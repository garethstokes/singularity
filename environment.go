package singularity

import (
  "math/rand"
  "time"
)

type Environment struct {
  Entities map[string] * Entity
  BoardSize * Point
}

func NewEnvironment() * Environment {
  environment := new(Environment)
  environment.Entities = make(map[string] * Entity,0)
  environment.BoardSize = &Point{1000,1000}
  return environment
}

func (e * Environment) AddPlayer(name string) {
  player := new(Entity)
  player.Name = name

  // position the player on the field randomly
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  p := &Point{r.Intn(e.BoardSize.X),r.Intn(e.BoardSize.Y)}
  player.Position = p

  e.Entities[name] = player
}

func (e * Environment) Step(playername string, move * Move) {
  player := e.Entities[playername]

  switch move.Action {
    case ACTION_MOVE_FORWARD:
      // extropolate
      player.Position.X += 10;
  }
}
