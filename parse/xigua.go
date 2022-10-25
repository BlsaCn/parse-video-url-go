package parse

import (
	"encoding/base64"
	"fmt"
	"github.com/BlsaCn/parse-video-url-go/tools"
	"github.com/tidwall/gjson"
	"net/url"
	"regexp"
	"strings"
)

type xiGua struct {
}

func (t xiGua) ParseShareUr(shareUrl string) (*ParseInfo, error) {
	parse, err := url.Parse(tools.Location(shareUrl))
	if err != nil {
		return nil, err
	}
	// 解析视频id
	videoId := strings.Replace(parse.Path, "/video/", "", 1)
	videoId = strings.Replace(videoId, "/", "", 1)
	return t.ParseVideoId(videoId)
}

func (t xiGua) ParseVideoId(videoId string) (*ParseInfo, error) {
	dataStr, err := tools.NewHttp().SendHttp(fmt.Sprintf("https://m.365yg.com/i%s/info/", videoId))
	if err != nil {
		return nil, err
	}

	data := gjson.Get(dataStr, "data")
	var p = &ParseInfo{
		AuthorName:   data.Get("media_user.screen_name").String(),
		AuthorAvatar: data.Get("media_user.avatar_url").String(),
		Title:        data.Get("title").String(),
		VideoUrl:     t.getVideoUrl(videoId),
		CoverUrl:     data.Get("poster_url").String(),
		CommentNum:   data.Get("comment_count").Int(),
		LikeNum:      data.Get("digg_count").Int(),
	}
	return p, nil
}

func (t xiGua) getVideoUrl(videoId string) string {
	cookie := "MONITOR_WEB_ID=7143612877764527652; __ac_nonce=06354165e0031397cdbb2; __ac_signature=_02B4Z6wo00f01fMszzAAAIDAeGdU0ds5es3zDMuAAB-hLkcyw3fUcsBLTtRrE0F5G49ooKwZz6ndN47fnhZx9zSVC6fgul9Gm0gDXVB77WhH564acTz3U67bXgn2Ve2-vAbsBDgsLMUWo2ive3; _tea_utm_cache_1768={%22utm_source%22:%22copy_link%22%2C%22utm_medium%22:%22android%22%2C%22utm_campaign%22:%22client_share%22}; _tea_utm_cache_1300={%22utm_source%22:%22copy_link%22%2C%22utm_medium%22:%22android%22%2C%22utm_campaign%22:%22client_share%22}; _tea_utm_cache_2285={%22utm_source%22:%22copy_link%22%2C%22utm_medium%22:%22android%22%2C%22utm_campaign%22:%22client_share%22}; ixigua-a-s=1; tt_scid=1RzrmUQ51j8q5QQYOPq-5V6RvaJewfqGuGfAF6294SbxsJGjStTZsFiOEL.YpAp.3e2c; ttwid=1%7Ccu9m9yb45Ydazbt1ZywzV5oW-kcMjhLOL6wl3BCCLfw%7C1666455550%7C46d6198ce31b93fe21f5e4ead3928e9316c6a87ab6691c696d9eaf0e33ef665a;"
	data, err := tools.NewHttp().
		SetAgent(tools.WIN).
		SetCookie(cookie).
		SendHttp(fmt.Sprintf("https://www.ixigua.com/%s", videoId))
	if err != nil {
		return ""
	}

	reg := regexp.MustCompile(`"main_url":"(.*?)"`)
	res := reg.FindStringSubmatch(data)
	if len(res) == 2 {
		str, err := base64.StdEncoding.DecodeString(res[1])
		if err != nil {
			return ""
		}
		return string(str)
	}

	return ""
}
