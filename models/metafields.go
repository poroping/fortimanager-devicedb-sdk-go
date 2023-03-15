package models

type Metafields struct {
	Name           *string          `json:"name,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Value          *string          `json:"value,omitempty"`
	DynamicMapping []DynamicMapping `json:"dynamic_mapping,omitempty"`
}

type DynamicMapping struct {
	Scope []Scope `json:"_scope,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Scope struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Vdom *string `json:"vdom,omitempty"`
	Text *string `json:"text,omitempty"`
}
