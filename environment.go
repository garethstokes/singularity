package singularity

import (
  "math/rand"
  "time"
)

type Environment struct {
  Entities map[string] * Entity
  BoardSize * Vector
}

func NewEnvironment() * Environment {
  environment := new(Environment)
  environment.Entities = make(map[string] * Entity,0)
  environment.BoardSize = &Vector{1000,1000}
  return environment
}

func (e * Environment) RandomScalar() (* Vector) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  v := new(Vector)
  v.X = float64(r.Intn(int(e.BoardSize.X)))
  v.Y = float64(r.Intn(int(e.BoardSize.Y)))
  return v
}

func (e * Environment) RandomNormalisedVector() (* Vector) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  v := new(Vector)
  v.X = (r.Float64() * 2 - 1) * 10
  v.Y = (r.Float64() * 2 - 1) * 10
  return v.Normalise()
}

func (e * Environment) AddPlayer(name string) {
  player := new(Entity)
  player.Name = name

  // position the player on the field randomly
  player.Position = e.RandomScalar()
  player.Direction = e.RandomNormalisedVector()

  e.Entities[name] = player
}

func (e * Environment) Step(playername string, move * Move) {
  player := e.Entities[playername]

  switch move.Action {
    case ACTION_MOVE_FORWARD:
      // extropolate
      n := player.Direction.Normalise().Scale(10)
      player.Position = player.Position.Add(n)
  }
}
