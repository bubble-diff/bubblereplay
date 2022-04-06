package models

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

type FilterConfig struct {
	// HttpPathRegexFilter 正则过滤Http请求路径
	HttpPathRegexFilter []string `json:"http_path_regex_filter" bson:"http_path_regex_filter"`
}

// Drop 返回true时，应过滤掉该包。
func (fc FilterConfig) Drop(ctx context.Context, req *http.Request) bool {
	// 无过滤配置时，默认不过滤
	if len(fc.HttpPathRegexFilter) == 0 {
		return false
	}

	for _, pathRegex := range fc.HttpPathRegexFilter {
		matched, err := regexp.MatchString(pathRegex, req.URL.Path)
		if err != nil {
			log.Printf("[Drop] incorrect regular expression => %s, path=%s, caused by %s", pathRegex, req.URL.Path, err)
		}
		if matched {
			return false
		}
	}
	return true
}
