package format

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRelTime(t *testing.T) {
	Convey("RelTime", t, func() {
		Convey("just now", func() {
			n := RelTime{time.Now()}
			So(n.String(), ShouldEqual, "just now")
		})
		Convey("1 day ago", func() {
			n := RelTime{time.Now().Add(-24 * time.Hour)}
			So(n.String(), ShouldEqual, "1 day ago")
		})
		Convey("days ago", func() {
			n := RelTime{time.Now().Add(-48 * time.Hour)}
			So(n.String(), ShouldEqual, "2 days ago")
		})
		Convey("minutes from now", func() {
			n := RelTime{time.Now().Add(2 * time.Minute)}
			So(n.String(), ShouldEqual, "2 minutes from now")
		})
	})
}

func ExampleRelTime() {
	t := RelTime{time.Now().Add(-5 * time.Minute)}
	fmt.Println(t)
	// Output: 5 minutes ago
}
func ExampleRelTime_future() {
	t := RelTime{time.Now().Add(5 * time.Minute)}
	fmt.Println(t)
	// Output: 5 minutes from now
}
