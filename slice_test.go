package toolkit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaginateSlice(t *testing.T) {
	Convey("slice paginate", t, func() {
		Convey("slice paginate with invalid param", func() {
			slice := []int{1, 2, 3, 4}
			newSlice := PaginateSlice(slice, 0, 0)
			So(len(newSlice), ShouldEqual, 0)
			newSlice = PaginateSlice(slice, 0, -1)
			So(len(newSlice), ShouldEqual, 0)
			newSlice = PaginateSlice(slice, -1, 1)
			So(len(newSlice), ShouldEqual, 1)
			So(newSlice[0], ShouldEqual, 1)
		})
		Convey("slice paginate", func() {
			slice := []int{1, 2, 3, 4}
			newSlice := PaginateSlice(slice, 1, 2)
			So(newSlice[0], ShouldEqual, 1)
			So(newSlice[1], ShouldEqual, 2)
			So(len(newSlice), ShouldEqual, 2)
			newSlice = PaginateSlice(slice, 2, 2)
			So(newSlice[0], ShouldEqual, 3)
			So(newSlice[1], ShouldEqual, 4)
			So(len(newSlice), ShouldEqual, 2)
		})
		Convey("slice paginate param out of range", func() {
			slice := []int{1, 2, 3, 4}
			newSlice := PaginateSlice(slice, 3, 2)
			So(len(newSlice), ShouldEqual, 0)
		})
	})
}

func TestSliceRemoveDuplication(t *testing.T) {
	Convey("slice remove duplication", t, func() {
		Convey("slice empty", func() {
			intSlice := []int{}
			SliceRemoveDuplication(&intSlice)
			So(len(intSlice), ShouldEqual, 0)
			strSlice := []string{}
			SliceRemoveDuplication(&strSlice)
			So(len(strSlice), ShouldEqual, 0)
		})
		Convey("slice not empty", func() {
			intSlice := []int{1, 2, 3, 3, 1, 2, 4, 4}
			SliceRemoveDuplication(&intSlice)
			So(len(intSlice), ShouldEqual, 4)
			type MyStruct struct {
				A string
				B int
			}
			ms1 := MyStruct{A: "a", B: 1}
			ms2 := MyStruct{A: "b", B: 1}
			ms3 := MyStruct{A: "a", B: 2}
			ms4 := MyStruct{A: "b", B: 2}
			sSlice := []MyStruct{ms1, ms2, ms1, ms3, ms3, ms2, ms4, ms1}
			SliceRemoveDuplication(&sSlice)
			So(len(sSlice), ShouldEqual, 4)
			cpms1 := ms1
			ptrms1 := &ms1
			psSlice := []*MyStruct{&ms1, &ms2, &ms1, &ms3, &ms3, &ms2, &ms4, &cpms1, ptrms1}
			SliceRemoveDuplication(&psSlice)
			So(len(psSlice), ShouldEqual, 5)
		})
	})
}

func TestSliceCompare(t *testing.T) {
	Convey("切片比较", t, func() {
		slice1 := []int{}
		slice2 := []int{}
		So(SliceCompare(slice1, slice2), ShouldBeTrue)
		slice1 = []int{1, 2, 3, 4}
		slice2 = []int{1, 2, 3, 4}
		So(SliceCompare(slice1, slice2), ShouldBeTrue)
		slice2 = []int{3, 1, 2, 4}
		So(SliceCompare(slice1, slice2), ShouldBeTrue)
		slice2 = []int{3, 1, 2, 2, 4}
		So(SliceCompare(slice1, slice2), ShouldBeTrue)
		slice2 = []int{1, 2, 3}
		So(SliceCompare(slice1, slice2), ShouldBeFalse)
		slice2 = []int{1, 2, 3, 5}
		So(SliceCompare(slice1, slice2), ShouldBeFalse)
	})
}
