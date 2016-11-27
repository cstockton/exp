package executor

import (
	"strings"
	"testing"
)

func TestNewCommand(t *testing.T) {
	command := newPrintCommand("TestNewCommand")

	if command == nil {
		t.Error("got nill command")
	}
}

func TestNewCommandResult(t *testing.T) {
	result := NewCommandResult()

	if result == nil {
		t.Error("got nill command result")
	}
}

func TestCommandRun(t *testing.T) {
	echoText := "TestCommandRun"
	command := newPrintCommand(echoText)

	if command == nil {
		t.Error("got nill command")
	}

	result := command.Run()
	if strings.TrimSpace(result.Stdout) != echoText {
		t.Errorf("got stdout %q, stderr %q; want %q on stdout", result.Stdout, result.Stderr, echoText)
	}
}

func TestCommandRunGo(t *testing.T) {
	command := NewCommand("go", "help")

	if command == nil {
		t.Error("got nill command")
	}

	result := command.Run()
	if !strings.Contains(result.Stdout, "Go is a tool for managing Go source code") {
		t.Errorf("got stdout %q, stderr %q; wanted go help text on stdout", result.Stdout, result.Stderr)
	}
}
