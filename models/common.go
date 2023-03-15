package models

const CmdbBasePath = "/pm/config/device/"

// "/pm/config/device/${device}/vdom/${vdom}/"

type FmgCmdbRequest struct {
	ID      int                     `json:"id"`
	Verbose int                     `json:"verbose"`
	Method  string                  `json:"method"`
	Params  []FmgCmdbRequestPayload `json:"params"`
	Session string                  `json:"session"`
}

type CmdbRequest struct {
	HTTPMethod string
	Payload    []byte
	Params     CmdbRequestParams
	Path       string
	Mkey       *string
	NoVdom     bool
}

type FmgCmdbRequestPayload struct {
	Data interface {
	} `json:"data"`
	URL string `json:"url"`
}

type CmdbRequestParams struct {
	Action               string    `json:"action,omitempty"`
	Adom                 string    `json:"adom,omitempty"`
	AllowAppend          *bool     `json:"allow_append,omitempty"`
	After                string    `json:"after,omitempty"`
	Before               string    `json:"before,omitempty"`
	Datasource           *bool     `json:"datasource,omitempty"`
	ExcludeDefaultValues *bool     `json:"exclude-default-values,omitempty"`
	Filter               *[]string `json:"filter,omitempty"`
	Format               *[]string `json:"format,omitempty"`
	Meta                 *bool     `json:"meta,omitempty"`
	Mkey                 string    `json:"mykey,omitempty"`
	PlainTextPassword    *bool     `json:"plain-text-password,omitempty"`
	Scope                string    `json:"scope,omitempty"`
	Sort                 *[]string `json:"sort,omitempty"`
	Type                 string    `json:"type,omitempty"`
	Vdom                 string    `json:"vdom,omitempty"`
	WithMeta             *bool     `json:"with_meta,omitempty"`
}

type CmdbResponse struct {
	HTTPMethod      string      `json:"http_method,omitempty"`
	Size            *int64      `json:"size,omitempty"`
	MatchedCount    *int64      `json:"matched_count,omitempty"`
	NextIdx         *int64      `json:"next_idx,omitempty"`
	Revision        *string     `json:"revision,omitempty"`
	RevisionChanged *bool       `json:"revision_changed,omitempty"`
	OldRevision     *string     `json:"old_revision,omitempty"`
	Results         interface{} `json:"results,omitempty"`
	Vdom            *string     `json:"vdom,omitempty"`
	Path            *string     `json:"path,omitempty"`
	ChildPath       *string     `json:"child_path,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Mkey            interface{} `json:"mkey,omitempty"`
	ChildMkey       interface{} `json:"child_mkey,omitempty"`
	Status          string      `json:"status,omitempty"`
	HTTPStatus      int64       `json:"http_status,omitempty"`
	Serial          *string     `json:"serial,omitempty"`
	Version         *string     `json:"version,omitempty"`
	Build           *int64      `json:"build,omitempty"`
	FullPath        *string     `json:"FullPath,omitempty"`
}

type FmgCmdbResponse struct {
	ID      int64    `json:"id,omitempty"`
	Result  []Result `json:"result,omitempty"`
	Session *string  `json:"session,omitempty"`
}

type Result struct {
	Data   interface{} `json:"data,omitempty"`
	Status Status      `json:"status,omitempty"`
	URL    string      `json:"url,omitempty"`
}

type Status struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
