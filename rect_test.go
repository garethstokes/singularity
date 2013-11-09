package singularity

import (
  "testing"
)

func makeRect(position * Vector) Rect {
  rect := Rect{}
  rect.Size = &Size{ 25, 25 }
  rect.Position = position
  return rect
}

func TestOverlapsWhenOnSamePosition(t * testing.T) {
  r1 := makeRect(&Vector{ 100, 100 })
  r2 := makeRect(&Vector{ 100, 100 })

  if r1.OverlapsWith(r2) == false {
    t.Fatalf("on same space, should have overlaped")
  }
}

func TestOverlapsWhenWithSize(t * testing.T) {
  r1 := makeRect(&Vector{ 105, 115 })
  r2 := makeRect(&Vector{ 100, 110 })

  if r1.OverlapsWith(r2) == false {
    t.Fatalf("should have overlaped")
  }
}

func TestNotOverlapingA(t * testing.T) {
  r1 := makeRect(&Vector{ 100, 100 })
  r2 := makeRect(&Vector{ 200, 200 })

  if r1.OverlapsWith(r2) {
    t.Fatalf("should not have overlaped")
  }
}

func TestNotOverlapingB(t * testing.T) {
  r1 := makeRect(&Vector{ 200, 100 })
  r2 := makeRect(&Vector{ 100, 200 })

  if r1.OverlapsWith(r2) {
    t.Fatalf("should not have overlaped")
  }
}

func TestNotOverlapingC(t * testing.T) {
  r1 := makeRect(&Vector{ 131, 599 })
  r2 := makeRect(&Vector{ 509, 577 })

  if r1.OverlapsWith(r2) == false {
    t.Fatalf("should have overlaped")
  }
}
