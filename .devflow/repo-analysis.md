This repository, "Reddit-clone," is an impressive demonstration of a high-performance, concurrent social platform simulator built in Go, leveraging the Proto Actor framework. It's designed to model realistic user behavior and large-scale interactions, focusing on the backend operations of a Reddit-like system.

## Repository Overview

1.  **Project Type**: This project is a **CLI (Command Line Interface) simulation tool**. It is not a live, interactive web application but rather a backend simulator designed to test the performance and behavioral aspects of an actor-based architecture under various load conditions. Its primary output is console logs and aggregated statistics of the simulated environment.

2.  **Architecture**: The core architecture is **actor-based, concurrent, and event-driven**.
    *   It utilizes the **Proto Actor framework** (`github.com/asynkron/protoactor-go`) to manage concurrent entities (actors) that communicate via asynchronous message passing.
    *   The system comprises two main actors: an `Engine` actor that encapsulates all business logic and state, and a `Simulator` actor that generates user actions and drives the simulation.
    *   User behavior, particularly subreddit popularity, is modeled using a **Zipf distribution** (`math/rand.Zipf`) to create a realistic "long tail" effect, where a few subreddits are highly popular, and many are less so.
    *   The entire simulation runs in-memory within a single Go process, making it an efficient tool for performance benchmarking on a single machine.

3.  **Technology Stack**:
    *   **Language**: Go (specifically Go 1.23 as per `go.mod`, though the `README.md` states 1.19+).
    *   **Concurrency Framework**: Proto Actor (`github.com/asynkron/protoactor-go`).
    *   **Concurrency Model**: Actor-based messaging.
    *   **Distribution Modeling**: Zipf distribution (`math/rand.Zipf`) for realistic simulation of popularity.
    *   **Architecture Style**: Concurrent, event-driven.

4.  **Entry Points**:
    *   The main entry point for the application is `main.go`.
    *   Execution begins with the `main()` function, which parses command-line flags to configure simulation parameters (number of users, subreddits, actions, and duration).
    *   It then initializes the Proto Actor system and spawns the `Engine` and `Simulator` actors, initiating the simulation.

## File Analysis

### File: README.md

1.  **Purpose**: Provides comprehensive documentation for the project.
2.  **Role**: Serves as the primary source of information for understanding the project's goals, features, technical stack, architecture, installation, usage, and performance characteristics. It's crucial for onboarding new developers and users, offering a high-level overview and detailed explanations of core concepts like the Karma System and Zipf Distribution.
3.  **Key Functions/Classes**: N/A (documentation file).
4.  **Dependencies**: N/A.
5.  **Business Logic**: N/A.

### File: main.go

1.  **Purpose**: The application's entry point and configuration manager.
2.  **Role**: Orchestrates the setup and teardown of the simulation environment. It's responsible for:
    *   Parsing command-line arguments (`-users`, `-subreddits`, `-actions`, `-time`) to customize simulation parameters.
    *   Initializing the Proto Actor system (`actor.NewActorSystem()`).
    *   Spawning the `Engine` and `Simulator` actors (`system.Root.Spawn(props)`), which are the core components of the simulation.
    *   Controlling the overall duration of the simulation using `time.Sleep()`.
    *   Stopping the actors gracefully at the end (`system.Root.Stop()`).
3.  **Key Functions/Classes**:
    *   `main()`: The primary function that executes when the program starts.
    *   `flag.Int()` and `flag.Parse()`: Used for command-line argument parsing.
    *   `actor.NewActorSystem()`: Creates the global actor system context.
    *   `actor.PropsFromProducer()`: Defines how new instances of `Engine` and `Simulator` actors should be created.
    *   `system.Root.Spawn()`: Initiates the `Engine` and `Simulator` actors.
