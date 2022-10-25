package main

import (
	"github.com/BlsaCn/parse-video-url-go/parse"
	"gopkg.in/ffmt.v1"
)

func main() {
	info, _ := parse.ByMsg("https://h5.pipix.com/s/MBxwbnj 打开皮皮虾")
	// info, _ := parse.ByMsg("https://v.douyin.com/MhPKjN2/ 打开抖音)
	// info, _ := parse.ByMsg("https://v.ixigua.com/Mm8mnoy 打开西瓜视频")

	// 头条
	// var_dump(Parse::TouTiao('https://m.toutiao.com/is/Mu13Xqu/'));

	// 哔哩哔哩
	// var_dump(Parse::BiLiBiLi('https://b23.tv/3uPg39'));

	_, _ = ffmt.P(info)
}
