<h1>BlsaCn/parse-video-url-go</h1>
<p>

</p>

## 短视频去水印

集成了：皮皮虾、抖音、西瓜、头条、哔哩哔哩、等等

## 安装

~~~
go get github.com/BlsaCn/parse-video-url-go
~~~

使用：(可以参考parse_test.go)
==

    匹配平台：ByMsg($msg)

成功返回

```
[
    'authorName' => '',  // 作者名称
    'authorAvatar' => '',// 作者头像地址
    'title' => '',       // 视频标题
    'videoUrl' => '',    // 视频播放地址
    'coverUrl' => '',    // 视频封面地址
    'commentNum' => 0,   // 评论数量
    'likeNum' => 0,      // 点赞数量
]

&parse-video-url-go/parse.ParseInfo{
 AuthorName:   string("爱母虾的大爷")
 AuthorAvatar: string("https://p9-ppx.byteimg.com/obj/tos-cn-i-8gu37r9deh/2d6e8a8c1a994001950c6e9598b155b4")
 Title:        string("羊会咩咩，鸭会嘎嘎，鸡会什么")
 VideoUrl:     string("http://v6-cdn-tos.ppxvod.com/f490d733a4170048e2ee2e8a462dc68e/6356ba0f/video/tos/cn/tos-cn-ve-0076/658
27b50c65941a4a6eb3809aaea621a/?a=1319&ch=0&cr=0&dr=3&cd=0%7C0%7C0%7C0&cv=1&br=1588&bt=1588&cs=0&ds=3&eid=2048&ft=FIJbvNN6V~6w
bLMFq8dzJLeOYZlc9R2dd2bLjOljtiZm&mime_type=video_mp4&qs=0&rc=Zjw6OWc2Nzk2OzlkNGdlZ0BpanFla2Q6ZnJyZzMzNGYzM0AtMzEwYWIzNjYxYF5j
YWJiYSNuMWlhcjRfLTJgLS1kMWFzcw%3D%3D&l=20221024231500010208102080175DFB98")
 CoverUrl:     string("https://p6-ppx-sign.byteimg.com/tos-cn-p-0076/8151df2c97c0421596496e5b9900d339_1665892503~tplv-f3gpral
wbh-logo.jpeg?x-expires=1698160501&x-signature=i4aZxLcsBFXwOEnbHKV6mLEkNpY%3D")
 CommentNum:   int64(318)
 LikeNum:      int64(5599)
}

```

失败返回

```
&parse-video-url-go/parse.ParseInfo{
}
```

