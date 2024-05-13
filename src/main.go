package main

import (
    "log"
    "os"
    "fmt"
    "net/smtp"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Error loading .env file")
    }
    
    var subject string
    var body string

    // !!!   Uncomment these to write these in code   !!!
    // subject = "This is a test email"
    // body = "Body test of the email"

    args := os.Args
    if len(os.Args) < 3 && subject == "" && body == "" {
        fmt.Println("Usage: go run main.go <subject> <body>\nMake sure to use double quotes in the subject and body if you are going to use spaces!")
        os.Exit(1)
    } else if subject == "" && body == "" {
        subject = args[1]
        body = args[2]
    }

    fromEmail := os.Getenv("EMAIL_USERNAME")
    toEmail := os.Getenv("TO_EMAIL")

    from := fromEmail
    to := []string{toEmail}
    addr := "smtp.gmail.com:587"
    host := "smtp.gmail.com"

    emailErr := sendEmail(from, to, addr, host, subject, body)
    if emailErr!= nil {
        log.Println(emailErr)
    } else {
        log.Println("Email sent successfully!")
    }
}

func sendEmail(from string, to []string, addr, host, subject, body string) error {
    emailUsername := os.Getenv("EMAIL_USERNAME")
    emailPassword := os.Getenv("EMAIL_PASSWORD")

    var msg string
    msg = "From: " + from + "\r\n" +
    "To: " + to[0] + "\r\n" +
    "Subject: " + subject + "\r\n\r\n" +
    body + "\r\n"

    msgBytes := []byte(msg)

    auth := smtp.PlainAuth("", emailUsername, emailPassword, host)

    err := smtp.SendMail(addr, auth, from, to, msgBytes)
    if err != nil {
        return err
    }

    return nil
}
