package executor

import (
	"runtime"
	"strconv"
)

var printCmdMap = map[string][]string{
	"windows": {"cmd", "/C", "echo"},
	"*":       {"echo"},
}

var sleepCmdMap = map[string][]string{
	"windows": {"cmd", "/C", "ping", "1.1.1.1", "-n", "1", "-w"},
	"*":       {"sleep"},
}

func getOsCmd(cmds map[string][]string) (cmd []string) {
	cmd, ok := cmds[runtime.GOOS]

	if !ok {
		cmd = cmds["*"]
	}
	return
}

func getOsPrintCmd() []string {
	return getOsCmd(printCmdMap)
}

func getOsSleepCmd() []string {
	return getOsCmd(sleepCmdMap)
}

func getPrintCmd(echoText string) (cmdName string, cmdArgs []string) {
	printCmd := getOsPrintCmd()
	cmdName = printCmd[0]
	cmdArgs = append(printCmd[1:], echoText)
	return
}

func getSleepCmd(duration int) (cmdName string, cmdArgs []string) {
	if runtime.GOOS == "windows" {
		duration *= 1000
	}
	sleepCmd := getOsSleepCmd()
	cmdName = sleepCmd[0]
	cmdArgs = append(sleepCmd[1:], strconv.Itoa(duration))
	return
}

func newPrintCommand(echoText string) *Command {
	cmdName, cmdArgs := getPrintCmd(echoText)

	return NewCommand(cmdName, cmdArgs...)
}

func newSleepCommand(duration int) *Command {
	cmdName, cmdArgs := getSleepCmd(duration)

	return NewCommand(cmdName, cmdArgs...)
}
