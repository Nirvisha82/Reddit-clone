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
