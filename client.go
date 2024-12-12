package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EngineClient struct {
	engine *Engine
}

func NewEngineClient(engine *Engine) *EngineClient {
	return &EngineClient{engine: engine}
}

func (ec *EngineClient) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var user struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.registerUser(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "User already exists"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	msg_str := fmt.Sprintf("User u/%s registered", user.Username)
	json.NewEncoder(w).Encode(map[string]string{"message": msg_str})
}

func (ec *EngineClient) HandleCreateSubreddit(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var subreddit struct {
		Name    string `json:"name"`
		Creator string `json:"creator"`
	}
	var response struct {
		msg string
	}
	if err := json.NewDecoder(r.Body).Decode(&subreddit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errs := e.createSubreddit(subreddit.Name, subreddit.Creator)
	if errs != nil {
		w.WriteHeader(http.StatusFailedDependency)
		msg := fmt.Sprintf("r/%s could not be created as u/%s does not exist.", subreddit.Name, subreddit.Creator)
		response.msg = msg
	} else {
		msg := fmt.Sprintf("r/%s created by u/%s.", subreddit.Name, subreddit.Creator)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response.msg = msg

	}
	json.NewEncoder(w).Encode(map[string]string{"message": response.msg})
}

func (ec *EngineClient) HandleJoinSubreddit(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var joinsub struct {
		SubName string
		User    string
	}
	var response struct {
		msg string
	}

	if err := json.NewDecoder(r.Body).Decode(&joinsub); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errs := e.joinSubreddit(joinsub.SubName, joinsub.User)
	if errs != nil {
		w.WriteHeader(http.StatusFailedDependency)
		msg := fmt.Sprintf("Join Subreddit could not be complete because: %s", errs)
		response.msg = msg
	} else {
		msg := fmt.Sprintf("r/%s joined by u/%s ", joinsub.SubName, joinsub.User)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response.msg = msg

	}
	json.NewEncoder(w).Encode(map[string]string{"message": response.msg})

}
func (ec *EngineClient) HandleLeaveSubreddit(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var request struct {
		SubName string
		User    string
	}
	var response struct {
		msg string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errs := e.leaveSubreddit(request.SubName, request.User)
	if errs != nil {
		w.WriteHeader(http.StatusFailedDependency)
		msg := fmt.Sprintf("Leave Subreddit could not be complete because: %s", errs)
		response.msg = msg
	} else {
		msg := fmt.Sprintf("r/%s left by u/%s ", request.SubName, request.User)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		response.msg = msg

	}
	json.NewEncoder(w).Encode(map[string]string{"message": response.msg})

}
func (ec *EngineClient) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var request struct {
		SubredditName string `json:"subredditName"`
		Author        string `json:"author"`
		Title         string `json:"title"`
		Content       string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postID := fmt.Sprintf("Post %d", len(e.posts)+1)
	errs := e.createPost(postID, request.SubredditName, request.Author, request.Title, request.Content)
	if errs != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(map[string]string{"error": errs.Error()})
	} else {
		w.WriteHeader(http.StatusCreated)
		msg := fmt.Sprintf("u/%s made a post to r/%s. PostID: %s", request.Author, request.SubredditName, postID)
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
	}
}

func (ec *EngineClient) HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var request struct {
		PostID   string `json:"postID"`
		ParentID string `json:"parentID"`
		Author   string `json:"author"`
		Content  string `json:"content"`
	}
	var response struct {
		ReplyID string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentID := fmt.Sprintf("Comment%d", len(e.posts[request.PostID].Comments)+1)
	err := e.createComment(request.PostID, request.ParentID, commentID, request.Author, request.Content)

	w.Header().Set("Content-Type", "application/json")
	reply_str := fmt.Sprintf("u/%s replied to %s", request.Author, request.ParentID)
	if err != nil {
		if reply_str == err.Error() {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		if request.ParentID != request.PostID {
			msg := fmt.Sprintf("u/%s replied to %s", request.Author, request.ParentID)
			for _, comment := range e.posts[request.PostID].Comments {
				// fmt.Printf("Checking %s if is parent of %s", comment.ID, ParentID)
				if comment.ID == request.ParentID {
					response.ReplyID = fmt.Sprintf("Reply%d", len(comment.Children))
					break
				}
			}
			json.NewEncoder(w).Encode(map[string]string{"message": msg, "ReplyID": response.ReplyID})
		} else {
			msg := fmt.Sprintf("u/%s commented on %s", request.Author, request.PostID)
			json.NewEncoder(w).Encode(map[string]string{"message": msg, "commentID": commentID})
		}
	}
}
func (ec *EngineClient) HandleVote(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var request struct {
		PostID   string `json:"postID"`
		UserID   string `json:"userID"`
		IsUpvote bool   `json:"isUpvote"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.vote(request.PostID, request.UserID, request.IsUpvote)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		voteType := "upvoted"
		if !request.IsUpvote {
			voteType = "downvoted"
		}
		msg := fmt.Sprintf("User %s %s post %s", request.UserID, voteType, request.PostID)
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
	}
}

func (ec *EngineClient) HandleGetFeed(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var user struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// username := r.URL.Query().Get("username")
	// if username == "" {
	// 	http.Error(w, "Username is required", http.StatusBadRequest)
	// 	return
	// }

	userdata, exists := e.users[user.Username]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var feed []*Post
	for _, post := range e.posts {
		if contains(userdata.SubscribedSubreddits, post.SubredditName) {
			feed = append(feed, post)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(feed) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "No posts in feed"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": user.Username,
		"feed":     feed,
	})

	log_str := fmt.Sprintf("[SHOW FEED] Feed retrieved for user %s", user.Username)
	e.logUserAction(user.Username, log_str)
}
func (ec *EngineClient) HandleSendDirectMessage(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	var request struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.sendDirectMessage(request.From, request.To, request.Content)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("Direct message sent from u/%s to u/%s", request.From, request.To)
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
	}
}
func (ec *EngineClient) getLogs(w http.ResponseWriter, r *http.Request) {
	e := ec.engine
	e.printAllUserActions()
	e.getSimulationStats()
	e.printSubredditPostsAndComments()
	json.NewEncoder(w).Encode(map[string]string{"message": "Logs and Summary written on the terminal"})

}
