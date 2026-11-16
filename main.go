package main

import (
    "fmt"
    "solid-go-example/notification"
    "solid-go-example/payment"
)

// LSP — любая реализация PaymentMethod подходит
func ProcessPayment(p payment.PaymentMethod) {
    fmt.Println(p.Pay(100))
}

func main() {
    // OCP + LSP
    stripe := payment.Stripe{}
    paypal := payment.PayPal{}

    ProcessPayment(stripe)
    ProcessPayment(paypal)

    // DIP — зависимость от интерфейса, а не от реализации
    email := notification.EmailNotifier{}
    sms := notification.SMSNotifier{}

    op1 := notification.OrderProcessor{Notifier: email}
    op2 := notification.OrderProcessor{Notifier: sms}

    op1.Process()
    op2.Process()
}
