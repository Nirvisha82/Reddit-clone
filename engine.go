package main

import (
	"fmt"
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
}

func NewEngine() *Engine {
	return &Engine{
		users:       make(map[string]*User),
		subreddits:  make(map[string]*Subreddit),
		posts:       make(map[string]*Post),
		userActions: make(map[string]*UserActions),
	}
}

func (e *Engine) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Engine started")
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
		// e.createComment(msg.PostID, msg.Author, msg.Content)
		e.createComment(msg.PostID, msg.ParentID, msg.CommentID, msg.Author, msg.Content)
	case *Vote:
		e.vote(msg.PostID, msg.UserID, msg.IsUpvote)
	case *SendDirectMessage:
		e.sendDirectMessage(msg.From, msg.To, msg.Content)
	case *SharePostViaDirectMessage:
		e.sharePostViaDirectMessage(msg.From, msg.To, msg.PostID)
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

func (e *Engine) registerUser(username string) {
	if _, exists := e.users[username]; !exists {
		e.users[username] = &User{Username: username, Karma: 0}
		//fmt.Printf("[REGISTER USER] User registered: %s\n", username)
		e.logUserAction(username, "[REGISTER USER]  Registerd as new user")

	}
}

func (e *Engine) createSubreddit(name, creator string) {
	if _, exists := e.subreddits[name]; !exists {
		e.subreddits[name] = &Subreddit{Name: name, Creator: creator, Members: []string{creator}}
		//fmt.Printf("[CREATE SUB] Subreddit created: %s by %s\n", name, creator)
		log_str := fmt.Sprintf("[CREATE SUB]     Subreddit created: %s by %s", name, creator)

		e.logUserAction(creator, log_str)
	}
}

func (e *Engine) joinSubreddit(subredditName, username string) {
	if subreddit, exists := e.subreddits[subredditName]; exists {
		if user, userExists := e.users[username]; userExists {
			if !contains(subreddit.Members, username) {
				subreddit.Members = append(subreddit.Members, username)
				user.SubscribedSubreddits = append(user.SubscribedSubreddits, subredditName)
				//fmt.Printf("[JOIN SUB] %s joined subreddit %s\n", username, subredditName)
				log_str := fmt.Sprintf("[JOIN SUB]       %s joined subreddit %s", username, subredditName)
				e.logUserAction(username, log_str)
			}
		}
	}
}

func (e *Engine) leaveSubreddit(subredditName, username string) {
	if subreddit, exists := e.subreddits[subredditName]; exists {
		if user, userExists := e.users[username]; userExists {
			if contains(subreddit.Members, username) {
				subreddit.Members = remove(subreddit.Members, username)
				user.SubscribedSubreddits = remove(user.SubscribedSubreddits, subredditName)
				//fmt.Printf("[LEAVE SUB] %s left subreddit %s\n", username, subredditName)
				log_str := fmt.Sprintf("[LEAVE SUB]      %s left subreddit %s", username, subredditName)

				e.logUserAction(username, log_str)
			}
		}
	}
}

func (e *Engine) createPost(postID, subredditName, author, title, content string) {
	if _, exists := e.subreddits[subredditName]; exists {

		e.posts[postID] = &Post{ID: postID, SubredditName: subredditName, Author: author, Title: title, Content: content}
		// by default upvote for post by author when posted & increased karma
		post := e.posts[postID]
		post.Upvotes++
		e.users[post.Author].Karma++
		// fmt.Printf("[POST] Post created in %s by %s: %s\n", subredditName, author, title)
		log_str := fmt.Sprintf("[POST]           %s created in %s by %s: %s", postID, subredditName, author, title)

		e.logUserAction(author, log_str)
		// }
	}
}

func (e *Engine) createComment(postID, parentID, commentID, author, content string) {
	if post, exists := e.posts[postID]; exists {
		newComment := &Comment{ID: commentID, ParentID: parentID, Author: author, Content: content}

		if parentID == postID {
			post.Comments = append(post.Comments, newComment)
			e.users[author].Karma++
		} else {
			// fmt.Println(("Adding Child Comment"))
			e.addChildComment(post.Comments, parentID, author, newComment)
		}

		log_str := fmt.Sprintf("[POST Comment]   %s commented on post %s: %s", author, postID, content)
		e.logUserAction(author, log_str)
	}
}

func (e *Engine) addChildComment(comments []*Comment, parentID string, author string, newComment *Comment) {
	for _, comment := range comments {
		if comment.ID == parentID {
			commentReply := fmt.Sprintf("Reply %d to %s", len(comment.Children)+1, comment.ID)
			newComment.Content = commentReply
			comment.Children = append(comment.Children, newComment)
			// log_str := fmt.Sprintf("[POST Comment]   %s commented on comment %s: %s", author, comment.ID, newComment.Content)
			// log_str := fmt.Sprintf("[COMMENT REPLY]  %s commented on %s: %s", author, comment.ID, newComment.Content)
			e.users[author].Karma++
			log_str := fmt.Sprintf("[COMMENT REPLY]  %s commented on %s: %s", author, comment.ID, commentReply)
			e.logUserAction(author, log_str)

			return
		}
		e.addChildComment(comment.Children, parentID, author, newComment)
	}
}

func (e *Engine) vote(postID, userID string, isUpvote bool) {
	if post, exists := e.posts[postID]; exists {
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
		//fmt.Printf("[VOTE] %s %s post %s\n", userID, voteType, postID)
		log_str := fmt.Sprintf("[VOTE]           %s %s post %s", userID, voteType, postID)
		e.logUserAction(userID, log_str)

	}
}

func (e *Engine) sendDirectMessage(from, to, content string) {
	if fromUser, exists := e.users[from]; exists {
		if toUser, exists := e.users[to]; exists {
			message := &DirectMessage{From: from, To: to, Content: content}
			fromUser.SentMessages = append(fromUser.SentMessages, message)
			toUser.ReceivedMessages = append(toUser.ReceivedMessages, message)
			//fmt.Printf("[Direct Message] DM sent from %s to %s: %s\n", from, to, content)
			log_str := fmt.Sprintf("[Direct Message] DM sent to %s: %s", to, content)
			e.logUserAction(from, log_str)
		}
	}
}

func (e *Engine) sharePostViaDirectMessage(from, to, postID string) {
	if fromUser, exists := e.users[from]; exists {
		if toUser, exists := e.users[to]; exists {
			if post, postExists := e.posts[postID]; postExists {
				// Create a message that includes the shared post
				message := &DirectMessage{
					From:       from,
					To:         to,
					Content:    fmt.Sprintf("Check out this post: %s", post.Title),
					SharedPost: post,
				}
				fromUser.SentMessages = append(fromUser.SentMessages, message)
				toUser.ReceivedMessages = append(toUser.ReceivedMessages, message)
				
				log_str := fmt.Sprintf("[SHARE POST]     %s shared post '%s' (ID: %s) with %s", from, post.Title, postID, to)
				e.logUserAction(from, log_str)
				
				log_str = fmt.Sprintf("[RECEIVED SHARE] Received shared post '%s' from %s", post.Title, from)
				e.logUserAction(to, log_str)
			}
		}
	}
}

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
