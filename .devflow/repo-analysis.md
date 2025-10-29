This repository, "Reddit-clone," presents a fascinating implementation of a Reddit-like social platform simulator built in Go, leveraging the Proto Actor framework. It focuses on demonstrating a high-performance, concurrent system capable of modeling realistic user behavior and large-scale interactions.

## Repository Overview

1.  **Project Type**: This project is a **CLI (Command Line Interface) simulation tool**. It's designed to simulate the backend operations of a social media platform like Reddit, rather than being a live, interactive web application or a general-purpose library. Its primary goal is to test the performance and behavior of an actor-based architecture under various load conditions.

2.  **Architecture**: The core architecture is **actor-based, concurrent, and event-driven**.
    *   It utilizes the **Proto Actor framework** to manage concurrent entities (actors) that communicate via asynchronous message passing.
    *   The system is composed of two main actors: an `Engine` actor that encapsulates all business logic and state, and a `Simulator` actor that generates user actions and drives the simulation.
    *   User behavior, particularly subreddit popularity, is modeled using a **Zipf distribution** to create a realistic "long tail" effect.
    *   The entire simulation runs in-memory within a single Go process, making it a powerful tool for performance benchmarking on a single machine.

3.  **Technology Stack**:
    *   **Language**: Go (specifically Go 1.19+ as per README, though `go.mod` indicates Go 1.23).
    *   **Concurrency Framework**: Proto Actor (`github.com/asynkron/protoactor-go`).
    *   **Concurrency Model**: Actor-based messaging.
    *   **Distribution Modeling**: Zipf distribution (`math/rand.Zipf`) for realistic simulation of popularity.
    *   **Architecture Style**: Concurrent, event-driven.

4.  **Entry Points**:
    *   The main entry point for the application is `main.go`.
    *   Execution starts with the `main()` function, which parses command-line flags to configure simulation parameters (number of users, subreddits, actions, and duration).
    *   It then initializes the Proto Actor system and spawns the `Engine` and `Simulator` actors, kicking off the simulation.

## File Analysis

### File: README.md

*   **Purpose**: Provides comprehensive documentation for the project.
*   **Role**: Serves as the primary source of information for understanding the project's goals, features, technical stack, architecture, installation, usage, and performance characteristics. It's crucial for onboarding new developers and users.
*   **Key Functions/Classes**: N/A (documentation file).
*   **Dependencies**: N/A.
*   **Business Logic**: N/A.

### File: main.go

*   **Purpose**: The application's entry point and configuration manager.
*   **Role**: Orchestrates the setup and teardown of the simulation environment. It's responsible for:
    *   Parsing command-line arguments to customize simulation parameters.
    *   Initializing the Proto Actor system.
    *   Spawning the `Engine` and `Simulator` actors, which are the core components of the simulation.
    *   Controlling the overall duration of the simulation.
    *   Stopping the actors gracefully at the end.
*   **Key Functions/Classes**:
    *   `main()`: The primary function that executes when the program starts. It uses the `flag` package to define and parse command-line options like `-users`, `-subreddits`, `-actions`, and `-time`.
    *   `actor.NewActorSystem()`: Creates the global actor system context.
    *   `actor.PropsFromProducer()`: Used to define how new instances of `Engine` and `Simulator` actors should be created.
    *   `system.Root.Spawn()`: Initiates the `Engine` and `Simulator` actors within the actor system.
    *   `time.Sleep()`: Pauses the main goroutine for the specified simulation duration.
    *   `system.Root.Stop()`: Sends stop signals to the spawned actors.
*   **Dependencies**: `flag`, `fmt`, `math/rand`, `time`, `github.com/asynkron/protoactor-go/actor`. It implicitly depends on `engine.go` and `simulator.go` by creating instances of `NewEngine()` and `NewSimulator()`.
*   **Business Logic**: Contains no direct Reddit-like business logic itself, but it sets the stage for the simulation and its parameters.

### File: engine.go

