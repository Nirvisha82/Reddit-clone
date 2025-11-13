# Post Bookmarking Feature Implementation

## Analysis

The Reddit-clone project is an actor-based simulation system built with Go and Proto Actor framework. Currently, it supports user registration, subreddit management, post creation, voting, commenting, and direct messaging. The issue requests adding **post bookmarking functionality** - allowing users to bookmark posts they like and unbookmark them.

### Current Architecture Overview
- **Engine Actor**: Centralized state management for users, subreddits, posts, and user actions
- **Simulator Actor**: Generates realistic user behavior patterns
- **Models**: Data structures representing entities (User, Post, Subreddit, Comment, DirectMessage)
- **Messages**: Communication protocol between actors (RegisterUser, CreatePost, Vote, etc.)

### Key Observations
1. The `User` struct currently tracks: Username, Karma, SubscribedSubreddits, SentMessages, ReceivedMessages
2. The `Post` struct contains: ID, SubredditName, Author, Title, Content, Upvotes, Downvotes, Comments
3. The Engine uses a centralized map-based storage pattern for all entities
4. User actions are logged with timestamps for detailed simulation tracking
5. The simulator randomly generates various user actions with predefined probabilities

### Implementation Strategy
To implement bookmarking, we need to:
1. Add a bookmarks collection to the `User` model
2. Create two new message types: `BookmarkPost` and `UnbookmarkPost`
3. Implement bookmark/unbookmark logic in the Engine actor
4. Add simulator actions to randomly bookmark/unbookmark posts
5. Add logging and statistics for bookmarked posts

---

## Affected Files

### 1. **models.go** - Data Model Changes
- Modify `User` struct to include bookmarks collection

### 2. **messages.go** - Message Type Additions
- Add `BookmarkPost` message type
- Add `UnbookmarkPost` message type
- Optional: Add `GetBookmarks` message type for retrieving user bookmarks

### 3. **engine.go** - Business Logic Implementation
- Add bookmark handling in the `Receive` method
- Implement `bookmarkPost()` method
- Implement `unbookmarkPost()` method
- Optional: Implement `getBookmarks()` method
- Update statistics printing to include bookmark information

### 4. **simulator.go** - Simulation Behavior
- Add bookmark/unbookmark actions to `simulateAction()` method
- Implement `simulateBookmarkPost()` method
- Implement `simulateUnbookmarkPost()` method
- Update action probability distribution

---

## Code Examples

### Example 1: User Model Update (models.go)
```go
type User struct {
	Username             string
	Karma                int
	SubscribedSubreddits []string
	BookmarkedPosts      []string          // NEW: Track bookmarked post IDs
	SentMessages         []*DirectMessage
	ReceivedMessages     []*DirectMessage
}
```

### Example 2: Message Types (messages.go)
```go
type BookmarkPost struct {
	PostID   string
	Username string
}

type UnbookmarkPost struct {
	PostID   string
	Username string
}

type GetBookmarks struct {
	Username string
}
```

### Example 3: Engine Message Handler (engine.go - Receive method)
```go
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
		e.createComment(msg.PostID, msg.ParentID, msg.CommentID, msg.Author, msg.Content)
	case *Vote:
		e.vote(msg.PostID, msg.UserID, msg.IsUpvote)
	case *SendDirectMessage:
		e.sendDirectMessage(msg.From, msg.To, msg.Content)
	case *GetFeed:
		e.getFeed(msg.Username)
	case *BookmarkPost:                    // NEW
		e.bookmarkPost(msg.PostID, msg.Username)
	case *UnbookmarkPost:                  // NEW
		e.unbookmarkPost(msg.PostID, msg.Username)
	case *GetBookmarks:                    // NEW
		e.getBookmarks(msg.Username)
	case *GetSimulationStats:
		e.getSimulationStats()
	case *PrintUserActions:
		e.printAllUserActions()
	case *PrintSubredditPostsAndComments:
		e.printSubredditPostsAndComments()
	}
}
```

