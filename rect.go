package singularity

import (
  "github.com/garethstokes/singularity/log"
)

type Rect struct {
  Size * Size         `json:"size"`
  Position * Vector   `json:"position"`
}

/*
 * OverlapsWith
 *
 * - If the left side of v1 is to the right of v2
 *   then they don't overlap
 *
 * - If the right side of v2 is to the left of v1
 *   then they don't overlap
 *
 * - If the top of v1 is under the bottom of v2 
 *   then they don't overlap
 *
 * - If the bottom of v1 is above the top of v2
 *   then they don't overlap
 *
 */
func (r1 Rect) OverlapsWith(r2 Rect) bool {
  if r1.Position.X > (r2.Position.X + float64(r2.Size.Width)) {
    //log.Info("a")
    return false
  }

  if (r1.Position.X + float64(r1.Size.Width)) < r1.Position.X {
    //log.Info("b")
    return false
  }

  if r1.Position.Y > (r2.Position.Y + float64(r2.Size.Height)) {
    //log.Info("c")
    return false
  }

  if (r1.Position.Y + float64(r1.Size.Height)) < r2.Position.Y {
    //log.Info("d")
    return false
  }

  log.Infof("COLLIDE: %@, %@", r1, r2)
  return true
}
