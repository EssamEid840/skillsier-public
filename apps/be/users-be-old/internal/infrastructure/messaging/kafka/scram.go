package kafka

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/xdg-go/scram"
)

// HashGeneratorFcn is a function that returns a hash.Hash
type HashGeneratorFcn func() hash.Hash

var (
	// SHA256 is the SHA256 hash generator
	SHA256 HashGeneratorFcn = sha256.New
	// SHA512 is the SHA512 hash generator
	SHA512 HashGeneratorFcn = sha512.New
)

// XDGSCRAMClient implements the sarama.SCRAMClient interface
// This is required for SCRAM-SHA-512 authentication with Kafka
type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	HashGeneratorFcn HashGeneratorFcn
}

// Begin starts the SCRAM authentication conversation
func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	client, err := x.HashGeneratorFcn.NewClient(userName, password, authzID)  // ✅ Use our wrapper
	if err != nil {
		return err
	}
	x.Client = client
	x.ClientConversation = x.Client.NewConversation()
	return nil
}
// Step processes a single step in the SCRAM authentication
func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

// Done indicates if the SCRAM authentication is complete
func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}

// scram creates a new SCRAM client
func (f HashGeneratorFcn) NewClient(userName, password, authzID string) (*scram.Client, error) {
	// Create hash generator for SCRAM
	hashGen := scram.HashGeneratorFcn(f)
	return hashGen.NewClient(userName, password, authzID)  // ✅ Correct API
}


