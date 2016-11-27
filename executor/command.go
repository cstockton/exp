package executor

import (
	"bytes"
	"os/exec"
	"time"
)

const (

	// DefaultCommandTimeout todo
	DefaultCommandTimeout = 30 * time.Second
)

// CommandTimeout todo
var CommandTimeout = DefaultCommandTimeout

// CommandResult todo
type CommandResult struct {
	Command  *Command
	Stdout   string
	Stderr   string
	Success  bool
	Error    error
	Started  time.Time
	Finished time.Time
}

// NewCommandResult todo
func NewCommandResult() *CommandResult {
	return &CommandResult{
		Command: nil,
		Stdout:  "",
		Stderr:  "",
		Error:   nil,
		Success: false,
	}
}

// Command todo
type Command struct {
	Path      string
	Arguments []string
	Timeout   time.Duration
}

// NewCommand todo
func NewCommand(command string, arguments ...string) *Command {
	return &Command{
		Path:      command,
		Arguments: arguments,
		Timeout:   CommandTimeout,
	}
}

// Run todo
func (c *Command) Run() CommandResult {
	result := *NewCommandResult()
	cmd := exec.Command(c.Path, c.Arguments...)
	done := make(chan bool, 1)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	result.Command = c
	result.Started = time.Now()

	if err := cmd.Start(); err != nil {
		result.Error = err
		return result
	}

	go func() {
		err := cmd.Wait()
		result.Finished = time.Now()

		if err != nil {
			result.Error = err
		}

		result.Stdout = stdout.String()
		result.Stderr = stderr.String()
		result.Success = cmd.ProcessState.Success()

		done <- true
	}()

	select {
	case <-time.After(c.Timeout):
		return result
	case <-done:
		return result
	}
}
