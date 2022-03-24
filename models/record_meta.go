package models

type RecordMeta struct {
	CosKey   string  `json:"cos_key"`
	Path     string  `json:"path"`
	DiffRate float64 `json:"diff_rate"`
}
