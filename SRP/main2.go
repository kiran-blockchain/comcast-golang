package main

import (
    "fmt"
    "time"
)

// Authentication responsibility
type Authenticator struct{}

func (a *Authenticator) Authenticate(username, password string) bool {
    // Authentication logic
   // fmt.Println("Authenticating user...")
    return username == "admin" && password == "password"
}

// Logging responsibility
type Logger struct{}

func (l *Logger) Log(message string) {
    fmt.Println(time.Now().Format(time.RFC3339), message)
}

func main() {
    authenticator := &Authenticator{}
    logger := &Logger{}

    username := "admin"
    password := "password"
    
    if authenticator.Authenticate(username, password) {
        logger.Log("Authentication successful for user: " + username)
    } else {
        logger.Log("Authentication failed for user: " + username)
    }
}
