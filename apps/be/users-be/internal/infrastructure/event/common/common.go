package common

import "users.go/m/internal/entities/outbox"

const (
	UserTopic string = "user"
)

var AggregateTypeToTopic = map[string]string{
	string(outbox.UserOutboxAggregateType): UserTopic,
}
