package contententity

import "time"

type ContentEntity struct {
	ID             string                 // Unique ID for the content entity
	ConfigEntityID string                 // ID of the associated config entity
	Fields         map[string]interface{} // Map of field values
	CreatedAt      time.Time              // Timestamp of creation
	UpdatedAt      time.Time              // Timestamp of last update
	Metadata       map[string]string      // Additional metadata
}
