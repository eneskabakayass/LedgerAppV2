package workers

import (
	"LedgerV2/pkg/models"
	"sync/atomic"
)

func StartWorker(p *Processor) {
	defer p.Wg.Done()
	for tx := range p.TransactionQueue {
		acc, exists := p.Accounts[tx.UserID]
		if !exists {
			acc = &models.Account{}
			p.Accounts[tx.UserID] = acc
		}
		acc.UpdateBalance(int64(tx.Amount))
		atomic.AddUint64(&p.Stats.Processed, 1)
	}
}
