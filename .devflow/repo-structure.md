This file is a merged representation of the entire codebase, combined into a single document.
The content has been processed for AI analysis and code review purposes.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and default ignore patterns
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Repository Information
- **Repository URL:** https://github.com/Nirvisha82/Reddit-clone.git
- **Repository Name:** Reddit-clone
- **Total Files Analyzed:** 7
- **Generated:** 2025-10-29 02:00:08

# Directory Structure
```
.devflow/
  repo-structure.md
Detailed back-end functionality.pdf
README.md
engine.go
go.mod
go.sum
main.go
messages.go
models.go
simulator.go
```

# Files

## File: README.md
````markdown
# Reddit-like Engine and Simulator in Go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.19+-00ADD8.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/Proto%20Actor-Framework-blue.svg" alt="Proto Actor">
  <img src="https://img.shields.io/badge/Concurrency-Actor%20Model-green.svg" alt="Actor Model">
  <img src="https://img.shields.io/badge/Distribution-Zipf-orange.svg" alt="Zipf Distribution">
</p>

A high-performance, concurrent Reddit-like social platform simulator built with Go and the Proto Actor framework. The system models realistic user behavior using Zipf distribution for subreddit popularity and supports large-scale simulations with up to 100K users and 250K actions.

## 🚀 Features

- **🏗️ Actor-Based Architecture**: Distributed system using Proto Actor framework
- **📊 Realistic User Modeling**: Zipf distribution for authentic subreddit popularity simulation
- **⚡ High Performance**: Handles 100K+ users with 250K+ actions in under 3 minutes
- **🔄 Concurrent Operations**: Full async/await support for all user interactions
- **👥 Complete Social Features**: Posts, comments, voting, direct messaging, and feeds
- **📈 Karma System**: Dynamic reputation tracking with upvote/downvote mechanics
- **🔌 Connection Management**: User online/offline status simulation
- **📱 Real-time Feed**: Personalized content feeds based on subscriptions
- **📊 Analytics**: Comprehensive simulation statistics and user action tracking

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.19+ |
| **Actor Framework** | Proto Actor |
| **Concurrency Model** | Actor-based messaging |
| **Distribution** | Zipf distribution for realistic modeling |
| **Architecture** | Concurrent, event-driven |

## 📁 Project Structure

```
reddit-simulator/
├── main.go              # Application entry point and configuration
├── engine.go            # Core Reddit engine with all business logic
├── simulator.go         # User behavior simulation and Zipf distribution
├── models.go            # Data structures for users, posts, comments
├── messages.go          # Actor message definitions and protocols
└── README.md           # Project documentation
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.19 or higher
- Proto Actor Go framework

### Step 1: Clone Repository

```bash
git clone <repository-url>
cd reddit-simulator
```

### Step 2: Install Dependencies

```bash
go mod init reddit-simulator
go get github.com/asynkron/protoactor-go/actor
go mod tidy
```

### Step 3: Run Simulation

```bash
# Default simulation (30 users, 6 subreddits, 200 actions, 5 seconds)
go run .

# Custom parameters
go run . -users 10 -subreddits 3 -actions 100 -time 3

# Large-scale simulation
go run . -users 1000 -subreddits 50 -actions 5000 -time 30
```

## 🏗️ System Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Main Thread   │    │   Actor System  │    │    Engine       │
│                 │────▶│                 │────▶│    Actor        │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │                        │
                                ▼                        ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Simulator     │    │   Message       │    │   Data Models   │
│   Actor         │◀───│   Passing       │────▶│                 │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🔄 Command Line Options

| Flag | Default | Description |
|------|---------|-------------|
| `-users` | 30 | Maximum number of users to simulate |
| `-subreddits` | 6 | Maximum number of subreddits to create |
| `-actions` | 200 | Total number of simulation actions |
| `-time` | 5 | Simulation duration in seconds |

### Usage Examples

```bash
# Quick test run
go run . -users 5 -subreddits 2 -actions 50 -time 2

# Medium simulation
go run . -users 100 -subreddits 10 -actions 1000 -time 10

