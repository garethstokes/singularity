package singularity

import (
  "math"
)

const (
  tolerance = 0.0001
)

type Vector struct {
  X float64
  Y float64
}

func (v1 * Vector) Add(v2 * Vector) (* Vector) {
  return &Vector{
    math.Floor(v1.X + v2.X),
    math.Floor(v1.Y + v2.Y),
  }
}

func (v1 * Vector) Subtract(v2 * Vector) (* Vector) {
  return &Vector{
    math.Floor(v1.X - v2.X),
    math.Floor(v1.Y - v2.Y),
  }
}

func (v * Vector) Length() float64 {
  return (v.X * v.X) + (v.Y * v.Y)
}

func (v * Vector) Magnitude() float64 {
  return math.Sqrt(v.Length())
}

/*
 * converts the vector to a unit vector
 */
func (v * Vector) Normalise() (* Vector) {
  m := math.Sqrt( (v.X * v.X) + (v.Y * v.Y) )
  if m < tolerance {
    m = 1
  }

  v.X = v.X / m
  v.Y = v.Y / m

  if math.Abs(v.X) < tolerance {
    v.X = 0.0
  }

  if math.Abs(v.Y) < tolerance {
    v.Y = 0.0
  }

  return v
}

func (v * Vector) Scale(factor float64) (* Vector) {
  result := new(Vector)
  result.X = v.X * factor
  result.Y = v.X * factor
  return result
}
