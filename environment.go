package singularity

import (
  "math/rand"
  "time"
  "github.com/garethstokes/singularity/log"
)

type Environment struct {
  Entities map[string] * Entity
  BoardSize * Vector
}

func NewEnvironment() * Environment {
  environment := new(Environment)
  environment.Entities = make(map[string] * Entity,0)
  environment.BoardSize = &Vector{800,800}
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

func (e * Environment) AddPlayer(name string, classifier string) {
  player := new(Entity)
  player.Name = name
  player.Classifier = classifier

  // position the player on the field randomly
  player.Position = e.RandomScalar()
  player.Direction = e.RandomNormalisedVector()

  e.Entities[name] = player
}

func (e * Environment) Step(playername string, move * Move) {
  player := e.Entities[playername]
  player.Action = move.Action

  e.stepEntity(player, move)
}

func (e * Environment) stepEntity(entity * Entity, move * Move) {
  switch move.Action {
    case ACTION_MOVE_FORWARD:
      // extropolate
      n := entity.Direction.Normalise().Scale(2)
      log.Infof("n: %@", n)
      entity.Position = entity.Position.Add(n)
    case ACTION_MOVE_BACKWARD:
      n := entity.Direction.Normalise().Scale(10)
      entity.Position = entity.Position.Subtract(n)
    case ACTION_MOVE_TURN:
      d := move.Direction.Normalise()
      entity.Direction = d
    case ACTION_MOVE_STOP:
      // do nothing (regen stamina maybe?)
  }

  e.bounceEntityOfWallsIfNeeded(entity, move);
}

func (e * Environment) bounceEntityOfWallsIfNeeded(entity * Entity, move * Move) {
  //log.Infof("direction: %@", player.Direction)
  //log.Infof("position: %@", player.Position)
  if entity.Position.X < 0 || entity.Position.X > e.BoardSize.X {
    entity.Direction = entity.Direction.NegateX()

    if entity.Position.X < 0 {
      entity.Position.X = 0
    } else {
      entity.Position.X = e.BoardSize.X
    }

    // ...and try again
    e.stepEntity(entity, move)
  }

  if entity.Position.Y < 0 || entity.Position.Y > e.BoardSize.Y {
    entity.Direction = entity.Direction.NegateY()

    if entity.Position.Y < 0 {
      entity.Position.Y = 0
    } else {
      entity.Position.Y = e.BoardSize.Y
    }

    // ...and try again
    e.stepEntity(entity, move)
  }
}

func (e * Environment) PlayerIsOutside(player * Entity) bool {
  if player.Position.X < 0 || player.Position.X > e.BoardSize.X {
    return true
  }

  if player.Position.Y < 0 || player.Position.Y > e.BoardSize.Y {
    return true
  }

  return false;
}
