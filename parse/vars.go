package parse

const (
	PiPiXia = "pipixia"
	Douyin  = "douyin"
	XiGua   = "xigua"
)

type urlParse interface {
	// ParseShareUr 解析分享的url
	ParseShareUr(shareUrl string) (*ParseInfo, error)
}
type videoIdParse interface {
	ParseVideoId(videoId string) (*ParseInfo, error)
}

// ParseInfo 最终返回的结构体信息
type ParseInfo struct {
	AuthorName   string `json:"author_name"`   // 作者名称
	AuthorAvatar string `json:"author_avatar"` // 作者头像地址
	Title        string `json:"title"`         // 视频标题
	VideoUrl     string `json:"video_url"`     // 视频播放地址
	CoverUrl     string `json:"cover_url"`     // 视频封面地址
	CommentNum   int64  `json:"comment_num"`   // 评论数量
	LikeNum      int64  `json:"like_num"`      // 点赞数量
}

type videoInfo struct {
	VideoDomain  string // 分享视频的域名
	Pattern      string // 匹配视频链接的正则表达式
	UrlParse     urlParse
	videoIdParse videoIdParse
}

var VideoInfoMap = map[string]videoInfo{
	PiPiXia: {
		VideoDomain:  "h5.pipix.com",
		Pattern:      `.*?(h5.pipix.com/s\/[0-9a-zA-Z]{7}\/?).*?`,
		UrlParse:     piPiXia{},
		videoIdParse: piPiXia{},
	},
	Douyin: {
		VideoDomain:  "v.douyin.com",
		Pattern:      `.*?(v.douyin.com\/[0-9a-zA-Z]{7}\/?).*?`,
		UrlParse:     &douYin{},
		videoIdParse: douYin{},
	},
	XiGua: {
		VideoDomain:  "v.ixigua.com",
		Pattern:      `.*?(v.ixigua.com\/[0-9a-zA-Z]{7}\/?).*?`,
		UrlParse:     xiGua{},
		videoIdParse: xiGua{},
	},
}
