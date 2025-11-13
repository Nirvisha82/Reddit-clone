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
	action := rand.Intn(10)
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
	case 8:
		s.simulateBookmarkPost(context)
	case 9:
		s.simulateUnbookmarkPost(context)
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

func (s *Simulator) simulateBookmarkPost(context actor.Context) {
	if len(s.posts) > 0 {
		context.Send(s.enginePID, &BookmarkPost{
			PostID:   s.randomPost(),
			Username: s.randomUser(),
		})
	}
}

func (s *Simulator) simulateUnbookmarkPost(context actor.Context) {
	if len(s.posts) > 0 {
		context.Send(s.enginePID, &UnbookmarkPost{
			PostID:   s.randomPost(),
			Username: s.randomUser(),
		})
	}
}

func (s *Simulator) printSimulationStats(context actor.Context) {
	fmt.Println("\nSimulation completed. Requesting final statistics...")
	context.Send(s.enginePID, &PrintSubredditPostsAndComments{})
	context.Send(s.enginePID, &GetSimulationStats{})
}

func (s *Simulator) printUserActions(context actor.Context) {
	context.Send(s.enginePID, &PrintUserActions{})
}
