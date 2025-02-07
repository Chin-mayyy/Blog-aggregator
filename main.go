package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Chin-mayyy/Blog_aggregator/internal/config"
	"github.com/Chin-mayyy/Blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type State struct {
	cfg *config.Config
	db  database.Queries
}

func main() {
	cfg, err := config.ReadFile()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", "postgres://chinmay:postgres@localhost:5432/gator?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &State{
		db:  *dbQueries,
		cfg: &cfg,
	}

	cmds := Commands{
		registeredCommand: make(map[string]func(*State, Command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", GetUsers)
	cmds.register("agg", agg)
	cmds.register("addfeed", middlewareLoggedIn(addFeed))
	cmds.register("feeds", feeds)
	cmds.register("follow", follow)
	cmds.register("following", following)
	cmds.register("unfollow", middlewareLoggedIn(unfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArguments := os.Args[2:]

	err = cmds.run(programState, Command{command_name: cmdName, arguments: cmdArguments})
	if err != nil {
		log.Fatal(err)
	}
}
