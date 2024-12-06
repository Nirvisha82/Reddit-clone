package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type Engine struct {
	users       map[string]*User
	subreddits  map[string]*Subreddit
	posts       map[string]*Post
	userActions map[string]*UserActions
	// server      *http.Server
	router *http.ServeMux
}

func (e *Engine) SetupRoutes() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/register", e.HandleRegisterUser)
	// mux.HandleFunc("/subreddit", e.HandleCreateSubreddit)
	// return mux
	e.router.HandleFunc("/register", e.HandleRegisterUser)
	e.router.HandleFunc("/createsub", e.HandleCreateSubreddit)
	e.router.HandleFunc("/joinsub", e.HandleJoinSubreddit)
	e.router.HandleFunc("/leavesub", e.HandleLeaveSubreddit)
	e.router.HandleFunc("/createpost", e.HandleCreatePost)
	e.router.HandleFunc("/createcomment", e.HandleCreateComment)
	e.router.HandleFunc("/vote", e.HandleVote)

	// Add more routes here
}

func NewEngine() *Engine {
	e := &Engine{
		users:       make(map[string]*User),
		subreddits:  make(map[string]*Subreddit),
		posts:       make(map[string]*Post),
		userActions: make(map[string]*UserActions),
		router:      http.NewServeMux(),
	}
	e.SetupRoutes()
	return e
}

func (e *Engine) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Engine started")
		go func() {
			if err := http.ListenAndServe(":8080", e.router); err != nil {
				fmt.Printf("Server error: %v\n", err)
			}
		}()
	case *RegisterUser:
		e.registerUser(msg.Username)
	case *CreateSubreddit:
		e.createSubreddit(msg.Name, msg.Creator)
	case *JoinSubreddit:
		e.joinSubreddit(msg.SubredditName, msg.Username)
	case *LeaveSubreddit:
		e.leaveSubreddit(msg.SubredditName, msg.Username)
	case *CreatePost:
		e.createPost(msg.PostID, msg.SubredditName, msg.Author, msg.Title, msg.Content)
	case *CreateComment:
		e.createComment(msg.PostID, msg.ParentID, msg.CommentID, msg.Author, msg.Content)
	case *Vote:
		e.vote(msg.PostID, msg.UserID, msg.IsUpvote)
	case *SendDirectMessage:
		e.sendDirectMessage(msg.From, msg.To, msg.Content)
	case *GetFeed:
		e.getFeed(msg.Username)
	case *GetSimulationStats:
		e.getSimulationStats()
	case *PrintUserActions:
		e.printAllUserActions()
	case *PrintSubredditPostsAndComments:
		e.printSubredditPostsAndComments()
	}
}

// Request handlers
func (e *Engine) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
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

