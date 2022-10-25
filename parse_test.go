package main

import (
	"github.com/BlsaCn/parse-video-url-go/parse"
	"testing"
)

func TestByMsg(t *testing.T) {
	info, _ := parse.ByMsg("https://v.ixigua.com/Mm8mnoy 打开西瓜视频")
	if info.AuthorName == "" {
		t.Fatal("AuthorName fail")
	}
	if info.AuthorAvatar == "" {
		t.Fatal("AuthorAvatar fail")
	}
	if info.Title == "" {
		t.Fatal("Title fail")
	}
	if info.VideoUrl == "" {
		t.Fatal("VideoUrl fail")
	}
	if info.CoverUrl == "" {
		t.Fatal("CoverUrl fail")
	}
	if info.CommentNum == 0 {
		t.Fatal("CommentNum fail")
	}
	if info.LikeNum == 0 {
		t.Fatal("LikeNum fail")
	}

}

func TestByMsg2(t *testing.T) {
	info, _ := parse.ByMsg("https://v.ixigua.com/1Mm8mnoy 打开西瓜视频")
	if info.AuthorName != "" {
		t.Fatal("AuthorName fail")
	}
	if info.AuthorAvatar != "" {
		t.Fatal("AuthorAvatar fail")
	}
	if info.Title != "" {
		t.Fatal("Title fail")
	}
	if info.VideoUrl != "" {
		t.Fatal("VideoUrl fail")
	}
	if info.CoverUrl != "" {
		t.Fatal("CoverUrl fail")
	}
	if info.CommentNum != 0 {
		t.Fatal("CommentNum fail")
	}
	if info.LikeNum != 0 {
		t.Fatal("LikeNum fail")
	}
}
