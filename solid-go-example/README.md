# SOLID Principles in Go

Этот проект демонстрирует использование принципов SOLID на языке Go.

## Реализованные принципы:

### 1. Single Responsibility Principle (SRP)
- order/order.go — только структура данных
- order/service.go — только бизнес-логика

### 2. Open/Closed Principle (OCP)
- Интерфейс PaymentMethod
- Новые способы оплаты добавляются без изменения существующего кода

### 3. Liskov Substitution Principle (LSP)
- Любая реализация PaymentMethod может использоваться в ProcessPayment

### 4. Interface Segregation Principle (ISP)
- Используются небольшие и узконаправленные интерфейсы

### 5. Dependency Inversion Principle (DIP)
- OrderProcessor зависит от интерфейса Notifier

## Запуск

go run main.go