*   **Purpose**: Implements the core business logic and state management for the Reddit-like platform.
*   **Role**: Acts as the central "brain" of the simulated Reddit system. It's an actor that receives messages representing user actions (e.g., register, post, vote) and updates its internal state accordingly. It maintains all the data for users, subreddits, and posts.
*   **Key Functions/Classes**:
    *   `Engine` struct: Holds the entire state of the simulated Reddit platform using maps: `users`, `subreddits`, `posts`, and `userActions`.
    *   `NewEngine()`: Constructor for the `Engine` struct.
    *   `Receive(context actor.Context)`: The main message processing method for the `Engine` actor. It uses a `switch` statement to dispatch incoming messages (defined in `messages.go`) to specific handler methods.
    *   `registerUser(username string)`: Creates a new `User` entry.
    *   `createSubreddit(name, creator string)`: Creates a new `Subreddit` and adds the creator as a member.
    *   `joinSubreddit(subredditName, username string)`: Adds a user to a subreddit's member list and updates the user's subscriptions.
    *   `leaveSubreddit(subredditName, username string)`: Removes a user from a subreddit.
    *   `createPost(postID, subredditName, author, title, content string)`: Creates a new `Post`, automatically upvotes it by the author, and increments the author's karma.
    *   `createComment(postID, parentID, commentID, author, content string)`: Adds a new comment, supporting nested replies. Increments the author's karma.
    *   `addChildComment(comments []*Comment, parentID string, author string, newComment *Comment)`: Recursively finds the parent comment to add a reply.
    *   `vote(postID, userID string, isUpvote bool)`: Updates post upvote/downvote counts and the author's karma based on the vote.
    *   `sendDirectMessage(from, to, content string)`: Stores direct messages in both sender's sent messages and receiver's received messages.
    *   `getFeed(username string)`: Generates a personalized feed for a user based on their subscribed subreddits.
    *   `logUserAction(username, action string)`: Records a timestamped action for a user, used for detailed simulation output.
    *   `getSimulationStats()`, `printAllUserActions()`, `printSubredditPostsAndComments()`, `printComments()`: Methods for generating and printing various statistics and detailed logs at the end of the simulation.
    *   `contains()`, `remove()`: Helper functions for slice manipulation.
*   **Dependencies**: `fmt`, `sort`, `strings`, `time`, `github.com/asynkron/protoactor-go/actor`. It heavily relies on data structures defined in `models.go` and message types from `messages.go`.
*   **Business Logic**: Implements all the core rules of a Reddit-like platform: user registration, subreddit management, content creation (posts, comments), voting mechanics (including karma calculation), direct messaging, and personalized content feeds.

### File: go.mod

*   **Purpose**: Manages the Go module dependencies for the project.
*   **Role**: Defines the module path (`reddit-clone`), specifies the minimum Go version required (Go 1.23), and lists all direct and transitive dependencies. This file ensures that the project builds with the correct versions of its external libraries, primarily `github.com/asynkron/protoactor-go`.
*   **Key Functions/Classes**: N/A (configuration file).
*   **Dependencies**: N/A.
*   **Business Logic**: N/A.

### File: messages.go

*   **Purpose**: Defines the various message types (structs) used for communication between actors.
*   **Role**: Establishes the communication protocol for the actor system. Each struct represents a distinct command or event that can be sent to an actor, primarily from the `Simulator` to the `Engine`.
*   **Key Functions/Classes**: N/A (pure data structs).
*   **Dependencies**: `time` (for the `Timestamp` field in `UserAction`).
*   **Business Logic**: N/A (defines the *verbs* or *actions* that trigger business logic in the `Engine`).

### File: models.go

*   **Purpose**: Defines the data structures (models) that represent the entities within the Reddit-like platform.
*   **Role**: Provides the blueprint for the state managed by the `Engine` actor. These structs (`User`, `Subreddit`, `Post`, `Comment`, `DirectMessage`) encapsulate the attributes and relationships of the simulated entities.
*   **Key Functions/Classes**: N/A (pure data structs).
*   **Dependencies**: N/A.
*   **Business Logic**: N/A (defines the *nouns* or *data* that the business logic operates on).

### File: simulator.go

