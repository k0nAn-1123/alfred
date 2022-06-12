package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int
}

func TestJsonMarshal(t *testing.T) {
	Convey("Testing JsonMarshal", t, func() {
		case1 := TestStruct{
			Name: "Jack",
			Age:  12,
		}
		str, err := JsonMarshal(case1)
		So(str, ShouldEqual, "{\"Name\":\"Jack\",\"Age\":12}")
		So(err, ShouldBeNil)
	})
}

func TestJsonUnmarshal(t *testing.T) {
	Convey("Testing JsonUnmarshal", t, func() {
		str := "{\"Name\":\"Jack\",\"Age\":12}"
		case1 := &TestStruct{}
		err := JsonUnmarshal(str, case1)
		So(case1.Name, ShouldEqual, "Jack")
		So(case1.Age, ShouldEqual, 12)
		So(err, ShouldBeNil)
	})
}
