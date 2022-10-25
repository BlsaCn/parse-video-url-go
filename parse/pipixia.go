package parse

import (
	"github.com/BlsaCn/parse-video-url-go/tools"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

const piPiXiaUrl = "https://is.snssdk.com/bds/cell/detail/?cell_type=1&aid=1319&app_name=super&cell_id="

type piPiXia struct {
}

func (t piPiXia) ParseShareUr(shareUrl string) (*ParseInfo, error) {
	parse, err := url.Parse(tools.AfterLocationUrl(shareUrl))
	if err != nil {
		return nil, err
	}
	// 解析视频id
	videoId := strings.Replace(parse.Path, "/item/", "", 1)
	return t.ParseVideoId(videoId)
}

func (t piPiXia) ParseVideoId(videoId string) (*ParseInfo, error) {
	dataStr, err := tools.NewHttp().SendHttp(piPiXiaUrl + videoId)
	if err != nil {
		return nil, err
	}
	data := gjson.Get(dataStr, "data.data.item")
	var p = &ParseInfo{
		AuthorName:   data.Get("author.name").String(),
		AuthorAvatar: data.Get("author.avatar.download_list.1.url").String(),
		Title:        data.Get("share.title").String(),
		VideoUrl:     data.Get("origin_video_download.url_list.0.url").String(),
		CoverUrl:     data.Get("origin_video_download.cover_image.download_list.0.url").String(),
		CommentNum:   data.Get("stats.comment_count").Int(),
		LikeNum:      data.Get("stats.like_count").Int(),
	}
	return p, nil
}
