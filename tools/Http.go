package tools

import (
	"io/ioutil"
	"net/http"
)

// AfterLocationUrl 获取重定向后的地址
func AfterLocationUrl(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return url
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cookie", "msToken=VYLbLLi_qqfyoDXRN0iYf9uNwJZnfqA5Y4RI9JsVfiImlPBRUwAFSSAlGwbFHMD045foFL48UvFChaPL86s3wyxKxfzEGbst0nk5KUtjTsd2qHdVY5tS9WU=")
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return url
	}
	return res.Request.URL.String()
}

// Location 获取重定向地址
func Location(url string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(url)
	if err != nil {
		return url
	}
	return res.Header.Get("Location")
}

type Http struct {
	UserAgent string
	Cookie    string
}

func NewHttp() *Http {
	return &Http{
		UserAgent: "",
		Cookie:    "",
	}
}
func (h *Http) SetAgent(userAgent string) *Http {
	h.UserAgent = userAgent
	return h
}
func (h *Http) SetCookie(cookie string) *Http {
	h.Cookie = cookie
	return h
}

// SendHttp 发送http接口请求，获取数据
func (h *Http) SendHttp(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	if len(h.UserAgent) > 0 {
		req.Header.Set("User-Agent", h.UserAgent)
	}
	if len(h.Cookie) > 0 {
		req.Header.Set("Cookie", h.Cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyText), nil
}
