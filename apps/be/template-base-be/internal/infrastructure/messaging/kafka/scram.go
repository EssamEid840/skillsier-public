package kafka

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/IBM/sarama"
)

// HashGeneratorFcn is a function that returns a hash generator
type HashGeneratorFcn func() hash.Hash

var (
	// SHA256 is the hash generator for SCRAM-SHA-256
	SHA256 HashGeneratorFcn = sha256.New
	
	// SHA512 is the hash generator for SCRAM-SHA-512
	SHA512 HashGeneratorFcn = sha512.New
)

// SCRAMClient implements the sarama.SCRAMClient interface
type SCRAMClient struct {
	*sarama.SCRAMClient
	HashGeneratorFcn HashGeneratorFcn
}

// Begin starts the SCRAM authentication process
func (s *SCRAMClient) Begin(userName, password, authzID string) (err error) {
	s.SCRAMClient, err = s.HashGeneratorFcn().(*sarama.SCRAMClient).Begin(userName, password, authzID)
	return err
}

// Step processes a challenge from the server
func (s *SCRAMClient) Step(challenge string) (response string, err error) {
	return s.SCRAMClient.Step(challenge)
}

// Done indicates whether the authentication is complete
func (s *SCRAMClient) Done() bool {
	return s.SCRAMClient.Done()
}