func (e *Engine) HandleCreateSubreddit(w http.ResponseWriter, r *http.Request) {
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

func (e *Engine) HandleJoinSubreddit(w http.ResponseWriter, r *http.Request) {
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
func (e *Engine) HandleLeaveSubreddit(w http.ResponseWriter, r *http.Request) {
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
func (e *Engine) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
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

func (e *Engine) HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	var request struct {
		PostID   string `json:"postID"`
		ParentID string `json:"parentID"`
		Author   string `json:"author"`
		Content  string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentID := fmt.Sprintf("Comment%d", len(e.posts[request.PostID].Comments)+1)
	err := e.createComment(request.PostID, request.ParentID, commentID, request.Author, request.Content)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusCreated)
		msg := fmt.Sprintf("u/%s commented on post %s", request.Author, request.PostID)
		json.NewEncoder(w).Encode(map[string]string{"message": msg, "commentID": commentID})
	}
}
func (e *Engine) HandleVote(w http.ResponseWriter, r *http.Request) {
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

// Functions
func (e *Engine) registerUser(username string) error {
	if _, exists := e.users[username]; exists {
		return fmt.Errorf("user already exists")
	}
	e.users[username] = &User{Username: username, Karma: 0}
	fmt.Printf("u/%s registered!\n", username)
	e.logUserAction(username, "[REGISTER USER] Registered as new user")
	return nil
}

func (e *Engine) createSubreddit(name, creator string) error {
	if _, exists := e.subreddits[name]; !exists {
		if _, userExists := e.users[creator]; userExists { //only create sub if user exists
			e.subreddits[name] = &Subreddit{Name: name, Creator: creator, Members: []string{creator}}
			//fmt.Printf("[CREATE SUB] Subreddit created: %s by %s\n", name, creator)
			log_str := fmt.Sprintf("[CREATE SUB]     Subreddit created: %s by %s", name, creator)
			e.logUserAction(creator, log_str)
			return nil
		}
	}
	return fmt.Errorf("the user does not exist, therefore subreddit cannot be created")
}

func (e *Engine) joinSubreddit(subredditName, username string) error {
	if subreddit, exists := e.subreddits[subredditName]; exists {
		if user, userExists := e.users[username]; userExists {
			if !contains(subreddit.Members, username) {
				subreddit.Members = append(subreddit.Members, username)
				user.SubscribedSubreddits = append(user.SubscribedSubreddits, subredditName)
				//fmt.Printf("[JOIN SUB] %s joined subreddit %s\n", username, subredditName)
				log_str := fmt.Sprintf("[JOIN SUB]       %s joined subreddit %s", username, subredditName)
				e.logUserAction(username, log_str)
				return nil
			}
		} else {
			return fmt.Errorf("u/%s does not exist", username)

		}

	}
	return fmt.Errorf("r/%s does not exist", subredditName)
}
func (e *Engine) leaveSubreddit(subredditName, username string) error {
	if subreddit, exists := e.subreddits[subredditName]; exists {
		if user, userExists := e.users[username]; userExists {
			if contains(subreddit.Members, username) {
				subreddit.Members = remove(subreddit.Members, username)
				user.SubscribedSubreddits = remove(user.SubscribedSubreddits, subredditName)
				//fmt.Printf("[LEAVE SUB] %s left subreddit %s\n", username, subredditName)
				log_str := fmt.Sprintf("[LEAVE SUB]      %s left subreddit %s", username, subredditName)

				e.logUserAction(username, log_str)
				return nil
			}
		} else {
			return fmt.Errorf("u/%s does not exist", username)

		}

	}
	return fmt.Errorf("r/%s does not exist", subredditName)
}

func (e *Engine) createPost(postID, subredditName, author, title, content string) error {
	if _, exists := e.subreddits[subredditName]; exists {
		fmt.Printf("Creating a post in %s\n", subredditName)
		if _, userExists := e.users[author]; userExists {
			fmt.Printf("Author %s Verified", author)
			e.posts[postID] = &Post{ID: postID, SubredditName: subredditName, Author: author, Title: title, Content: content}
			// by default upvote for post by author when posted & increased karma
			post := e.posts[postID]
			post.Upvotes++
			e.users[post.Author].Karma++
			// fmt.Printf("[POST] Post created in %s by %s: %s\n", subredditName, author, title)
			log_str := fmt.Sprintf("[POST]           %s created in %s by %s: %s", postID, subredditName, author, title)

			e.logUserAction(author, log_str)
			return nil
		} else {
			return fmt.Errorf("u/%s does not exist", author)
		}
	}
	fmt.Printf("Searching for %s", subredditName)
	return fmt.Errorf("r/%s does not exist", subredditName)

}

func (e *Engine) createComment(postID, parentID, commentID, author, content string) error {
	if post, exists := e.posts[postID]; exists {
		if _, userExists := e.users[author]; userExists {
			newComment := &Comment{ID: commentID, ParentID: parentID, Author: author, Content: content}

			if parentID == postID {
				post.Comments = append(post.Comments, newComment)
				e.users[author].Karma++
			} else {

				if !e.addChildComment(post.Comments, parentID, author, newComment) {
					return fmt.Errorf("parent comment %s does not exist", parentID)
				}
			}

			log_str := fmt.Sprintf("[POST Comment]   %s commented on post %s: %s", author, postID, content)
			e.logUserAction(author, log_str)
			return nil
		}
		return fmt.Errorf("u/%s does not exist", author)
	}
	return fmt.Errorf("post %s does not exist", postID)
}

// func (e *Engine) addChildComment(comments []*Comment, parentID string, author string, newComment *Comment) {
// 	for _, comment := range comments {
// 		if comment.ID == parentID {
// 			commentReply := fmt.Sprintf("Reply %d to %s", len(comment.Children)+1, comment.ID)
// 			newComment.Content = commentReply
// 			comment.Children = append(comment.Children, newComment)
// 			// log_str := fmt.Sprintf("[POST Comment]   %s commented on comment %s: %s", author, comment.ID, newComment.Content)
// 			// log_str := fmt.Sprintf("[COMMENT REPLY]  %s commented on %s: %s", author, comment.ID, newComment.Content)
// 			e.users[author].Karma++
// 			log_str := fmt.Sprintf("[COMMENT REPLY]  %s commented on %s: %s", author, comment.ID, commentReply)
// 			e.logUserAction(author, log_str)

//				return
//			}
//			e.addChildComment(comment.Children, parentID, author, newComment)
//		}
//	}

// adding child comment not behaving as exoected. instead create another method to add a reply.
func (e *Engine) addChildComment(comments []*Comment, parentID string, author string, newComment *Comment) bool {
	for _, comment := range comments {
		if comment.ID == parentID {
			commentReply := fmt.Sprintf("Reply %d to %s", len(comment.Children)+1, comment.ID)
			fmt.Println(commentReply)
			newComment.Content = commentReply
			newComment.ID = fmt.Sprintf("Child Comment %d", len(comment.Children)+1)
			comment.Children = append(comment.Children, newComment)
			e.users[author].Karma++
			log_str := fmt.Sprintf("[COMMENT REPLY] %s commented on %s: %s", author, comment.ID, commentReply)
			e.logUserAction(author, log_str)
			return true
		}
		if e.addChildComment(comment.Children, parentID, author, newComment) {
			return true
		}
	}
	return false
}
func (e *Engine) vote(postID, userID string, isUpvote bool) error {
	if post, exists := e.posts[postID]; exists {
		if _, userExists := e.users[userID]; userExists {
			if isUpvote {
				post.Upvotes++
				e.users[post.Author].Karma++
			} else {
				post.Downvotes++
				e.users[post.Author].Karma--
			}
			voteType := "upvoted"
			if !isUpvote {
				voteType = "downvoted"
			}
			log_str := fmt.Sprintf("[VOTE] %s %s post %s", userID, voteType, postID)
			e.logUserAction(userID, log_str)
			return nil
		}
		return fmt.Errorf("user %s does not exist", userID)
	}
	return fmt.Errorf("post %s does not exist", postID)
}

func (e *Engine) sendDirectMessage(from, to, content string) error {
	if fromUser, fromExists := e.users[from]; fromExists {
		if toUser, toExists := e.users[to]; toExists {
			message := &DirectMessage{From: from, To: to, Content: content}
			fromUser.SentMessages = append(fromUser.SentMessages, message)
			toUser.ReceivedMessages = append(toUser.ReceivedMessages, message)
			log_str := fmt.Sprintf("[Direct Message] DM sent to %s: %s", to, content)
			e.logUserAction(from, log_str)
			return nil
		}
		return fmt.Errorf("recipient %s does not exist", to)
	}
	return fmt.Errorf("sender %s does not exist", from)
}

// func (e *Engine) vote(postID, userID string, isUpvote bool) {
// 	if post, exists := e.posts[postID]; exists {
// 		if isUpvote {
// 			post.Upvotes++
// 			e.users[post.Author].Karma++
// 		} else {
// 			post.Downvotes++
// 			e.users[post.Author].Karma--
// 		}
// 		voteType := "upvoted"
// 		if !isUpvote {
// 			voteType = "downvoted"
// 		}
// 		//fmt.Printf("[VOTE] %s %s post %s\n", userID, voteType, postID)
// 		log_str := fmt.Sprintf("[VOTE]           %s %s post %s", userID, voteType, postID)
// 		e.logUserAction(userID, log_str)

// 	}
// }

// func (e *Engine) sendDirectMessage(from, to, content string) {
// 	if fromUser, exists := e.users[from]; exists {
// 		if toUser, exists := e.users[to]; exists {
// 			message := &DirectMessage{From: from, To: to, Content: content}
// 			fromUser.SentMessages = append(fromUser.SentMessages, message)
// 			toUser.ReceivedMessages = append(toUser.ReceivedMessages, message)
// 			//fmt.Printf("[Direct Message] DM sent from %s to %s: %s\n", from, to, content)
// 			log_str := fmt.Sprintf("[Direct Message] DM sent to %s: %s", to, content)
// 			e.logUserAction(from, log_str)
// 		}
// 	}
// }

func (e *Engine) getFeed(username string) {
	if user, exists := e.users[username]; exists {
		var feed []*Post
		for _, post := range e.posts {
			if contains(user.SubscribedSubreddits, post.SubredditName) {
				feed = append(feed, post)
			}
		}
		//fmt.Printf("[SHOW FEED] Feed for user %s -----\n ", username)
		log_str := fmt.Sprintf("[SHOW FEED]      Feed for user %s ----- ", username)

		e.logUserAction(username, log_str)

		for _, post := range feed {
			//fmt.Printf("%s (in %s)\n", post.Title, post.SubredditName)
			log_str := fmt.Sprintf("                 %s (in %s)", post.Title, post.SubredditName)
			e.logUserAction(username, log_str)
		}
	}
}

func (e *Engine) getSimulationStats() {
	fmt.Println("\n\n----Simulation Statistics----")
	fmt.Printf("Total Users: %d\n", len(e.users))
	fmt.Printf("Total Subreddits: %d\n", len(e.subreddits))
	fmt.Printf("Total Posts: %d\n", len(e.posts))

	fmt.Println("\nUser Karma:")
	var users []string
	for username := range e.users {
		users = append(users, username)
	}
	sort.Strings(users)
	for _, username := range users {
		fmt.Printf("%s: %d karma\n", username, e.users[username].Karma)
	}

	fmt.Println("\nPost Statistics:")
	var posts []string
	for postID := range e.posts {
		posts = append(posts, postID)
	}
	sort.Strings(posts)
	for _, postID := range posts {
		post := e.posts[postID]
		fmt.Printf("%s by %s in %s: %d upvotes, %d downvotes, %d comments\n",
			postID, post.Author, post.SubredditName, post.Upvotes, post.Downvotes, len(post.Comments))
	}
}
func (e *Engine) logUserAction(username, action string) {
	if _, exists := e.userActions[username]; !exists {
		e.userActions[username] = &UserActions{Username: username, Actions: []UserAction{}}
	}
	e.userActions[username].Actions = append(e.userActions[username].Actions, UserAction{
		Action:    action,
		Timestamp: time.Now(),
	})
}
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func remove(slice []string, item string) []string {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (e *Engine) printAllUserActions() {
	fmt.Println("------- Printing User Actions --------")
	for _, userActions := range e.userActions {
		fmt.Printf("\n%s Actions:\n", userActions.Username)
		for _, action := range userActions.Actions {
			// fmt.Printf("- %s: %s\n", action.Timestamp.Format(time.RFC3339), action.Action)
			fmt.Printf("%s\n", action.Action)
		}
	}
}

func (e *Engine) printSubredditPostsAndComments() {
	fmt.Println("\n--- Subreddit-wise Posts and Comments ---")
	for subredditName := range e.subreddits {
		fmt.Printf("\nSubreddit: %s", subredditName)
		subredditPosts := 0
		for _, post := range e.posts {
			if post.SubredditName == subredditName {
				subredditPosts++
				fmt.Printf("\n>Post %d: %s by %s\n", subredditPosts, post.Title, post.Author)
				fmt.Printf(" Content: %s\n", post.Content)
				fmt.Printf(" Upvotes: %d | Downvotes: %d\n", post.Upvotes, post.Downvotes)
				if len(post.Comments) > 0 {
					fmt.Println(" Comments:")
					e.printComments(post.Comments, 1)

				} else {
					fmt.Println(" No comments yet.")
				}
			}
		}
		if subredditPosts == 0 {
			fmt.Println("\nNo posts in this subreddit yet.")
			fmt.Printf("\n\n-> Summary:\n Total %d Posts.\n Total %d members.\n\n", subredditPosts, len(e.subreddits[subredditName].Members))
		} else {
			fmt.Printf("\n\n-> Summary:\n Total %d Posts.\n Total %d members.\n\n", subredditPosts, len(e.subreddits[subredditName].Members))
		}
	}
}

func (e *Engine) printComments(comments []*Comment, depth int) {
	for _, comment := range comments {
		indent := strings.Repeat("  ", depth)
		fmt.Printf("%s- %s: %s\n", indent, comment.Author, comment.Content)
		if len(comment.Children) > 0 {
			e.printComments(comment.Children, depth+1)
		}
	}
}
