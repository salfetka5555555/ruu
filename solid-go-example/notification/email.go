package notification

import "fmt"

// Реализация уведомления по email
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
    fmt.Println("Email:", message)
}
