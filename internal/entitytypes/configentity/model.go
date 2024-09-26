package configentity

type FieldDefinition struct {
	Name     string            // Field name (e.g., "title")
	Type     string            // Data type (e.g., "string", "text", "int")
	Required bool              // Whether the field is required
	Settings map[string]string // Additional settings (e.g., max length)
}

type ConfigEntity struct {
	ID          string            // Unique ID for the config entity
	Name        string            // Name of the content type
	Fields      []FieldDefinition // List of field definitions
	PathPattern string            // Path alias pattern (e.g., "/articles/{title}")
	Metadata    map[string]string // Additional metadata
}
