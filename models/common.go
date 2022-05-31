package models

const CmdbBasePath = "/pm/config/device/"

// "/pm/config/device/${device}/vdom/${vdom}/"

type CmdbRequest struct {
	ID      int                  `json:"id"`
	Verbose int                  `json:"verbose"`
	Method  string               `json:"method"`
	Params  []CmdbRequestPayload `json:"params"`
	Session string               `json:"session"`
}

type CmdbRequestPayload struct {
	Data interface {
	} `json:"data"`
	URL string `json:"url"`
}

// type CmdbRequestParams struct {
// 	AllowAppend *bool
// 	Device      string
// 	Mkey        string
// 	Vdom        string
// }

type CmdbResponse struct {
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
