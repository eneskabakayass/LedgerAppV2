package workers

import (
	"LedgerV2/pkg/models"
	"sync/atomic"
)

func StartWorker(p *Processor) {
	defer p.Wg.Done()

	for tx := range p.TransactionQueue {
		userID := tx.FromUserID

		acc, exists := p.Accounts[userID]
		if !exists {
			acc = &models.Account{}
			p.Accounts[userID] = acc
		}

		acc.UpdateBalance(int64(tx.Amount))

		atomic.AddUint64(&p.Stats.Processed, 1)
	}
}
