package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetEnv() {

	reader := bufio.NewReader(os.Stdin)

	// List of environment variables to get from the user
	envVars := []string{
		"DATABASE_URL",
		"GITHUB_CLIENT_ID",
		"GITHUB_CLIENT_SECRET",
		"NEXTAUTH_SECRET",
		"NEXTAUTH_URL",
		"ADMIN_MAIL",
		"ADMIN_PASS",
	}

	// Open the .env file
	file, err := os.OpenFile("/Younicorn/.env", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get each environment variable from the user
	for _, envVar := range envVars {
		fmt.Printf("Enter %s: ", envVar)
		value, _ := reader.ReadString('\n')

		// Write the environment variable to the file
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", envVar, strings.TrimSpace(value)))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Environment variables saved to .env")
}
