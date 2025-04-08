package services

type BalanceService interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}