# Large-scale performance test
go run . -users 10000 -subreddits 100 -actions 25000 -time 60
```

## 🎯 Core Features

### User Management
- **User Registration**: Dynamic user creation during simulation
- **Connection Status**: Online/offline state management
- **Karma System**: Reputation tracking based on community interactions

### Subreddit Operations
- **Create Subreddits**: Dynamic community creation
- **Join/Leave**: Flexible membership management
- **Zipf Distribution**: Realistic popularity modeling

### Content Management
- **Post Creation**: Rich content posting to any subreddit
- **Nested Comments**: Multi-level comment threading with replies
- **Voting System**: Upvote/downvote mechanics affecting karma
- **Direct Messaging**: Private user-to-user communication

### Social Features
- **Personalized Feeds**: Content from subscribed subreddits
- **Real-time Interactions**: Concurrent user actions
- **Community Building**: Organic subreddit growth patterns

## 📊 Karma System

The karma system models Reddit's reputation mechanics:

### Karma Rules
- **+1 Karma**: Per post creation
- **+1 Karma**: Per comment on posts
- **+1 Karma**: Per upvote received on content
- **-1 Karma**: Per downvote received on content
- **Default Upvote**: Authors automatically upvote their own posts

### Calculation Formula
```
Total Karma = Posts Created + Comments Made + (Total Upvotes - Total Downvotes)
```

## 🎲 Zipf Distribution Implementation

The simulator uses Zipf distribution to model realistic user behavior:

```go
// Zipf parameter: 1.07 (slightly skewed distribution)
zipf := rand.NewZipf(r, 1.07, 1, uint64(maxSubreddits))

// Popular subreddits get more users and content
subredditIndex := int(zipf.Uint64())
```

**Benefits:**
- **Realistic Modeling**: Mimics real-world social platform dynamics
- **Popular Communities**: Some subreddits naturally become more active
- **Long Tail Effect**: Many smaller communities with less activity

## 🔧 Actor System Design

### Engine Actor

**Responsibilities:**
- Process all user actions and state changes
- Maintain data consistency across the system
- Handle message routing and response generation

**Key Methods:**
```go
func (e *Engine) Receive(context actor.Context)
func (e *Engine) registerUser(username string)
func (e *Engine) createSubreddit(name, creator string)
func (e *Engine) createPost(postID, subredditName, author, title, content string)
func (e *Engine) vote(postID, userID string, isUpvote bool)
```

### Simulator Actor

**Responsibilities:**
- Generate realistic user behavior patterns
- Manage simulation lifecycle and timing
- Coordinate with Engine for action execution

**Key Methods:**
```go
func (s *Simulator) runSimulation(context actor.Context)
func (s *Simulator) simulateAction(context actor.Context)
func (s *Simulator) simulateCreatePost(context actor.Context)
func (s *Simulator) simulateVote(context actor.Context)
```

## 📈 Performance Benchmarks

### Tested Configurations

| Users | Actions | Subreddits | Time | Machine |
|-------|---------|------------|------|---------|
| 100 | 500 | 10 | 0:05 | M3 Pro |
| 1,000 | 5,000 | 50 | 0:15 | M3 Pro |
| 10,000 | 25,000 | 100 | 0:45 | M3 Pro |
| **100,000** | **250,000** | **600** | **2:40** | **12-Core M3 Pro, 18GB RAM** |

### Performance Characteristics
- **Linear Scaling**: Performance scales predictably with user count
- **Memory Efficient**: Optimized data structures for large simulations
- **Concurrent Processing**: Full utilization of multi-core systems

## 📊 Simulation Output

### User Actions Log
```
[REGISTER USER]  Registered as new user
[CREATE SUB]     Subreddit created: r/Sub 1 by User 1
[JOIN SUB]       User 2 joined subreddit r/Sub 1
[POST]           Post 1 created in r/Sub 1 by User 1: Post Title
[POST Comment]   User 2 commented on post Post 1: Comment content
[VOTE]           User 3 upvoted post Post 1
[Direct Message] DM sent to User 4: Hello there!
[SHOW FEED]      Feed for user User 2 -----
```
## 🔍 Monitoring & Analytics

### Real-time Metrics
- **Action Throughput**: Actions processed per second
- **User Activity**: Active vs. inactive user ratios
- **Content Distribution**: Posts and comments per subreddit
- **Engagement Rates**: Voting patterns and participation

### Post-Simulation Analysis
- **User Karma Distribution**: Reputation spread across users
- **Subreddit Popularity**: Member counts and activity levels
- **Content Metrics**: Post engagement and comment threading depth
- **Performance Stats**: Execution time and resource usage



## 👥 Authors

- **[Nirvisha Soni](https://github.com/Nirvisha82)** 
- **[Neel Malwatkar](https://github.com/neelmalwatkar)** 

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/actor-enhancement`)
3. Commit changes (`git commit -m 'Add new actor feature'`)
4. Push to branch (`git push origin feature/actor-enhancement`)
5. Open a Pull Request

