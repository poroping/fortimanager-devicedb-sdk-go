package models

type TaskResultData struct {
	Adom       *int64  `json:"adom,omitempty"`
	EndTm      *int64  `json:"end_tm,omitempty"`
	Flags      *int64  `json:"flags,omitempty"`
	ID         *int64  `json:"id,omitempty"`
	Line       []Line  `json:"line,omitempty"`
	NumDone    *int64  `json:"num_done,omitempty"`
	NumErr     *int64  `json:"num_err,omitempty"`
	NumLines   *int64  `json:"num_lines,omitempty"`
	NumWarn    *int64  `json:"num_warn,omitempty"`
	Percent    *int64  `json:"percent,omitempty"`
	PID        *int64  `json:"pid,omitempty"`
	Src        *string `json:"src,omitempty"`
	StartTm    *int64  `json:"start_tm,omitempty"`
	State      *string `json:"state,omitempty"`
	Title      *string `json:"title,omitempty"`
	TotPercent *int64  `json:"tot_percent,omitempty"`
	User       *string `json:"user,omitempty"`
}

type Line struct {
	Detail  *string   `json:"detail,omitempty"`
	EndTm   *int64    `json:"end_tm,omitempty"`
	Err     *int64    `json:"err,omitempty"`
	History []History `json:"history,omitempty"`
	IP      *string   `json:"ip,omitempty"`
	Name    *string   `json:"name,omitempty"`
	OID     *int64    `json:"oid,omitempty"`
	Percent *int64    `json:"percent,omitempty"`
	POID    *int64    `json:"poid,omitempty"`
	StartTm *int64    `json:"start_tm,omitempty"`
	State   *string   `json:"state,omitempty"`
	Vdom    *string   `json:"vdom,omitempty"`
}

type History struct {
	Detail  *string     `json:"detail,omitempty"`
	Name    *string     `json:"name,omitempty"`
	Percent *int64      `json:"percent,omitempty"`
	State   *int64      `json:"state,omitempty"`
	Vdom    interface{} `json:"vdom"`
}
