package main

import "users_orchestrator/cmd"

func main() {
	defer cmd.Stop()
	cmd.Start()
}
