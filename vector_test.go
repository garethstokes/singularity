package singularity_test

import (
  "math"
  "testing"
  "github.com/garethstokes/singularity"
)

func TestNormaliseVector(t * testing.T) {
  v := new(singularity.Vector)
  v.X = 100
  v.Y = 100

  v = v.Normalise()

  var unit = float64(1)
  m := v.Magnitude()

  if math.Ceil(m) != unit {
    t.Fatalf("expected %f but found %f\n", unit, m)
  }
}

func TestAddVectors(t * testing.T) {
  v1 := &singularity.Vector {10,10}
  v2 := &singularity.Vector {5,5}

  vector := v1.Add(v2)

  t.Logf("%@\n", vector)


}
