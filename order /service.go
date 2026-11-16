package order

// SRP — отдельный сервис для логики
type OrderService struct{}

func (s *OrderService) CreateOrder(amount float64) Order {
    return Order{
        ID:     1,
        Amount: amount,
    }
}
