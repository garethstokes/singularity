package singularity

import (
  "math"
)

type Vector struct {
  X float64   `json:"x"`
  Y float64   `json:"y"`
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

func (v1 * Vector) Multiply(s float64) (* Vector) {
  return &Vector{
    v1.X * s,
    v1.Y * s,
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
  return v.Multiply(1.0/v.Length())
}

func (v * Vector) Scale(factor float64) (* Vector) {
  result := new(Vector)
  result.X = v.X * factor
  result.Y = v.Y * factor
  return result
}

func (v * Vector) Negate() (* Vector) {
  return &Vector{-v.X,-v.Y}
}

func (v * Vector) NegateX() (* Vector) {
  return &Vector{-v.X,v.Y}
}

func (v * Vector) NegateY() (* Vector) {
  return &Vector{v.X,-v.Y}
}
