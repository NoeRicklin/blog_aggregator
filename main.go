package main

import(
	"log"
	"fmt"
	"github.com/NoeRicklin/blog_aggregator/internal/config"
)

const userName = "glueckskeks"

func main() {
	cfg, err := config.Read()
	if err != nil { log.Fatal(err) }

	s := state{ cfg: *cfg }

	c := commands{ cmds: make(map[string]func(*state, command) error }
}
