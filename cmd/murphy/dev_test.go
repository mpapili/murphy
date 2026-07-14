package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mpapili/murphy/internal/proto"
)

func TestDefaultSpawnOccurrence(t *testing.T) {
	o := defaultSpawnOccurrence("system-test")
	if o.Type != "spawn" {
		t.Errorf("type=%q want spawn", o.Type)
	}
	if o.Location.GameID != "system-test" {
		t.Errorf("gameId=%q", o.Location.GameID)
	}
	if o.Location.OffsetX != 8 || o.Location.OffsetZ != 8 {
		t.Errorf("offsets = (%d,%d) want (8,8)", o.Location.OffsetX, o.Location.OffsetZ)
	}
}

func TestHandleDevOccurrenceMethod(t *testing.T) {
	h := handleDevOccurrence(nil, "g1")
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dev/occurrence", nil)
	h(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("GET status=%d want 405", rec.Code)
	}
}

func TestHandleDevOccurrenceNoNATS(t *testing.T) {
	h := handleDevOccurrence(nil, "g1")
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/dev/occurrence", nil)
	h(rec, req)
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("POST status=%d want 503", rec.Code)
	}
	var body map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["ok"] != false {
		t.Errorf("ok=%v want false", body["ok"])
	}
}

func TestHandleDevOccurrenceBodyTypeOverride(t *testing.T) {
	// Without NATS we still parse body before publish — confirm path used.
	// Publish fails at nil nc after decode.
	h := handleDevOccurrence(nil, "g1")
	rec := httptest.NewRecorder()
	payload := `{"type":"weather","location":{"chunkX":1,"offsetX":3,"offsetZ":4}}`
	req := httptest.NewRequest(http.MethodPost, "/dev/occurrence", strings.NewReader(payload))
	h(rec, req)
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status=%d want 503 (nats down)", rec.Code)
	}
}

func TestDevOccurrenceBodyDecode(t *testing.T) {
	// Ensure optional fields map cleanly into Occurrence defaults shape.
	raw := `{"type":"spawn","location":{"gameId":"x","chunkX":-1,"chunkZ":2,"offsetX":1,"offsetZ":2},"magnitude":0.5,"duration":10,"startTick":99}`
	var body devOccurrenceBody
	if err := json.Unmarshal([]byte(raw), &body); err != nil {
		t.Fatal(err)
	}
	if body.Type != "spawn" || body.Location == nil || body.Location.ChunkX != -1 {
		t.Fatalf("decode mismatch: %+v", body)
	}
	if body.Magnitude == nil || *body.Magnitude != 0.5 {
		t.Errorf("magnitude")
	}
	var _ proto.Occurrence // keep import useful if refactor
}
