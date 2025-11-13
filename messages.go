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

type BookmarkPost struct {
	PostID   string
	Username string
}

type UnbookmarkPost struct {
	PostID   string
	Username string
}