### Example 4: Bookmark Implementation (engine.go)
```go
func (e *Engine) bookmarkPost(postID, username string) {
	if post, postExists := e.posts[postID]; postExists {
		if user, userExists := e.users[username]; userExists {
			// Check if already bookmarked
			if !contains(user.BookmarkedPosts, postID) {
				user.BookmarkedPosts = append(user.BookmarkedPosts, postID)
				log_str := fmt.Sprintf("[BOOKMARK]       %s bookmarked post %s by %s in %s", 
					username, postID, post.Author, post.SubredditName)
				e.logUserAction(username, log_str)
			}
		}
	}
}

func (e *Engine) unbookmarkPost(postID, username string) {
	if post, postExists := e.posts[postID]; postExists {
		if user, userExists := e.users[username]; userExists {
			// Check if bookmarked
			if contains(user.BookmarkedPosts, postID) {
				user.BookmarkedPosts = remove(user.BookmarkedPosts, postID)
				log_str := fmt.Sprintf("[UNBOOKMARK]     %s unbookmarked post %s by %s in %s", 
					username, postID, post.Author, post.SubredditName)
				e.logUserAction(username, log_str)
			}
		}
	}
}

func (e *Engine) getBookmarks(username string) {
	if user, exists := e.users[username]; exists {
		log_str := fmt.Sprintf("[SHOW BOOKMARKS] Bookmarks for user %s ----- ", username)
		e.logUserAction(username, log_str)
		
		for _, postID := range user.BookmarkedPosts {
			if post, postExists := e.posts[postID]; postExists {
				log_str := fmt.Sprintf("                 %s (in %s) by %s", 
					post.Title, post.SubredditName, post.Author)
				e.logUserAction(username, log_str)
			}
		}
	}
}
```

### Example 5: Simulator Integration (simulator.go)
```go
func (s *Simulator) simulateAction(context actor.Context) {
	action := rand.Intn(10)  // Changed from 8 to 10 to include new actions
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
	case 8:                              // NEW
		s.simulateBookmarkPost(context)
	case 9:                              // NEW
		s.simulateUnbookmarkPost(context)
	}
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
```

### Example 6: Statistics Update (engine.go - getSimulationStats method)
```go
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

	// NEW: Bookmark Statistics
	fmt.Println("\nUser Bookmarks:")
	for _, username := range users {
		bookmarkCount := len(e.users[username].BookmarkedPosts)
		if bookmarkCount > 0 {
			fmt.Printf("%s: %d bookmarked posts\n", username, bookmarkCount)
		}
	}

	fmt.Println("\nPost Statistics:")
	var posts []string
	for postID := range e.posts {
		posts = append(posts, postID)
	}
	sort.Strings(posts)
	for _, postID := range posts {
		post := e.posts[postID]
		
		// Count bookmarks for this post
		bookmarkCount := 0
		for _, user := range e.users {
			if contains(user.BookmarkedPosts, postID) {
				bookmarkCount++
			}
		}
		
		fmt.Printf("%s by %s in %s: %d upvotes, %d downvotes, %d comments, %d bookmarks\n",
			postID, post.Author, post.SubredditName, post.Upvotes, post.Downvotes, 
			len(post.Comments), bookmarkCount)
	}
}
```

---

## Implementation Steps

### Step 1: Update Data Models (models.go)
1. Open `models.go`
2. Locate the `User` struct definition
3. Add a new field: `BookmarkedPosts []string` to track bookmarked post IDs
4. Ensure the field is initialized as an empty slice in the `registerUser()` method

**Rationale**: This stores which posts each user has bookmarked, using post IDs as references.

### Step 2: Create Message Types (messages.go)
1. Open `messages.go`
2. Add three new message struct types:
   - `BookmarkPost` with fields: `PostID string` and `Username string`
   - `UnbookmarkPost` with fields: `PostID string` and `Username string`
   - `GetBookmarks` with field: `Username string` (optional, for retrieving bookmarks)
3. These follow the same pattern as existing message types like `Vote` and `SendDirectMessage`

**Rationale**: Messages are the communication protocol in the actor model; these new types enable bookmark operations.

### Step 3: Implement Engine Logic (engine.go)
1. Open `engine.go`
2. Update the `Receive()` method to handle the three new message types
3. Add case statements for `*BookmarkPost`, `*UnbookmarkPost`, and `*GetBookmarks`
4. Implement three new methods:
   - `bookmarkPost(postID, username string)`: Adds postID to user's BookmarkedPosts if not already present
   - `unbookmarkPost(postID, username string)`: Removes postID from user's BookmarkedPosts if present
   - `getBookmarks(username string)`: Logs all bookmarked posts for a user
