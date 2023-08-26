package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const timerFile = "/tmp/pomodoro.txt"

func main() {

	currentTime := time.Now().Format("15:04")

	// default alert
	if len(os.Args) < 2 {
		defaultAlert(currentTime)
		return
	}

	// -n: notification alert
	if os.Args[1] == "-n" {
		notificationAlert(currentTime)
		return
	}

	if os.Args[1] == "pmdr" {
		pomodoroAlert(currentTime)
		return
	}
}

func defaultAlert(currentTime string) {
	cmd := exec.Command("osascript", "-e", `display dialog "`+currentTime+`" with title "default alert"`)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to display dialog: %v\n", err)
	}
}

func notificationAlert(currentTime string) {
	cmd := exec.Command("osascript", "-e", `display notification "`+currentTime+`" with title "notification alert"`)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to display notification: %v\n", err)
	}
}

func pomodoroAlert(currentTime string) {
		if _, err := os.Stat(timerFile); os.IsNotExist(err) {
			fmt.Println("Have a nice pomodoro!")
			cmd := exec.Command("/bin/bash", "./pomodoro.sh")
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to execute pomodoro.sh: %v\n", err)
			}
		} else {
			dat, err := ioutil.ReadFile(timerFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read timerFile: %v\n", err)
			}
			endTimeStr := strings.TrimSpace(string(dat))
			endTime, err := strconv.Atoi(string(endTimeStr))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse timerFile: %v\n", err)
			}
			now := int(time.Now().Unix())
			timeLeft := endTime - now
			fmt.Println("You have", timeLeft, "seconds left")
		}
}
