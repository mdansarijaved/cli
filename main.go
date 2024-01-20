package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	if os.Getenv("USER") != "root" {
		log.Fatal("This program must be run as root (sudo)")
	}

	nodeInstall := exec.Command("bash", "-c", "curl -sL https://deb.nodesource.com/setup_20.x | sudo -E bash - && sudo apt-get install -y nodejs")

	// Start the command
	if err := nodeInstall.Start(); err != nil {
		log.Fatal(err)
	}

	// Start a goroutine to print a spinner
	go func() {
		spinner := []string{"-", "\\", "|", "/"}
		i := 0
		for {
			fmt.Printf("Installing node.... \r%s", spinner[i%len(spinner)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}()
	if err := nodeInstall.Wait(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Node installed")
	}

	cmd2 := exec.Command("bash", "-c", "sudo npm install -g pm2")

	// Start the command
	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		spinner := []string{"-", "\\", "|", "/"}
		i := 0
		for {
			fmt.Printf("Installing Pm2.... \r%s", spinner[i%len(spinner)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}()

	if err := cmd2.Wait(); err != nil {
		log.Fatal(err)
	}

	// Wait for the command to finish

}