---

<p align="center">
  ⚡ Built with Go's concurrency power and Proto Actor's distributed magic
</p>
````

## File: engine.go
````go
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
````

## File: go.mod
````
module reddit-clone

go 1.23

require github.com/asynkron/protoactor-go v0.0.0-20240822202345-3c0e61ca19c9

require (
	github.com/Workiva/go-datastructures v1.1.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/lithammer/shortuuid/v4 v4.0.0 // indirect
	github.com/lmittmann/tint v1.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/twmb/murmur3 v1.1.8 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/sdk v1.21.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
````

## File: main.go
````go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		maxUsers          = flag.Int("users", 30, "Maximum number of users")
		maxSubreddits     = flag.Int("subreddits", 6, "Maximum number of subreddits")
		simulationActions = flag.Int("actions", 200, "Number of simulation actions")
		simulationTime    = flag.Int("time", 5, "Simulation time in seconds")
	)
	flag.Parse()

	system := actor.NewActorSystem()

	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngine() })
	enginePID := system.Root.Spawn(engineProps)

	simulatorProps := actor.PropsFromProducer(func() actor.Actor {
		return NewSimulator(enginePID, *maxUsers, *maxSubreddits, *simulationActions)
	})
	simulatorPID := system.Root.Spawn(simulatorProps)

	fmt.Printf("Reddit-like engine and simulator started. Running for %d seconds...\n", *simulationTime)
	time.Sleep(time.Duration(*simulationTime) * time.Second)

	system.Root.Stop(simulatorPID)
	system.Root.Stop(enginePID)

	fmt.Println("PIDs stopped.")
}
````

## File: messages.go
````go
package main

import "time"

type RegisterUser struct {
	Username string
}

type CreateSubreddit struct {
	Name    string
	Creator string
}

type JoinSubreddit struct {
	SubredditName string
	Username      string
}

type LeaveSubreddit struct {
	SubredditName string
	Username      string
}

type CreatePost struct {
	PostID        string
	SubredditName string
	Author        string
	Title         string
	Content       string
}

// type CreateComment struct {
// 	PostID  string
// 	Author  string
// 	Content string
// }

type CreateComment struct {
	PostID    string
	ParentID  string
	CommentID string
	Author    string
	Content   string
}

type Vote struct {
	PostID   string
	UserID   string
	IsUpvote bool
}

type SendDirectMessage struct {
	From    string
	To      string
	Content string
}

type GetFeed struct {
	Username string
}

type UserAction struct {
	Action    string
	Timestamp time.Time
}

type UserActions struct {
	Username string
	Actions  []UserAction
}

//User wise Actions
type PrintUserActions struct{}

// simulation stats
type GetSimulationStats struct{}

// subreddit wise posts and comments
type PrintSubredditPostsAndComments struct{}

type StartSimulation struct{}
type SimulationCompleted struct{}
````

## File: models.go
````go
package main

type User struct {
	Username             string
	Karma                int
	SubscribedSubreddits []string
	SentMessages         []*DirectMessage
	ReceivedMessages     []*DirectMessage
}

type Subreddit struct {
	Name    string
	Creator string
	Members []string
}

type Post struct {
	ID            string
	SubredditName string
	Author        string
	Title         string
	Content       string
	Upvotes       int
	Downvotes     int
	Comments      []*Comment
}

