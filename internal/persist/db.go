package persist

import "log"

func ConnectPostgres(url string) {
	if url == "" {
		log.Printf("persist: DATABASE_URL empty; skipping Postgres")
		return
	}
	log.Printf("persist: Postgres stub ready (url configured)")
}
