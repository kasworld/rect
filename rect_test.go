package rect

import (
	"kasworld/goguelike/lib/ivector2d"
	"testing"
)

func TestIsIn(t *testing.T) {
	r1 := Rect{10, 10, 10, 10}
	r2 := Rect{10, 10, 10, 10}
	if !r1.IsIn(r2) {
		t.Fail()
	}
}
func TestNotIsIn(t *testing.T) {
	r1 := Rect{10, 10, 10, 11}
	r2 := Rect{10, 10, 10, 10}
	if r1.IsIn(r2) {
		t.Fail()
	}
}
func TestIsOverlap(t *testing.T) {
	r1 := Rect{10, 10, 10, 10}
	r2 := Rect{15, 15, 10, 10}
	if !r1.IsOverlap(r2) {
		t.Fail()
	}
}
func TestNotIsOverlap(t *testing.T) {
	r1 := Rect{10, 10, 10, 10}
	r2 := Rect{20, 10, 10, 10}
	if r1.IsOverlap(r2) {
		t.Fail()
	}
}

func TestMakeRectBy4Driect1(t *testing.T) {
	r := Rect{10, 10, 10, 10}
	p := ivector2d.Vt{15, 15}
	rs := [4]Rect{}
	for i := range rs {
		rs[i] = r.MakeRectBy4Driect(p, i)
		if !rs[i].IsIn(r) {
			t.Fail()
		}
	}
}

func TestMakeRectBy4Driect2(t *testing.T) {
	r := Rect{10, 10, 10, 10}
	p := ivector2d.Vt{15, 15}
	rs := [4]Rect{}
	for i := range rs {
		rs[i] = r.MakeRectBy4Driect(p, i)
	}
	if rs[0].IsOverlap(rs[1]) {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	r1 := Rect{10, 10, 10, 10}
	r2 := Rect{15, 15, 10, 10}
	if !(r1.Union(r2) == Rect{10, 10, 15, 15}) {
		t.Logf("result %v", r1.Union(r2))
		t.Fail()
	}
	if !(r2.Union(r1) == Rect{10, 10, 15, 15}) {
		t.Logf("result %v", r2.Union(r1))
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	r1 := Rect{10, 10, 10, 10}
	r2 := Rect{15, 15, 10, 10}
	if !(r1.Intersection(r2) == Rect{15, 15, 5, 5}) {
		t.Logf("result %v", r1.Intersection(r2))
		t.Fail()
	}
	if !(r2.Intersection(r1) == Rect{15, 15, 5, 5}) {
		t.Logf("result %v", r2.Intersection(r1))
		t.Fail()
	}
}