*   **Purpose**: Generates realistic user behavior patterns and drives the simulation by sending messages to the `Engine` actor.
*   **Role**: Acts as the "user base" of the simulation. It's an actor that continuously generates random actions (registering users, creating subreddits, posting, commenting, voting, etc.) based on predefined probabilities and distributions. It coordinates with the `Engine` actor by sending it the appropriate messages.
*   **Key Functions/Classes**:
    *   `Simulator` struct: Holds the `enginePID` (the address of the `Engine` actor), lists of `users`, `subreddits`, `posts`, `comments`, `userStatus` (online/offline), action counters, and a `rand.Zipf` distribution generator.
    *   `NewSimulator()`: Constructor for the `Simulator` struct, initializing the Zipf distribution.
    *   `Receive(context actor.Context)`: The main message processing method for the `Simulator` actor. Upon `actor.Started`, it kicks off the `runSimulation` goroutine.
    *   `runSimulation(context actor.Context)`: The main loop of the simulation. It registers initial users and subreddits, then repeatedly calls `simulateAction` until the `SIMULATION_ACTIONS` limit is reached.
    *   `simulateAction(context actor.Context)`: Randomly selects one of the various simulation actions (e.g., join subreddit, create post, vote) and calls the corresponding helper function.
    *   `simulateRegisterUser()`, `simulateCreateSubreddit()`, `simulateJoinSubreddit()`, `simulateLeaveSubreddit()`, `simulateCreatePost()`, `simulateCreateComment()`, `simulateVote()`, `simulateSendDirectMessage()`, `simulateGetFeed()`, `simulateConnection()`: These methods generate specific action messages with random data and send them to the `enginePID`. `simulateCreatePost` uses the Zipf distribution to select a subreddit.
    *   `randomUser()`, `randomSubreddit()`, `randomPost()`, `randomComment()`: Helper functions to randomly select existing entities for actions.
    *   `printSimulationStats()`, `printUserActions()`: Sends messages to the `Engine` to trigger the printing of final statistics and user action logs.
*   **Dependencies**: `fmt`, `math/rand`, `time`, `github.com/asynkron/protoactor-go/actor`. It depends on message types defined in `messages.go`.
*   **Business Logic**: Implements the simulation's rules for user behavior, including the probabilities of different actions, the application of Zipf distribution for subreddit popularity, and the overall flow of the simulation.

## System Relationships

1.  **Data Flow**:
    *   **Initialization**: `main.go` starts the `ActorSystem` and spawns the `Engine` and `Simulator` actors.
    *   **Simulation Loop**: The `Simulator` actor, running in its own goroutine, continuously generates various action messages (e.g., `RegisterUser`, `CreatePost`, `Vote`) using random logic and Zipf distribution.
    *   **Message Passing**: The `Simulator` sends these action messages to the `Engine` actor's PID (`context.Send(s.enginePID, message)`).
    *   **State Update**: The `Engine` actor receives these messages in its `Receive` method. Based on the message type, it calls the appropriate handler function (e.g., `createPost`, `vote`). These handlers modify the `Engine`'s internal state (maps holding `models.go` structs like `User`, `Post`, `Subreddit`).
    *   **Logging**: The `Engine` also logs individual user actions using `logUserAction` for detailed output.
    *   **Reporting**: At the end of the simulation, the `Simulator` sends `GetSimulationStats` and `PrintUserActions` messages to the `Engine`. The `Engine` then processes these to print aggregated statistics and detailed user action logs to the console.

2.  **Key Components**:
    *   **Proto Actor System**: The underlying framework that enables concurrent, message-driven communication.
    *   **`Engine` Actor**: The central repository for all application state and business logic. It's the single source of truth for the simulated Reddit data.
    *   **`Simulator` Actor**: The driver of the simulation, responsible for generating realistic user interactions and feeding them to the `Engine`.
    *   **`Models` (in `models.go`)**: The data structures that define the entities within the Reddit-like platform (users, posts, subreddits, comments, messages).
    *   **`Messages` (in `messages.go`)**: The communication protocol, defining the types of interactions and data exchanged between actors.

3.  **Integration Points**:
    *   `main.go` integrates with `engine.go` and `simulator.go` by creating and managing their actor instances.
    *   `simulator.go` integrates directly with `engine.go` by sending messages to its `actor.PID`. This is the primary communication channel.
    *   Both `engine.go` and `simulator.go` rely on `models.go` for defining the structure of the data they manipulate and `messages.go` for the types of messages they send and receive.

