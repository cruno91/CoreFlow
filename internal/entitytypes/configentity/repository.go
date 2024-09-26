package configentity

import (
	"CoreFlow/internal/database"
	"encoding/json"
)

func GetConfigEntityByID(id string) (*ConfigEntity, error) {
	var entity ConfigEntity
	query := `SELECT id, name, fields, path_pattern, metadata FROM config_entities WHERE id = $1`
	row := database.DB.QueryRow(query, id)

	var fieldsJSON, metadataJSON []byte
	err := row.Scan(&entity.ID, &entity.Name, &fieldsJSON, &entity.PathPattern, &metadataJSON)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON fields
	if err := json.Unmarshal(fieldsJSON, &entity.Fields); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(metadataJSON, &entity.Metadata); err != nil {
		return nil, err
	}

	return &entity, nil
}
