package notification

import "fmt"

// Реализация SMS уведомления
type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
    fmt.Println("SMS:", message)
}
