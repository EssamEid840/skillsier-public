package postgres

import (
	"fmt"
	"sync"
)

// entityRegistry holds all registered entity models for auto-migration
var (
	entityRegistry = &EntityRegistry{
		models: make([]interface{}, 0),
		mu:     sync.RWMutex{},
	}
)

// EntityRegistry manages entity models for auto-migration
type EntityRegistry struct {
	models []interface{}
	mu     sync.RWMutex
}

// Register registers entity models for auto-migration
// This should be called in init() functions of entity packages
func Register(models ...interface{}) {
	entityRegistry.Register(models...)
}

// Register adds models to the registry
func (r *EntityRegistry) Register(models ...interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, model := range models {
		if model == nil {
			continue
		}
		r.models = append(r.models, model)
	}
}

// GetModels returns all registered models
func (r *EntityRegistry) GetModels() []interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to prevent external modification
	models := make([]interface{}, len(r.models))
	copy(models, r.models)
	return models
}

// Clear clears all registered models (used for testing)
func (r *EntityRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.models = make([]interface{}, 0)
}

// Count returns the number of registered models
func (r *EntityRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.models)
}

// GetRegisteredModels returns all registered entity models
func GetRegisteredModels() []interface{} {
	return entityRegistry.GetModels()
}

// PrintRegisteredModels prints all registered models (for debugging)
func PrintRegisteredModels() {
	models := entityRegistry.GetModels()
	fmt.Printf("Registered %d entity models for auto-migration:\n", len(models))
	for i, model := range models {
		fmt.Printf("  %d. %T\n", i+1, model)
	}
}