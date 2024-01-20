package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func main() {

	if os.Getenv("USER") != "root" {
		log.Fatal("This program must be run as root (sudo)")
	}
	fmt.Println("Installing Nodejs and Pm2")
	if _, err := exec.Command("node", "-v").Output(); err != nil {

		nodeInstall := exec.Command("bash", "-c", "curl -sL https://deb.nodesource.com/setup_20.x | sudo -E bash - && sudo apt-get install -y nodejs")

		// Start the command
		if err := nodeInstall.Start(); err != nil {
			log.Fatal(err)
		}

		// Start a goroutine to print a spinner
		nodespinner := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Build our new spinner
		nodespinner.Start()                                                    // Start the spinner
		if err := nodeInstall.Wait(); err != nil {
			nodespinner.Stop()
			log.Fatal(err)
		} else {
			nodespinner.Stop()
			fmt.Println("Node installed")
		}
	} else {
		fmt.Println("Node already installed")
	}

	InstallPm2 := exec.Command("bash", "-c", "sudo npm install -g pm2")

	if err := InstallPm2.Start(); err != nil {
		log.Fatal(err)
	}

	pm2spinner := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Build our new spinner
	pm2spinner.Start()                                                    // Start the spinner

	// Start the command

	if err := InstallPm2.Wait(); err != nil {
		pm2spinner.Stop()
		log.Fatal(err)
	} else {
		pm2spinner.Stop()
		fmt.Println("Pm2 installed")
	}

	nginxInstall := exec.Command("bash", "-c", "sudo apt-get install -y nginx")

	if err := nginxInstall.Start(); err != nil {
		log.Fatal(err)
	}

	nginxSpinner := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Build our new spinner
	nginxSpinner.Start()                                                    // Start the spinner

	if err := nginxInstall.Wait(); err != nil {
		nginxSpinner.Stop()
		log.Fatal(err)
	} else {
		nginxSpinner.Stop()
		fmt.Println("Nginx installed")
	}

	// Wait for the command to finish

}