4.  **API/Interface Design**:
    *   The system uses an **actor-based messaging interface**. Components communicate by sending Go structs (defined in `messages.go`) to the `actor.PID` of the target actor.
    *   The `Engine` actor effectively exposes a "message-based API" through its `Receive` method, where each message type (`*RegisterUser`, `*CreatePost`, etc.) acts as a distinct endpoint or command.
    *   There are no traditional external APIs (like REST or gRPC) as this is an in-memory simulation.

## Development Insights

1.  **Code Quality**:
    *   **Organization**: The project exhibits good modularity with clear separation of concerns into different Go files (`main`, `engine`, `simulator`, `models`, `messages`). This makes the codebase easy to navigate and understand.
    *   **Readability**: The code is generally clean, well-structured, and uses descriptive variable and function names. The `README.md` is exceptionally well-written and provides excellent context.
    *   **Consistency**: The use of the Proto Actor framework is consistent throughout the actor-related files.
    *   **Error Handling**: Explicit error handling is minimal, which is common and acceptable for a simulation project where the focus is on functional behavior and performance rather than robust error recovery. For a production system, this would need significant enhancement.
    *   **Logging**: The custom `logUserAction` in `engine.go` is a good approach for capturing simulation events in a structured way, complementing the `fmt.Printf` statements.

2.  **Design Patterns**:
    *   **Actor Model**: This is the foundational design pattern. It promotes concurrency, isolation, and message-driven communication, which are well-suited for simulating independent entities and their interactions.
    *   **Command Pattern**: The message structs in `messages.go` effectively act as commands that are sent to the `Engine` actor, which then executes the corresponding business logic.
    *   **Centralized State (within an Actor)**: The `Engine` actor centralizes all the application's state (users, subreddits, posts). While this is a common approach within a single actor to ensure consistency, it could be a bottleneck in a truly distributed, high-throughput scenario if not carefully managed.

3.  **Potential Issues**:
    *   **Single `Engine` Actor Bottleneck**: For a truly massive, real-world Reddit-like system, a single `Engine` actor managing *all* state could become a performance bottleneck. While Proto Actor supports distribution, this implementation keeps all core state and logic within one actor. For the stated simulation scales (100K users, 250K actions), it performs well, but beyond that, a more granular distribution of state (e.g., `SubredditActor` for each subreddit, `UserActor` for each user) would be necessary.
    *   **No Persistence**: As a simulator, persistence is not a requirement. However, if this project were to evolve into a live application, integrating a database or persistent storage would be a major undertaking.
    *   **Limited Error Handling**: The lack of explicit error handling (e.g., what if a user tries to join a non-existent subreddit?) means the simulation might proceed with invalid states if the `Simulator` were to send malformed messages. For a simulation, this is often overlooked, but critical for production.
    *   **Global Random Seed**: `rand.Seed(time.Now().UnixNano())` is called once in `main.go`. While the `Simulator` uses its own `rand.New` instance for Zipf, for highly reproducible simulation runs, a fixed seed might be preferable as a command-line option.

4.  **Scalability**:
    *   **Vertical Scaling**: The current design leverages Go's concurrency and the actor model effectively on a single machine. The performance benchmarks provided in the README (up to 100,000 users and 250,000 actions in under 3 minutes on an M3 Pro) demonstrate good vertical scalability and efficient utilization of multi-core processors. The `Engine` actor processes messages sequentially, preventing internal race conditions, but the overall system can handle high message throughput.
    *   **Horizontal Scaling**: The Proto Actor framework is inherently designed for distributed systems. While the current implementation centralizes the `Engine` actor, the architecture *could* be extended for horizontal scaling. This would involve:
        *   Breaking down the `Engine` into smaller, more specialized actors (e.g., `SubredditManagerActor`, `UserManagerActor`, or even individual actors per subreddit/user).
        *   Implementing consistent hashing or other distribution strategies to place these actors across multiple nodes.
        *   Introducing more complex message routing.
        *   This would be a significant architectural evolution, but the foundation is there.

5.  **Maintainability**:
    *   **Modular Design**: The clear separation of concerns into distinct files and actors greatly enhances maintainability. Changes to data structures (`models.go`) are isolated, and changes to business logic (`engine.go`) are contained within its