package parse

import (
	"errors"
	"parse-video-url-go/tools"
	"strings"
)

// ByMsg 是否为视频分享链接
func ByMsg(msg string) (*ParseInfo, error) {
	for source, info := range VideoInfoMap {
		url := tools.MatchUrl(info.Pattern, msg)
		if len(url) > 0 {
			if !strings.HasPrefix(url, "http") {
				url = "https://" + url
			}
			return VideoInfoMap[source].UrlParse.ParseShareUr(url)
		}
	}

	return nil, errors.New("匹配不到视频平台")
}
