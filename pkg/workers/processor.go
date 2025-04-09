package workers

import (
	"LedgerV2/pkg/models"
	"fmt"
	"sync"
	"sync/atomic"
)

type Processor struct {
	TransactionQueue chan *models.Transaction
	Workers          int
	Stats            Stats
	Accounts         map[int]*models.Account
	Wg               sync.WaitGroup
}

type Stats struct {
	Processed uint64
	Failed    uint64
}

func NewProcessor(workerCount int) *Processor {
	return &Processor{
		TransactionQueue: make(chan *models.Transaction, 100),
		Workers:          workerCount,
		Accounts:         make(map[int]*models.Account),
	}
}

func (p *Processor) Start() {
	for i := 0; i < p.Workers; i++ {
		p.Wg.Add(1)
		go StartWorker(p)
	}
}

func (p *Processor) Stop() {
	close(p.TransactionQueue)
	p.Wg.Wait()
}

func (p *Processor) PrintStats() {
	fmt.Printf("Processing: %d, Failed: %d\n", atomic.LoadUint64(&p.Stats.Processed), atomic.LoadUint64(&p.Stats.Failed))
}

func (p *Processor) AddTransaction(tx *models.Transaction) {
	p.TransactionQueue <- tx
}
