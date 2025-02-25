package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// confirm will prompt the user with the given toPrint string, and
// exit the program if N or n is input.
func confirm(p ...interface{}) {
	if yesEnabled {
		return
	}
	toPrint := fmt.Sprint(p...)
	toPrint = printer(color.FgYellow, "CONF", toPrint)
	fmt.Print(toPrint + " [Y/n]: ")
	var resp string
	fmt.Scanln(&resp)
	if strings.ToLower(strings.TrimSpace(resp)) == "n" {
		os.Exit(1)
	}
}

// ask will prompt the user with the given toPrint string, and
// return a boolean.
func ask(p ...interface{}) bool {
	if yesEnabled {
		return true
	}
	toPrint := fmt.Sprint(p...)
	toPrint = printer(color.FgBlue, "CONF", toPrint)
	fmt.Print(toPrint + " [Y/n]: ")
	var resp string
	fmt.Scanln(&resp)
	return !(strings.ToLower(strings.TrimSpace(resp)) == "n")
}

func pass(p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	var printStr string
	if checkCount != 0 {
		printStr = printer(color.FgGreen, fmt.Sprintf("%s:%d", "PASS", checkCount), toPrint)
	} else {
		printStr = printer(color.FgGreen, "PASS", toPrint)
	}
	fmt.Print(printStr)
}

func fail(p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	var printStr string
	if checkCount != 0 {
		printStr = printer(color.FgRed, fmt.Sprintf("%s:%d", "FAIL", checkCount), toPrint)
	} else {
		printStr = printer(color.FgRed, "FAIL", toPrint)
	}
	fmt.Print(printStr)
}

func warn(p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	var printStr string
	if checkCount != 0 {
		printStr = printer(color.FgYellow, fmt.Sprintf("%s:%d", "WARN", checkCount), toPrint)
	} else {
		printStr = printer(color.FgYellow, "WARN", toPrint)
	}
	fmt.Print(printStr)
}

func debug(p ...interface{}) {
	// Function inlining and dead code elimination means that debug calls
	// (and their strings) will be optimized out on phocus builds, with any
	// recent Go compiler.
	//
	// So, we can make as many debug calls as we want, and it won't make a
	// specific version of phocus easier to reverse. (Easier than it already
	// is, since source code is available.) This really only benefits those
	// with custom builds.
	if DEBUG_BUILD && debugEnabled {
		toPrint := fmt.Sprintln(p...)
		var printStr string
		if checkCount != 0 {
			printStr = printer(color.FgMagenta, fmt.Sprintf("%s:%d", "DBUG", checkCount), toPrint)
		} else {
			printStr = printer(color.FgMagenta, "DBUG", toPrint)
		}
		fmt.Print(printStr)
	}
}

func info(p ...interface{}) {
	if verboseEnabled {
		toPrint := fmt.Sprintln(p...)
		var printStr string
		if checkCount != 0 {
			printStr = printer(color.FgCyan, fmt.Sprintf("%s:%d", "INFO", checkCount), toPrint)
		} else {
			printStr = printer(color.FgCyan, "INFO", toPrint)
		}
		fmt.Print(printStr)
	}
}

func blue(head string, p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	printStr := printer(color.FgCyan, head, toPrint)
	fmt.Print(printStr)
}

func red(head string, p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	fmt.Print(printer(color.FgRed, head, toPrint))
}

func green(head string, p ...interface{}) {
	toPrint := fmt.Sprintln(p...)
	fmt.Print(printer(color.FgGreen, head, toPrint))
}

func printer(colorChosen color.Attribute, messageType, toPrint string) string {
	printer := color.New(colorChosen, color.Bold)
	printStr := "["
	printStr += printer.Sprintf(messageType)
	printStr += fmt.Sprintf("] %s", toPrint)
	return printStr
}
