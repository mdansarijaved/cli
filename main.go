package main

import (
	userinput "cliyounicorn/userInput"
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

	if _, err := exec.Command("pm2", "-v").Output(); err != nil {
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
	} else {
		fmt.Println("Pm2 already installed")
	}

	if _, err := exec.Command("nginx", "-v").Output(); err != nil {
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
	} else {
		fmt.Println("Nginx already installed")
	}

	if _, err := exec.Command("git", "--version").Output(); err != nil {
		gitInstall := exec.Command("bash", "-c", "sudo apt-get install -y git")

		if err := gitInstall.Start(); err != nil {
			log.Fatal(err)
		}

		gitSpinner := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Build our new spinner
		gitSpinner.Start()                                                    // Start the spinner

		if err := gitInstall.Wait(); err != nil {
			gitSpinner.Stop()
			log.Fatal(err)
		} else {

			gitSpinner.Stop()
			fmt.Println("Git installed")
		}
	} else {
		fmt.Println("Git already installed")
	}

	YounicornClone := exec.Command("bash", "-c", "cd / && git clone https://github.com/kratos-respawned/Younicorn.git")

	if err := YounicornClone.Start(); err != nil {
		log.Fatal(err)
	}

	YounicornSpinner := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Build our new spinner
	YounicornSpinner.Start()                                                    // Start the spinner

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Working Directory:", cwd)

	if err := YounicornClone.Wait(); err != nil {
		YounicornSpinner.Stop()
		log.Fatal(err)
	} else {
		YounicornSpinner.Stop()
		fmt.Println("Younicorn cloned")

	}
	userinput.GetEnv()
	// Wait for the command to finish

	npmRun := exec.Command("bash", "-c", "cd Younicorn && npm i && npm run build")

	if err := npmRun.Start(); err != nil {
		log.Fatal(err)
	}

}
