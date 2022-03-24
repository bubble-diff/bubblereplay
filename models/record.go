package models

type Record struct {
	TaskID  int64  `json:"task_id"`
	OldReq  []byte `json:"old_req"`
	OldResp []byte `json:"old_resp"`
	NewResp []byte `json:"new_resp"`
	// Diff 差异比对结果，形式为EBNF表达式。
	Diff    string `json:"diff,omitempty"`
	// DiffRate 差异比对率，基于Levenshtein算法给出。
	DiffRate float64 `json:"diff_rate"`
}
