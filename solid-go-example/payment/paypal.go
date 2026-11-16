package payment

// Еще одна реализация (OCP)
type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
    return "Paid with PayPal"
}
