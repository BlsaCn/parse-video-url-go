package parse

import (
	"github.com/tidwall/gjson"
	"net/url"
	"parse-video-url-go/tools"
	"strings"
)

const douYinUrl = "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids="

type douYin struct {
}

func (t douYin) ParseShareUr(shareUrl string) (*ParseInfo, error) {
	parse, err := url.Parse(tools.AfterLocationUrl(shareUrl))
	if err != nil {
		return nil, err
	}
	// 解析视频id
	videoId := strings.Replace(parse.Path, "/video/", "", 1)
	return t.ParseVideoId(videoId)
}
func (t douYin) ParseVideoId(videoId string) (*ParseInfo, error) {
	dataStr, err := tools.NewHttp().SendHttp(douYinUrl + videoId)
	if err != nil {
		return nil, err
	}
	data := gjson.Get(dataStr, "item_list.0")
	videoUrl := strings.Replace(data.Get("video.play_addr.url_list.0").String(), "/playwm/", "/play/", 1)
	var p = &ParseInfo{
		AuthorName:   data.Get("author.nickname").String(),
		AuthorAvatar: data.Get("author.avatar_larger.url_list.0").String(),
		Title:        data.Get("desc").String(),
		VideoUrl:     tools.AfterLocationUrl(videoUrl),
		CoverUrl:     data.Get("video.cover.url_list.0").String(),
		CommentNum:   data.Get("statistics.comment_count").Int(),
		LikeNum:      data.Get("statistics.digg_count").Int(),
	}
	return p, nil
}
