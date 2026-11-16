package notification

// DIP — абстракция уведомлений
type Notifier interface {
    Send(message string)
}

// Использует интерфейс, а не конкретный тип
type OrderProcessor struct {
    Notifier Notifier
}

func (op OrderProcessor) Process() {
    op.Notifier.Send("Order created")
}
