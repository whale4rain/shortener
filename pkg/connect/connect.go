package connect

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// client 全局http客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url可不可以请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