5. Follow the existing pattern: check if post and user exist, perform the action, log it
6. Use the existing `contains()` and `remove()` helper functions for slice operations
7. Update `getSimulationStats()` to display bookmark statistics

**Rationale**: The Engine actor is where all business logic lives; these implementations handle the core bookmark operations.

### Step 4: Integrate with Simulator (simulator.go)
1. Open `simulator.go`
2. Update `simulateAction()` method:
   - Change `rand.Intn(8)` to `rand.Intn(10)` to include two new action types
   - Add case 8 for `s.simulateBookmarkPost(context)`
   - Add case 9 for `s.simulateUnbookmarkPost(context)`
3. Implement `simulateBookmarkPost()` method:
   - Check if posts exist
   - Send a `BookmarkPost` message with a random post and random user
4. Implement `simulateUnbookmarkPost()` method:
   - Check if posts exist
   - Send an `UnbookmarkPost` message with a random post and random user

**Rationale**: The Simulator generates realistic user behavior; these additions make bookmarking a simulated user action.

### Step 5: Testing & Validation
1. Compile the project: `go build`
2. Run with small parameters to verify functionality:
   ```bash
   go run . -users=5 -subreddits=3 -actions=50 -time=10
   ```
3. Verify in output:
   - Bookmark/unbookmark actions appear in user action logs
   - Bookmark statistics display in final statistics
   - No errors or panics occur
4. Run with larger parameters to test scalability:
   ```bash
   go run . -users=100 -subreddits=20 -actions=1000 -time=30
   ```

### Step 6: Documentation Updates (Optional)
1. Update `README.md` to mention bookmarking feature
2. Add bookmark feature to the features list
3. Include example output showing bookmark statistics

---

## Technical Considerations

### 1. **Data Consistency**
- Bookmarks are stored as post IDs in the User struct
- No separate Post.Bookmarks field needed (can be computed if needed)
- The centralized Engine actor ensures consistency

### 2. **Performance**
- Using `contains()` and `remove()` for slice operations is O(n)
- For large bookmark lists, consider using a map[string]bool for O(1) lookups
- Current implementation is acceptable for simulation purposes

### 3. **Edge Cases Handled**
- Attempting to bookmark a non-existent post (silently ignored)
- Attempting to bookmark an already bookmarked post (prevented by contains check)
- Attempting to unbookmark a non-bookmarked post (prevented by contains check)
- Attempting to bookmark/unbookmark when user doesn't exist (silently ignored)

### 4. **Logging & Traceability**
- All bookmark actions are logged with timestamps via `logUserAction()`
- Logs include: username, action type, post ID, post author, and subreddit
- Enables detailed audit trails of user behavior

### 5. **Scalability Notes**
- Current implementation works well for simulations up to 100K users
- For production systems, consider:
  - Using a map for O(1) bookmark lookups: `BookmarkedPosts map[string]bool`
  - Implementing a separate BookmarkService actor for large-scale systems
  - Adding database persistence for bookmarks

---

## Potential Enhancements

### Future Improvements
1. **Bookmark Collections**: Allow users to organize bookmarks into named collections
2. **Bookmark Metadata**: Track when posts were bookmarked, add notes/tags
3. **Bookmark Sharing**: Allow users to share their bookmark collections
4. **Bookmark Notifications**: Notify users when bookmarked posts get new comments
5. **Trending Bookmarks**: Track most-bookmarked posts across the platform
6. **Bookmark Export**: Export bookmarks to various formats (JSON, CSV)

### Database Integration
When scaling to production:
```go
type BookmarkRecord struct {
	UserID      string
	PostID      string
	BookmarkedAt time.Time
	Notes       string
}
```

---

## Summary

This implementation adds post bookmarking functionality to the Reddit-clone simulator by:

- Extending the User model to track bookmarked posts
- Creating new message types for bookmark operations
- Implementing bookmark logic in the Engine actor
- Integrating with the Simulator for realistic behavior
- Adding statistics and logging for visibility
- Following existing patterns for consistency
- Maintaining performance for simulation scales

The changes are minimal, focused, and follow the established architecture patterns of the codebase. The implementation is backward-compatible and does not affect existing functionality.
