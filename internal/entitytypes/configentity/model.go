package configentity

import "github.com/google/uuid"

type FieldDefinition struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Required bool              `json:"required"`
	Settings map[string]string `json:"settings"`
}

type ConfigEntity struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Fields      []FieldDefinition `json:"fields"`
	PathPattern string            `json:"pathPattern"`
	Metadata    map[string]string `json:"metadata"`
}
