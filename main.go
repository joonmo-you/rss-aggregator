package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joonmo-you/aggregator/internal/config"
	"github.com/joonmo-you/aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	dbQueries := database.New(db)
	currentState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
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