4.  **Dependencies**: `flag`, `fmt`, `math/rand`, `time`, `github.com/asynkron/protoactor-go/actor`. It implicitly depends on `engine.go` and `simulator.go` by creating instances of `NewEngine()` and `NewSimulator()`.
5.  **Business Logic**: Contains no direct Reddit-like business logic itself, but it sets the stage for the simulation and its parameters.

### File: engine.go

1.  **Purpose**: Implements the core business logic and state management for the Reddit-like platform.
2.  **Role**: Acts as the central "brain" of the simulated Reddit system. It's an actor that receives messages representing user actions (e.g., register, post, vote, comment, direct message) and updates its internal state accordingly. It maintains all the data for users, subreddits, and posts in memory.
3.  **Key Functions/Classes**:
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
4.  **Dependencies**: `fmt`, `sort`, `strings`, `time`, `github.com/asynkron/protoactor-go/actor`. It heavily relies on data structures defined in `models.go` and message types from `messages.go`.
5.  **Business Logic**: Implements all the core rules of a Reddit-like platform: user registration, subreddit creation/joining/leaving, content creation (posts, comments, nested comments), voting mechanics (including karma calculation), direct messaging, and personalized content feeds.

### File: simulator.go

1.  **Purpose**: Generates realistic user behavior patterns and drives the simulation by sending messages to the `Engine` actor.
2.  **Role**: Acts as the "user base" of the simulation. It's an actor that continuously generates random actions (registering users, creating subreddits, posting, commenting, voting, etc.) based on predefined probabilities and distributions. It coordinates with the `Engine` actor by sending it the appropriate messages. It also manages the simulation lifecycle and triggers the final reporting.
3.  **Key Functions/Classes**:
    *   `Simulator` struct: Holds the `enginePID` (the address of the `Engine` actor), lists of `users`, `subreddits`, `posts`, `comments`, `userStatus` (online/offline), action counters, and a `rand.Zipf` distribution generator.
    *   `NewSimulator()`: Constructor for the `Simulator` struct, initializing the Zipf distribution (`rand.NewZipf`).
    *   `Receive(context actor.Context)`: The main message processing method for the `Simulator` actor. Upon `actor.Started`, it kicks off the `runSimulation` goroutine.
    *   `runSimulation(context actor.Context)`: The main loop of the simulation. It registers initial users and subreddits, then repeatedly calls `simulateAction` until the `SIMULATION_ACTIONS` limit is reached or the simulation time elapses.
    *   `simulateAction(context actor.Context)`: Randomly selects one of the various simulation actions (e.g., join subreddit, create post, vote) and calls the corresponding helper function.
    *   `simulateRegisterUser()`, `simulateCreateSubreddit()`, `simulateJoinSubreddit()`, `simulateLeaveSubreddit()`, `simulateCreatePost()`, `simulateCreateComment()`, `simulateVote()`, `simulateSendDirectMessage()`, `simulateGetFeed()`, `simulateConnection()`: These methods generate specific action messages with random data and send them to the `enginePID`. `simulateCreatePost` uses the Zipf distribution to select a subreddit, ensuring popular subreddits receive more content.
    *   `randomUser()`, `randomSubreddit()`, `randomPost()`, `randomComment()`: Helper functions to randomly select existing entities for actions.
    *   `printSimulationStats()`, `printUserActions()`: Sends messages to the `Engine` to trigger the printing of final statistics and user action logs.
4.  **Dependencies**: `fmt`, `math/rand`, `time`, `github.com/asynkron/protoactor-go/actor`. It depends on message types defined in `messages.go`.
5.  **Business Logic**: Implements the simulation's rules for user behavior, including the probabilities of different actions, the application of Zipf distribution for subreddit popularity, and the overall flow and timing of the simulation.

### File: models.go

