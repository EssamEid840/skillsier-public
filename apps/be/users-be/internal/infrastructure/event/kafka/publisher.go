package kafka

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaPublisher struct {
	client *kgo.Client
}

type KafkaProducerConfig struct {
	Host    string // e.g. "localhost:9092"
	GroupID string // not used by producer, kept for parity
}

func NewKafkaPublisher(cfg KafkaProducerConfig) *KafkaPublisher {
	if cfg.Host == "" {
		log.Fatal("kafka host is empty")
	}
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Host),
		// add TLS/SASL config here if needed
	)
	if err != nil {
		log.Fatalf("kafka: create client: %v", err)
	}
	return &KafkaPublisher{client: cl}
}

// EnsureTopics ensures the given topics exist (idempotent).
func (p *KafkaPublisher) EnsureTopics(ctx context.Context, topics []string) error {
	adm := kadm.NewClient(p.client)
	defer adm.Close()

	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	td, err := adm.ListTopics(cctx, topics...)
	if err != nil {
		log.Printf("kafka: ListTopics failed (will attempt create anyway): %v", err)
	}

	toCreate := make([]string, 0, len(topics))
	for _, t := range topics {
		if td != nil {
			if _, ok := td[t]; ok {
				continue // already exists
			}
		}
		toCreate = append(toCreate, t)
	}
	if len(toCreate) == 0 {
		return nil
	}

	cctx2, cancel2 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel2()
	res, err := adm.CreateTopics(cctx2, 1 /*partitions*/, 1 /*replication*/, nil, toCreate...)
	if err != nil {
		return err
	}
	for _, r := range res {
		if r.Err != nil && !errors.Is(r.Err, kerr.TopicAlreadyExists) {
			return r.Err
		}
	}
	return nil
}

// Publish matches event.EventPublisher: key is string, payload is []byte.
// If the topic is missing, create it and retry once.
// Return errors (do not log.Fatal) so the dispatcher can retry/backoff.
func (p *KafkaPublisher) Publish(ctx context.Context, topic string, key string, payload []byte) error {
	var keyBytes []byte
	if key != "" {
		keyBytes = []byte(key)
	}
	rec := &kgo.Record{
		Topic: topic,
		Key:   keyBytes,
		Value: payload,
	}

	// 1st attempt
	c1, cancel1 := context.WithTimeout(ctx, 5*time.Second)
	err := p.client.ProduceSync(c1, rec).FirstErr()
	cancel1()
	if err == nil {
		return nil
	}

	// If topic doesn't exist yet, create & retry once
	if errors.Is(err, kerr.UnknownTopicOrPartition) {
		if e := p.ensureTopic(ctx, topic, 1, 1); e != nil {
			log.Printf("kafka: ensureTopic(%s) failed: %v", topic, e)
			return err // original publish error; caller may retry later
		}
		c2, cancel2 := context.WithTimeout(ctx, 5*time.Second)
		defer cancel2()
		if e := p.client.ProduceSync(c2, rec).FirstErr(); e != nil {
			log.Printf("kafka: produce retry failed: %v", e)
			return e
		}
		return nil
	}

	// Other errors (network/auth/etc.)
	log.Printf("kafka: produce failed: %v", err)
	return err
}

// internal helper to create a single topic
func (p *KafkaPublisher) ensureTopic(ctx context.Context, topic string, partitions int32, replication int16) error {
	adm := kadm.NewClient(p.client)
	defer adm.Close()

	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := adm.CreateTopics(cctx, partitions, replication, nil, topic)
	if err != nil {
		return err
	}
	if tr, ok := res[topic]; ok && tr.Err != nil && !errors.Is(tr.Err, kerr.TopicAlreadyExists) {
		return tr.Err
	}
	return nil
}

func (p *KafkaPublisher) Close() {
	p.client.Close()
}
