package models

type AdvanceConfig struct {
	// IsRecursionDiff 递归diff json模式
	IsRecursionDiff       bool `json:"is_recursion_diff" bson:"is_recursion_diff"`
	// IsIgnoreArraySequence 忽略json数组类型的顺序
	// 如: [1, 2] == [2, 1]
	IsIgnoreArraySequence bool `json:"is_ignore_array_sequence" bson:"is_ignore_array_sequence"`
}
