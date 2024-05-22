package models

type Metafields struct {
	Name           *string          `json:"name,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Value          *string          `json:"value,omitempty"`
	DynamicMapping []DynamicMapping `json:"dynamic_mapping,omitempty"`
}