1.  **Purpose**: Defines the data structures (models) that represent the entities within the Reddit-like platform.
2.  **Role**: Provides the blueprint for the state managed by the `Engine` actor. These structs (`User`, `Subreddit`, `Post`, `Comment`, `DirectMessage`) encapsulate the attributes and relationships of the simulated entities, ensuring a consistent data model across the system.
3.  **Key Functions/Classes**: N/A (pure data structs).
4.  **Dependencies**: N/A.
5.  **Business Logic**: N/A (defines the *nouns* or *data* that the business logic in `engine.go` operates on).

### File: messages.go

1.  **Purpose**: Defines the various message types (structs) used for communication between actors.
2.  **Role**: Establishes the communication protocol for the actor system. Each struct represents a distinct command or event that can be sent to an actor, primarily from the `Simulator` to the `Engine`. This clear definition of messages ensures type safety and a well-defined interface for actor interactions.
3.  **Key Functions/Classes**: N/A (pure data structs).
4.  **Dependencies**: `time` (for the `Timestamp` field in `UserAction`).
5.  **Business Logic**: N/A (defines the *verbs* or *actions* that trigger business logic in the `Engine`).

### File: go.mod

1.  **Purpose**: Manages the Go module dependencies for the project.
2.  **Role**: Defines the module path (`reddit-clone`), specifies the minimum Go version required (Go 1.23), and lists all direct and transitive dependencies. This file ensures that the project builds with the correct versions of its external libraries, primarily `github.com/asynkron/protoactor-go`.
3.  **Key Functions/Classes**: N/A (configuration file).
4.  **Dependencies**: N/A.
5.  **Business Logic**: N/A.

## System Relationships

1.  **Data Flow**:
    *   **Initialization**: `main.go` starts the `ActorSystem` and spawns the `Engine` and `Simulator` actors.
    *   **Simulation Loop**: The `Simulator` actor, running in its own goroutine, continuously generates various action messages (e.g., `RegisterUser`, `CreatePost`, `Vote`) using random logic and Zipf distribution.
    *   **Message Passing**: The `Simulator` sends these action messages to the `Engine` actor's PID (`context.Send(s.enginePID, message)`).
    *   **State Update**: The `Engine` actor receives these messages in its `Receive` method. Based on the message type, it calls the appropriate handler function (e.g., `createPost`, `vote`). These handlers modify the `Engine`'s internal state (maps holding `models.go` structs like `User`, `Post`, `Subreddit`).
    *   **Logging**: The `Engine` also logs individual user actions using `logUserAction` for detailed output.
    *   **Reporting**: At the end of the simulation, the `Simulator` sends `GetSimulationStats` and `PrintUserActions` messages to the `Engine`. The `Engine` then processes these to print aggregated statistics and detailed user action logs to the console.

2.  **Key Components**:
    *   **Proto Actor System**: The underlying framework that enables concurrent, message-driven communication, providing the runtime for actors.
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
    *   **Organization**: The project exhibits good modularity with clear separation of concerns into different Go files (`main`, `engine`, `simulator`, `models`, `messages`). This makes the codebase easy to navigate and understand. The `README.md` is exceptionally well-written and provides excellent context.
    *   **Readability**: The code is generally clean, well-structured, and uses descriptive variable and function names. The use of `fmt.Printf` for logging is straightforward for a CLI tool.
    *   **Consistency**: The use of the Proto Actor framework is consistent throughout the actor-related files.
    *   **Error Handling**: Explicit error handling (e.g., checking if a user exists before performing an action) is present in `engine.go` (e.g., `if _, exists := e.users[username]; !exists`). However, the system primarily assumes valid messages from the `Simulator`. For a production system, more robust validation and error propagation would be necessary. For a simulation, this level is often acceptable.
    *   **Logging**: The custom `logUserAction` in `engine.go` is a good approach for capturing simulation events in a structured way, complementing the direct `fmt.Printf` statements.

