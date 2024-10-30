package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func init() {
	readEnvFile()
}

func GetInput(year uint16, day uint8) (string, error) {
	session := os.Getenv("SESSION")
	cookiejar, err := cookiejar.New(nil)

	parsedUrl, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))

	if err != nil {
		return "", fmt.Errorf("Error parsing URL: %v", err)
	}

	cookiejar.SetCookies(parsedUrl, []*http.Cookie{{
		Name:  "session",
		Value: session,
	}})

	if err != nil {
		return "", fmt.Errorf("Error creating cookie jar: %v", err)
	}

	client := &http.Client{
		Jar: cookiejar,
	}

	resp, err := client.Get(parsedUrl.String())

	if err != nil {
		return "", fmt.Errorf("Error getting URL: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	return string(body), nil
}

func readEnvFile() {
	file, err := os.ReadFile(".env")

	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}

		env := strings.Split(line, "=")
		os.Setenv(env[0], env[1])
	}
}
