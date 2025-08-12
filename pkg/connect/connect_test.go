package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("base", t, func() {
		url := "https://www.baidu.com"
		got := Get(url)
		// 断言
		convey.So(got, convey.ShouldEqual, true)
	})
	convey.Convey("shouldFalse", t, func() {
		url := "https://www.baidu.com/test"
		got := Get(url)
		// 断言
		convey.So(got, convey.ShouldEqual, false)
	})
}
