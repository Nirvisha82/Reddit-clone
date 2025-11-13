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
