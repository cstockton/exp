package executor

// Run todo
func Run(command string, arguments ...string) CommandResult {
	cmd := NewCommand(command, arguments...)
	return cmd.Run()
}

// RunAll todo
func RunAll(commands ...*Command) []*CommandResult {
	count := len(commands)
	results := make([]*CommandResult, count)
	incoming := make(chan CommandResult, count)

	for _, cmd := range commands {
		go func(cmd *Command) {
			incoming <- cmd.Run()
		}(cmd)
	}

	for i := 0; i < count; i++ {
		result := <-incoming
		results[i] = &result
	}

	return results
}

// RunSequential todo
func RunSequential(commands ...*Command) []*CommandResult {
	count := len(commands)
	results := make([]*CommandResult, count)

	for i, cmd := range commands {
		result := cmd.Run()
		results[i] = &result
	}

	return results
}
