package http

import (
	"LedgerV2/internal/http/handlers"
	"LedgerV2/internal/http/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(txHandler *handlers.TransactionHandler, userHandler *handlers.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)

	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.Route("/api/v1", func(api chi.Router) {

		api.Group(func(pub chi.Router) {
			pub.Post("/auth/register", handlers.RegisterHandler)
			pub.Post("/auth/login", handlers.LoginHandler)
			pub.Post("/auth/refresh", handlers.RefreshHandler)
		})

		api.Group(func(private chi.Router) {
			private.Route("/users", func(u chi.Router) {
				u.Get("/", userHandler.GetAllUsers)
				u.Get("/{id}", userHandler.GetUserByID)
				u.Put("/{id}", userHandler.UpdateUser)
				u.Delete("/{id}", userHandler.DeleteUser)
			})

			private.Route("/transactions", func(t chi.Router) {
				t.Post("/credit", txHandler.CreditHandler)
				t.Post("/debit", txHandler.DebitHandler)
				t.Post("/transfer", txHandler.TransferHandler)
				t.Get("/history", txHandler.TransactionHistoryHandler)
				t.Get("/{id}", txHandler.GetTransactionByIDHandler)
			})

			private.Route("/balances", func(b chi.Router) {
				b.Get("/current", handlers.CurrentBalanceHandler)
				b.Get("/historical", handlers.HistoricalBalanceHandler)
				b.Get("/at-time", handlers.BalanceAtTimeHandler)
			})
		})
	})

	return r
}
