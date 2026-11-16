package payment

// Конкретная реализация оплаты
type Stripe struct{}

func (s Stripe) Pay(amount float64) string {
    return "Paid with Stripe"
}
