package proteus

import (
	"testing"
)

func checkValue(t *testing.T, val1 interface{}, val2 interface{}) {
	if val1 != val2 {
		t.Errorf("Values are not equal: '%s' != '%s'", val1, val2)
	}
}

type Src struct {
	ValA string `dst:"Val1"`
	ValB string `dst:"Val2"`
}

type Dst struct {
	Val1 string
	Val2 string
}

func TestStructMap(t *testing.T) {
	src := Src{
		ValA: "value from A",
		ValB: "value from B",
	}
	dst := Dst{}
	New("dst").Map(src, &dst)
	checkValue(t, dst.Val1, src.ValA)
	checkValue(t, dst.Val2, src.ValB)
	dst = Dst{}
	New("dst").Map(&src, &dst)
	checkValue(t, dst.Val1, src.ValA)
	checkValue(t, dst.Val2, src.ValB)
}

type SrcEmbedded struct {
	Src `dst:""`
}

func TestEmbeddedStruct(t *testing.T) {
	src := SrcEmbedded{}
	src.ValA = "value from A"
	src.ValB = "value from B"
	dst := Dst{}
	New("dst").Map(src, &dst)
	checkValue(t, dst.Val1, src.ValA)
	checkValue(t, dst.Val2, src.ValB)
}

type SrcNested struct {
	SrcInner Src `dst:"DstInner"`
}

type DstNested struct {
	DstInner Dst
}

func TestNestedStructs(t *testing.T) {
	src := SrcNested{}
	src.SrcInner.ValA = "value from A"
	src.SrcInner.ValB = "value from B"
	dst := DstNested{}
	New("dst").Map(src, &dst)
	checkValue(t, dst.DstInner.Val1, src.SrcInner.ValA)
	checkValue(t, dst.DstInner.Val2, src.SrcInner.ValB)
}

type SrcEmbedded2 struct {
	Src `dst:"DstInner"`
}

func TestEmbeddedToNested(t *testing.T) {
	src := SrcEmbedded2{}
	src.ValA = "value from A"
	src.ValB = "value from B"
	dst := DstNested{}
	New("dst").Map(src, &dst)
	checkValue(t, dst.DstInner.Val1, src.ValA)
	checkValue(t, dst.DstInner.Val2, src.ValB)
}

type SrcNumbers struct {
	N1 int64 `dst:"No1"`
	N2 int32 `dst:"No1"`
}

type DstNumbers struct {
	No1 int64
	No2 int64
}

func TestAssignableTypes(t *testing.T) {
	src := SrcNumbers{N1: 1, N2: 3}
	dst := DstNumbers{}
	New("dst").Map(src, &dst)
	checkValue(t, src.N1, dst.No1)
	checkValue(t, int64(0), dst.No2)
}