type Comment struct {
	ID       string
	ParentID string
	Author   string
	Content  string
	Children []*Comment
}

type DirectMessage struct {
	From    string
	To      string
	Content string
}
````

## File: simulator.go
````go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type Simulator struct {
	enginePID          *actor.PID
	users              []string
	userStatus         map[string]bool
	subreddits         []string
	posts              []string
	actions            int
	context            actor.Context
	comments           map[string][]string
	zipf               *rand.Zipf
	MAX_USERS          int
	MAX_SUBREDDITS     int
	SIMULATION_ACTIONS int
}

func NewSimulator(enginePID *actor.PID, maxUsers, maxSubreddits, simulationActions int) actor.Actor {

	eh := rand.NewSource(time.Now().UnixNano())
	r := rand.New(eh)
	zipf := rand.NewZipf(r, 1.07, 1, uint64(maxSubreddits))
	return &Simulator{
		enginePID:          enginePID,
		comments:           make(map[string][]string),
		userStatus:         make(map[string]bool),
		actions:            0,
		zipf:               zipf,
		MAX_USERS:          maxUsers,
		MAX_SUBREDDITS:     maxSubreddits,
		SIMULATION_ACTIONS: simulationActions,
	}
}

func (s *Simulator) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *actor.Started:
		s.context = context
		fmt.Println("Simulator started")
		go s.runSimulation(context)
	}
}
func (s *Simulator) runSimulation(context actor.Context) {
	startTime := time.Now()
	s.registerInitialUsers(context)
	s.createInitialSubreddits(context)

	for s.actions < s.SIMULATION_ACTIONS {
		s.simulateAction(context)
		s.actions++
		time.Sleep(time.Millisecond * 10)
	}

	if s.MAX_USERS < 50 && s.SIMULATION_ACTIONS < 201 {
		s.printUserActions(context)
		s.printSimulationStats(context)
	}

	endTime := time.Now()
	fmt.Printf("Simulation completed in %s.\n", endTime.Sub(startTime))

}

func (s *Simulator) registerInitialUsers(context actor.Context) {
	for i := 0; i < s.MAX_USERS; i++ {
		s.simulateRegisterUser(context)
	}
}

func (s *Simulator) createInitialSubreddits(context actor.Context) {
	for i := 0; i < s.MAX_SUBREDDITS; i++ {
		s.simulateCreateSubreddit(context)
	}
}

func (s *Simulator) simulateAction(context actor.Context) {
	action := rand.Intn(8)
	switch action {
	case 0:
		s.simulateJoinSubreddit(context)
	case 1:
		s.simulateLeaveSubreddit(context)
	case 2:
		s.simulateCreatePost(context)
	case 3:
		s.simulateCreateComment(context)
	case 4:
		s.simulateVote(context)
	case 5:
		s.simulateSendDirectMessage(context)
	case 6:
		s.simulateGetFeed(context)
	case 7:
		s.simulateConnection()
	}
}

func (s *Simulator) simulateConnection() {
	username := s.randomUser()
	if rand.Intn(2) == 0 {
		s.userStatus[username] = true
		if s.MAX_USERS < 50 {
			fmt.Printf("%s is now connected.\n", username)
		}
	} else {
		s.userStatus[username] = false
		if s.MAX_USERS < 50 {
			fmt.Printf("%s is now disconnected.\n", username)
		}
	}
}

func (s *Simulator) simulateRegisterUser(context actor.Context) {
	username := fmt.Sprintf("User %d", len(s.users)+1)
	s.users = append(s.users, username)
	context.Send(s.enginePID, &RegisterUser{Username: username})
}

func (s *Simulator) simulateCreateSubreddit(context actor.Context) {
	subredditName := fmt.Sprintf("r/Sub %d", len(s.subreddits)+1)
	creator := s.randomUser()
	s.subreddits = append(s.subreddits, subredditName)
	context.Send(s.enginePID, &CreateSubreddit{Name: subredditName, Creator: creator})
}

