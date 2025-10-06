package daprpub

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	BaseURL    string // e.g. http://localhost:3501
	PubsubName string // e.g. kafka-pubsub
}

type DaprPublisher struct {
	cfg    Config
	client *http.Client
}

func New(cfg Config) *DaprPublisher {
	return &DaprPublisher{
		cfg: cfg,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (p *DaprPublisher) Publish(ctx context.Context, topic string, key string, payload []byte) error {
	// Partition key is optional; Kafka honors it via Dapr metadata.
	u := fmt.Sprintf("%s/v1.0/publish/%s/%s?metadata.partitionKey=%s",
		p.cfg.BaseURL,
		url.PathEscape(p.cfg.PubsubName),
		url.PathEscape(topic),
		url.QueryEscape(key),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("dapr publish failed: %s (%s)", resp.Status, string(b))
	}
	return nil
}

func (p *DaprPublisher) Close() {} // no-op
