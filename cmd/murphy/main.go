// Command murphy is the in-world occurrence generator microservice (skeleton).
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/mpapili/murphy/internal/bus"
	"github.com/mpapili/murphy/internal/ingest"
	"github.com/mpapili/murphy/internal/persist"
	"github.com/mpapili/murphy/internal/sched"
)

func main() {
	port := envOr("PORT", "8091")
	gameID := envOr("GAME_ID", "local-dev")
	natsURL := envOr("NATS_URL", "nats://localhost:4222")
	interval, _ := strconv.Atoi(envOr("OCCURRENCE_INTERVAL_SEC", "15"))

	persist.ConnectPostgres(os.Getenv("DATABASE_URL"))

	nc := bus.Connect(natsURL)
	if nc != nil {
		defer nc.Drain()
	}
	ingest.SubscribeState(nc, gameID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go sched.Run(ctx, nc, gameID, interval, 42)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":  "ok",
			"service": "murphy",
			"gameId":  gameID,
		})
	})

	srv := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		log.Printf("murphy listening on http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
	shctx, shcancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shcancel()
	_ = srv.Shutdown(shctx)
	log.Printf("murphy stopped")
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