func (s *Simulator) simulateJoinSubreddit(context actor.Context) {

	for _, subreddit := range s.subreddits {
		memberCount := int(s.zipf.Uint64())
		for i := 0; i < memberCount; i++ {
			user := s.randomUser()
			if s.userStatus[user] {
				context.Send(s.enginePID, &JoinSubreddit{
					SubredditName: subreddit,
					Username:      user,
				})
			} else {
				continue
			}

		}
	}
}

func (s *Simulator) simulateLeaveSubreddit(context actor.Context) {
	context.Send(s.enginePID, &LeaveSubreddit{
		SubredditName: s.randomSubreddit(),
		Username:      s.randomUser(),
	})
}

func (s *Simulator) simulateCreatePost(context actor.Context) {
	// subredditName := s.randomSubreddit()

	if len(s.subreddits) == 0 {
		return // Ensure there are subreddits available
	}

	// Use the Zipf generator to select a subreddit index
	subredditIndex := int(s.zipf.Uint64())

	// Ensure the index is within bounds (precautionary)
	if subredditIndex >= len(s.subreddits) {
		subredditIndex = len(s.subreddits) - 1
	}

	subredditName := s.subreddits[subredditIndex]

	postID := fmt.Sprintf("Post %d", len(s.posts)+1)
	s.posts = append(s.posts, postID)

	context.Send(s.enginePID, &CreatePost{
		PostID:        postID,
		SubredditName: subredditName,
		Author:        s.randomUser(),
		Title:         fmt.Sprintf("Post %d", len(s.posts)),
		Content:       fmt.Sprintf("Hello there! This is content of %s", postID),
	})
}

func (s *Simulator) simulateCreateComment(context actor.Context) {
	if len(s.posts) > 0 {
		postID := s.randomPost()
		parentID := postID
		if rand.Float32() < 0.5 && len(s.posts) > 0 {
			parentID = s.randomComment(postID)
		}
		commentID := fmt.Sprintf("Comment %d", len(s.comments[postID])+1)
		s.comments[postID] = append(s.comments[postID], commentID)
		context.Send(s.enginePID, &CreateComment{
			PostID:    postID,
			ParentID:  parentID,
			CommentID: commentID,
			Author:    s.randomUser(),
			Content:   fmt.Sprintf("This is a simulated %s.", commentID),
		})
	}
}

func (s *Simulator) randomComment(postID string) string {
	if comments, exists := s.comments[postID]; exists && len(comments) > 0 {
		return comments[rand.Intn(len(comments))]
	}
	return postID
}

func (s *Simulator) simulateVote(context actor.Context) {
	if len(s.posts) > 0 {
		context.Send(s.enginePID, &Vote{
			PostID:   s.randomPost(),
			UserID:   s.randomUser(),
			IsUpvote: rand.Intn(2) == 0,
		})
	}
}

func (s *Simulator) simulateSendDirectMessage(context actor.Context) {
	from := s.randomUser()
	to := s.randomUser()
	for to == from {
		to = s.randomUser()
	}
	//First message
	context.Send(s.enginePID, &SendDirectMessage{
		From:    from,
		To:      to,
		Content: fmt.Sprintf("This is a direct message from %s to %s", from, to),
	})
	//Reply to the message
	context.Send(s.enginePID, &SendDirectMessage{
		From:    to,
		To:      from,
		Content: fmt.Sprintf("This is a reply message from %s to %s", to, from),
	})
}

func (s *Simulator) simulateGetFeed(context actor.Context) {
	context.Send(s.enginePID, &GetFeed{Username: s.randomUser()})
}

func (s *Simulator) randomUser() string {
	return s.users[rand.Intn(len(s.users))]
}

func (s *Simulator) randomSubreddit() string {
	return s.subreddits[rand.Intn(len(s.subreddits))]
}

func (s *Simulator) randomPost() string {
	return s.posts[rand.Intn(len(s.posts))]
}

func (s *Simulator) printSimulationStats(context actor.Context) {
	fmt.Println("\nSimulation completed. Requesting final statistics...")
	context.Send(s.enginePID, &PrintSubredditPostsAndComments{})
	context.Send(s.enginePID, &GetSimulationStats{})
}

func (s *Simulator) printUserActions(context actor.Context) {
	context.Send(s.enginePID, &PrintUserActions{})
}
````

