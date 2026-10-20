package main

import (
	"log"
	"os"

	"github.com/joonmo-you/aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	currentState := &state{
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.registeredCommands[cmdName](currentState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
}
