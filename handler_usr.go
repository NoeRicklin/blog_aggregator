package main

import(
	"fmt"
	"github.com/NoeRicklin/blog_aggregator/internal/config"
)

type state struct {
	cfg 	*config.Config
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 { return fmt.Errorf("<login> needs 1 Argument\n") }
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil { return err }

	fmt.Printf("Logged in as <%s>\n", username)
	return nil
}

