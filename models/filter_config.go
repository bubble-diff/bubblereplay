package models

type FilterConfig struct {
	// HttpPathRegexFilter 正则过滤Http请求路径
	HttpPathRegexFilter []string `json:"http_path_regex_filter" bson:"http_path_regex_filter"`
}
