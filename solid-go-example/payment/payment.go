package payment

// OCP — можно расширять новыми методами оплаты
type PaymentMethod interface {
    Pay(amount float64) string
}
