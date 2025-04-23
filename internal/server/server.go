package server

import (
	"LedgerV2/internal/http"
	"LedgerV2/internal/http/handlers"
	"LedgerV2/pkg/repositories"
	"LedgerV2/pkg/services"
	"context"
	"github.com/rs/zerolog/log"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartWithService(port string, txService *services.TransactionService) {
	txHandler := &handlers.TransactionHandler{Service: txService}
	userService := services.NewUserService(repositories.UserRepo)
	userHandler := &handlers.UserHandler{Service: userService}
	router := http.NewRouter(txHandler, userHandler)

	srv := &stdhttp.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info().Msgf("Server started on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != stdhttp.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Shutdown error")
	} else {
		log.Info().Msg("Shutdown complete")
	}
}

func securityHeaders(next stdhttp.Handler) stdhttp.Handler {
	return stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}
