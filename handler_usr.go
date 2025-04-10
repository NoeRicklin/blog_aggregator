package main

import(
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 { return fmt.Errorf("<login> needs 1 Argument\n") }
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil { return err }

	fmt.Printf("Logged in as <%s>\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 { return fmt.Errorf("<register> needs 1 Argument\n") }
	username := cmd.args[0]
}
