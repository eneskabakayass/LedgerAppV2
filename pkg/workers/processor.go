package workers

import (
	"LedgerV2/pkg/models"
	"errors"
	"github.com/rs/zerolog/log"
	"sync"
)

type Stats struct {
	Processed uint64
}

type Processor struct {
	Accounts         map[string]*models.Account
	TransactionQueue chan *models.Transaction
	Wg               sync.WaitGroup
	Stats            Stats
}

func NewProcessor(workerCount int) *Processor {
	p := &Processor{
		Accounts:         make(map[string]*models.Account),
		TransactionQueue: make(chan *models.Transaction, 100),
	}

	for i := 0; i < workerCount; i++ {
		p.Wg.Add(1)
		go StartWorker(p)
	}

	return p
}

func (p *Processor) Stop() {
	close(p.TransactionQueue)
	p.Wg.Wait()
}

func (p *Processor) Deposit(userID string, amount float64) error {
	acc, exists := p.Accounts[userID]
	if !exists {
		acc = &models.Account{UserID: userID}
		p.Accounts[userID] = acc
	}
	acc.UpdateBalance(int64(amount))
	return nil
}

func (p *Processor) Withdraw(userID string, amount float64) error {
	acc, exists := p.Accounts[userID]
	if !exists {
		return errors.New("account not found")
	}
	if acc.Balance < int64(amount) {
		return errors.New("insufficient funds")
	}
	acc.UpdateBalance(-int64(amount))
	return nil
}

func (p *Processor) GetBalance(userID string) float64 {
	acc, exists := p.Accounts[userID]
	if !exists {
		return 0
	}
	return float64(acc.Balance)
}

func (p *Processor) Start() {
	for i := 0; i < len(p.Accounts); i++ {
		p.Wg.Add(1)
		go StartWorker(p)
	}
}

func (p *Processor) PrintStats() {
	log.Logger.Println("Processed transactions:", p.Stats.Processed)
}
