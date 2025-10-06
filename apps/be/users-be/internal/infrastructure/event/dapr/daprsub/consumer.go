package daprsub

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"users.go/m/internal/infrastructure/event"
)

// Config matches your existing shape.
type Config struct {
	AppPort    string // e.g. ":8082"
	PubsubName string // e.g. "kafka-pubsub"
	Topic      string // e.g. "user"
	Route      string // e.g. "/events/users"
}

// KeySelector lets you derive a partition key from the message.
// raw is the final payload delivered to your app (CloudEvent .data if present, otherwise raw body).
// ce is non-nil when the incoming message was a CloudEvent (minimal fields populated).
type KeySelector func(raw []byte, ce *CloudEvent) []byte

// Minimal CloudEvent envelope (sufficient for Dapr pub/sub HTTP).
type CloudEvent struct {
	ID          string          `json:"id,omitempty"`
	Type        string          `json:"type,omitempty"`
	SpecVersion string          `json:"specversion,omitempty"`
	Source      string          `json:"source,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`
}

// DaprConsumer is a generic HTTP subscriber for Dapr pubsub.
type DaprConsumer struct {
	cfg           Config
	keySelector   KeySelector
	rawCloudEvent bool // if true, deliver the full CloudEvent JSON instead of ce.data
}

func New(cfg Config) *DaprConsumer { return &DaprConsumer{cfg: cfg} }

// Optional: set a custom key selector (e.g., extract "aggregate_id" from payload).
func (c *DaprConsumer) WithKeySelector(sel KeySelector) *DaprConsumer {
	c.keySelector = sel
	return c
}

// Optional: deliver raw CloudEvent JSON to the channel (instead of ce.data).
func (c *DaprConsumer) WithRawCloudEvent() *DaprConsumer {
	c.rawCloudEvent = true
	return c
}

func (c *DaprConsumer) Start(ctx context.Context) (<-chan event.ConsumerMessage, error) {
	ch := make(chan event.ConsumerMessage, 128)

	r := gin.New()
	r.Use(gin.Recovery())

	// Programmatic subscription (works locally and in k8s without a Subscription CRD)
	r.GET("/dapr/subscribe", func(g *gin.Context) {
		type sub struct {
			PubsubName string `json:"pubsubname"`
			Topic      string `json:"topic"`
			Route      string `json:"route"`
		}
		g.JSON(http.StatusOK, []sub{{
			PubsubName: c.cfg.PubsubName,
			Topic:      c.cfg.Topic,
			Route:      c.cfg.Route,
		}})
	})

	// Health check (handy for readiness probes)
	r.GET("/healthz", func(g *gin.Context) { g.String(http.StatusOK, "ok") })

	// Dapr delivers messages here; may be CloudEvent or raw JSON.
	r.POST(c.cfg.Route, func(g *gin.Context) {
		body, _ := g.GetRawData()

		// Try CloudEvent unwrap.
		var ce CloudEvent
		isCE := false
		if json.Unmarshal(body, &ce) == nil && ce.SpecVersion != "" && len(ce.Data) > 0 {
			isCE = true
		}

		// Decide what to deliver to the app.
		var delivered []byte
		if c.rawCloudEvent && isCE {
			delivered = body // full CloudEvent
		} else if isCE {
			delivered = ce.Data // app payload only
		} else {
			delivered = body // raw non-CloudEvent
		}

		// Optional partition key selection.
		var key []byte
		if c.keySelector != nil {
			if isCE {
				key = c.keySelector(delivered, &ce)
			} else {
				key = c.keySelector(delivered, nil)
			}
			// Trim empty keys to nil for consistency
			if len(strings.TrimSpace(string(key))) == 0 {
				key = nil
			}
		}

		// Push to the channel for the caller to handle generically.
		ch <- event.ConsumerMessage{Key: key, Value: delivered}
		g.Status(http.StatusOK)
	})

	srv := &http.Server{Addr: c.cfg.AppPort, Handler: r}

	// Graceful shutdown when ctx is canceled
	go func() {
		<-ctx.Done()
		ctx2, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx2)
		close(ch)
	}()

	// Run server
	go func() { _ = srv.ListenAndServe() }()

	return ch, nil
}

func (c *DaprConsumer) Stop() { /* handled via ctx cancel in Start */ }