2.  **Design Patterns**:
    *   **Actor Model**: This is the foundational design pattern. It promotes concurrency, isolation, and message-driven communication, which are well-suited for simulating independent entities and their interactions.
    *   **Command Pattern**: The message structs in `messages.go` effectively act as commands that are sent to the `Engine` actor, which then executes the corresponding business logic.
    *   **Centralized State (within an Actor)**: The `Engine` actor centralizes all the application's state (users, subreddits, posts). This ensures consistency and simplifies state management within the single actor.

3.  **Potential Issues**:
    *   **Single `Engine` Actor Bottleneck**: For a truly massive, real-world Reddit-like system, a single `Engine` actor managing *all* state could become a performance bottleneck. While Proto Actor supports distribution, this implementation keeps all core state and logic within one actor. For the stated simulation scales (up to 100K users, 250K actions), it performs well, but beyond that, a more granular distribution of state (e.g., `SubredditActor` for each subreddit, `UserActor` for each user) would be necessary.
    *   **No Persistence**: As a simulator, persistence is not a requirement. However, if this project were to evolve into a live application, integrating a database or persistent storage would be a major undertaking.
    *   **Limited Input Validation**: While the `Engine` checks for existence of entities (e.g., `if _, exists := e.users[username]`), it doesn't extensively validate the *content* of messages (e.g., malformed post titles). For a simulation, this is often acceptable, but critical for production.
    *   **Global Random Seed**: `rand.Seed(time.Now().UnixNano())` is called once in `main.go`. While the `Simulator` uses its own `rand.New` instance for Zipf, for highly reproducible simulation runs, a fixed seed might be preferable as a command-line option to allow for exact replication of simulation scenarios.
    *   **`simulateJoinSubreddit` Logic**: The current `simulateJoinSubreddit` iterates through all subreddits and for each, attempts to add `memberCount` (from Zipf) random users. This could lead to a very high number of `JoinSubreddit` messages, potentially overwhelming the `Engine` or creating an unrealistic join rate. A more balanced approach might be to pick a random user and a random subreddit (using Zipf for subreddit selection) and have that user join.

4.  **Scalability**:
    *   **Vertical Scaling**: The current design leverages Go's concurrency and the actor model effectively on a single machine. The performance benchmarks provided in the README (up to 100,000 users and 250,000 actions in under 3 minutes on an M3 Pro) demonstrate good vertical scalability and efficient utilization of multi-core processors. The `Engine` actor processes messages sequentially, preventing internal race conditions, but the overall system can handle high message throughput.
    *   **Horizontal Scaling**: The Proto Actor framework is inherently designed for distributed systems. While the current implementation centralizes the `Engine` actor, the architecture *could* be extended for horizontal scaling. This would involve:
        *   Breaking down the `Engine` into smaller, more specialized actors (e.g., `SubredditManagerActor`, `UserManagerActor`, or even individual actors per subreddit/user).
        *   Implementing consistent hashing or other distribution strategies to place these actors across multiple nodes.
        *   Introducing more complex message routing and potentially persistence mechanisms.
        *   This would be a significant architectural evolution, but the foundation provided by Proto Actor is suitable for such an expansion.

5.  **Maintainability**:
    *   **Modular Design**: The clear separation of concerns into distinct files and actors greatly enhances maintainability. Changes to data structures (`models.go`) are isolated, and changes to business logic (`engine.go`) are contained within its `Receive` method and associated handlers.
    *   **Clear Interfaces**: The `messages.go` file explicitly defines the communication contracts between actors, making it easy to understand what actions can be performed and what data is exchanged.
    *   **Readability**: The code is generally easy to read and follow, with well-named functions and variables.
    *   **Testability**: While not explicitly shown, the actor model with clear message types lends itself well to unit and integration testing by sending specific messages to actors and asserting their state changes or outgoing messages.
    *   **Extensibility**: Adding new features (e.g., new user actions, new content types) would primarily involve defining new message types, adding a new case to the `Engine`'s `Receive` method, implementing the corresponding logic, and adding a new simulation function in `simulator.go`. This is a straightforward process due to the modular design.