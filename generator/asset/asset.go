package asset

import "time"

type (

	AssetVersion struct {
		Attributes map[string][]string
		Content    interface{}
	}
	AssetType string
	Asset     struct {
		AssetType
		CreatedAt  time.Time
		UpdatedAt  time.Time
		ID         string
		Title      string
		Attributes []string
		Versions   map[string]AssetVersion
	}
)
