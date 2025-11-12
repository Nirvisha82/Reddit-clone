package main

type User struct {
	Username             string
	Karma                int
	SubscribedSubreddits []string
	BookmarkedPosts      []string
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
