package outbox

type OutboxAggregateType string

const (
	UserOutboxAggregateType OutboxAggregateType = "user"
)

type OutboxEventType string

const (
	OutboxEventTypeUserCreated OutboxEventType = "user_created"
)
