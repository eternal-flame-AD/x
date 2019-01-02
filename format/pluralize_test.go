package format

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPluralize(t *testing.T) {
	Convey("Pluralize", t, func() {
		Convey("regular", func() {
			So(AutoPluralize(1, "car"), ShouldEqual, "1 car")
			So(AutoPluralize(2, "car"), ShouldEqual, "2 cars")
		})
		Convey("special", func() {
			So(AutoPluralize(1, "knife"), ShouldEqual, "1 knife")
			So(AutoPluralize(2, "knife"), ShouldEqual, "2 knives")
			So(AutoPluralize(1, "sheep"), ShouldEqual, "1 sheep")
			So(AutoPluralize(2, "sheep"), ShouldEqual, "2 sheep")
		})
	})
}
