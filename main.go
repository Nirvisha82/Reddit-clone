package main

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	system := actor.NewActorSystem()
	engineProps := actor.PropsFromProducer(func() actor.Actor { return NewEngine() })
	enginePID := system.Root.Spawn(engineProps)

	fmt.Println("Reddit engine started. Press Enter to stop...")
	fmt.Scanln()

	system.Root.Stop(enginePID)
	fmt.Println("Engine stopped.")
}
