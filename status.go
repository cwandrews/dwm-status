package main

import (
//	"os"
	"os/exec"
	"time"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

func getDate() string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, _ := now.Clock()
	weekDay := now.Weekday()
	timeStr := fmt.Sprintf("%s %s %d, %d | %d:%d",
		weekDay.String(), month.String(), day, year, hour, min)
	return timeStr
}

func getUnread() string {
	numUnread, _ := ioutil.ReadFile("/home/charlie/.num_unread")
	if string(numUnread) == "0" || string(numUnread) == "" {
		return ""
	} else {
		return fmt.Sprintf("Unread: %s", string(numUnread[:1]))
	}
}

func BatteryStats() string {
	ret, err := exec.Command("acpi").Output()
	if err != nil {
		return "acpi error"
	}
	stuffICareAbout := strings.SplitN(string(ret), ":", 2)[1]
	acpiSlice := strings.Split(stuffICareAbout, ", ")
	status := acpiSlice[0]
	perc := acpiSlice[1]
	remaining := strings.Split(acpiSlice[2], " ")[0]
	return fmt.Sprintf("%s, %s, %s", status, perc, remaining)

}

func main() {
	for {
		status := fmt.Sprintf("%s | %s | %s", getUnread(), getDate(), BatteryStats())
		err := exec.Command("xsetroot", "-name", status).Run()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}
}
