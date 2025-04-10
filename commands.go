
package main

import(
	"fmt"
	"github.com/NoeRicklin/blog_aggregator/internal/config"
	"github.com/NoeRicklin/blog_aggregator/internal/database"
)

type state struct {
	cfg 	*config.Config
	db		*database.Queries
}

type command struct {
	name	string
	args	[]string
}

type commands struct {
	cmds 	map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.cmds[name]; ok {
		return fmt.Errorf("Command Handler already exists\n")
	}

	c.cmds[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Command doesn't exist\n")
	}

	return f(s, cmd)
}

