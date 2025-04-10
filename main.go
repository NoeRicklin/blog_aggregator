package main

import _ "github.com/lib/pq"
import (
	"log"
	"database/sql"
	"os"
	"github.com/NoeRicklin/blog_aggregator/internal/config"
	"github.com/NoeRicklin/blog_aggregator/internal/database"
)

const userName = "glueckskeks"

func main() {
	cfg, err := config.Read()
	if err != nil { log.Fatal(err) }

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil { log.Fatal(err) }

	dbQueries := database.New(db)

	s := &state{ cfg: 	&cfg,
				db:		dbQueries }

	c := commands{ cmds: make(map[string]func(*state, command) error) }

	c.register("login", handlerLogin)

	input := os.Args
	if len(input) < 2 { log.Fatal("Requires arguments") }

	cmd := command{ name: input[1], args: input[2:] }

	err = c.run(s, cmd)
	if err != nil { log.Fatal(err) }
}
