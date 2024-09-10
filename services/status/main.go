package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"
)

var services Services

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
		return
	}

	file, err := os.ReadFile("status.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &services)
	if err != nil {
		return
	}

	checkStatus()

	ticker := time.NewTicker(15 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				checkStatus()
			}
		}
	}()

	select {}
}

func checkStatus() {
	command := strings.Fields("docker ps --format {{.Names}}")
	cmd := exec.Command(command[0], command[1:]...)

	stdout, err := cmd.Output()
	if err != nil {
		return
	}

	containers := strings.Split(string(stdout), "\n")

	for _, service := range services.Services {
		if !slices.Contains(containers, service) {
			sendDiscord(fmt.Sprintf("Service not running: %s", service), "@here", 0xFF0000)
		} else {
			sendDiscord(fmt.Sprintf("Service running: %s", service), "", 0x008000)
		}
	}
}

func sendDiscord(m string, mention string, color int) {
	message := Message{
		Username: "Status checker",
		Content:  mention,
		Embeds: []Embed{
			{
				Title:       "Status",
				Description: m + "\n" + fmt.Sprintf("*** %s ***", time.Now().Format(time.DateTime)),
				Color:       color,
			},
		},
	}

	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)
	if err != nil {
		return
	}

	_, err = http.Post(os.Getenv("DISCORD_URL"), "application/json", payload)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
}
