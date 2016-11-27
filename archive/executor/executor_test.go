package executor

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func printResult(result CommandResult) {
	fmt.Println("Result-------------------")
	fmt.Println("Result Cmd:", result.Command)
	fmt.Println("Result Error:", result.Error)
	fmt.Println("Result Duration: ", result.Finished.Sub(result.Started))
	fmt.Println("Result Stdout:", result.Stdout)
	fmt.Println("Result Stderr:", result.Stderr)
	fmt.Println("Result ProcessState", result.Success)
}

func printResults(results []*CommandResult) {
	for _, result := range results {
		printResult(*result)
	}
}

func verifyNoError(t *testing.T, result *CommandResult) {
	if result.Error != nil {
		t.Fatalf("got err: %v; stdout %q, stderr %q", result.Error, result.Stdout, result.Stderr)
	}
}

func verifyNoErrors(t *testing.T, results []*CommandResult) {
	for _, res := range results {
		verifyNoError(t, res)
	}
}

func verifyEcho(t *testing.T, result *CommandResult) {
	verifyNoError(t, result)

	if strings.TrimSpace(result.Stdout) != result.Command.Arguments[len(result.Command.Arguments)-1] {
		t.Errorf("got stdout %q, stderr %q; want %q on stdout", result.Stdout, result.Stderr, result.Command.Arguments[1])
	}
}

func verifyEchos(t *testing.T, results []*CommandResult) {
	for _, res := range results {
		verifyEcho(t, res)
	}
}

func TestRun(t *testing.T) {
	echoText := "TestRun"
	cmdName, cmdArgs := getPrintCmd(echoText)
	res := Run(cmdName, cmdArgs...)

	if res.Error != nil {
		t.Fatalf("got err: %v; stdout %q, stderr %q", res.Error, res.Stdout, res.Stderr)
	}

	if strings.TrimSpace(res.Stdout) != echoText {
		t.Errorf("got stdout %q, stderr %q; want %q on stdout", res.Stdout, res.Stderr, res.Command.Arguments[1])
	}
}

func TestRunAll(t *testing.T) {
	commands := []*Command{}

	for i := 0; i < 20; i++ {
		commands = append(commands, newPrintCommand("TestRun"+strconv.Itoa(i)))
	}

	results := RunAll(commands...)

	verifyEchos(t, results)
}

func TestRunSequential(t *testing.T) {
	asserts := make(map[int]string)
	asserts[0] = ""

	commands := []*Command{}
	commands = append(commands, newSleepCommand(1))

	for i := 0; i < 20; i++ {
		asserts[i] = "TestRun" + strconv.Itoa(i)
		commands = append(commands, newPrintCommand(asserts[i]))
	}

	results := RunSequential(commands...)

	// sleep can exit with error on win with ping hack
	verifyNoErrors(t, results[1:])

	for i, res := range results[1:] {
		if strings.TrimSpace(res.Stdout) != asserts[i] {
			t.Errorf("got stdout %q, stderr %q; want %q on stdout", res.Stdout, res.Stderr, asserts[i])
		}
	}
}

func TestRunNotSequential(t *testing.T) {
	commands := []*Command{}
	commands = append(commands, newSleepCommand(2))

	for i := 0; i < 20; i++ {
		commands = append(commands, newPrintCommand("TestRun"+strconv.Itoa(i)))
	}

	results := RunAll(commands...)
	verifyNoErrors(t, results[:len(results)-1])

	// verify the tail of results is the ping by checking for TestRun in all other
	// results stdout
	for _, res := range results[:len(results)-1] {
		if strings.Contains("TestRun", res.Stdout) {
			t.Errorf("got stdout %q, stderr %q; wanted string.Contains(TestRun) on stdout", res.Stdout, res.Stderr)
		}
	}
}
