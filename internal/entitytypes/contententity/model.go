package contententity

import (
	"github.com/google/uuid"
	"time"
)

type ContentEntity struct {
	ID             uuid.UUID              `json:"id"`
	ConfigEntityID uuid.UUID              `json:"configEntityId"`
	Bundle         string                 `json:"bundle"`
	Name           string                 `json:"name"`
	Fields         map[string]interface{} `json:"fields"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	Metadata       map[string]string      `json:"metadata"`
}
