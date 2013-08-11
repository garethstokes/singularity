package singularity_test

import (
  "testing"
  "github.com/garethstokes/singularity"
)

func TestAddPlayer(t * testing.T) {
  e := singularity.NewEnvironment()
  e.AddPlayer("garrydanger")

  player := e.Entities["garrydanger"]
  if player == nil {
    t.Fatal("null player found.")
  }

  if player.Position.X < 0 || player.Position.X > e.BoardSize.X {
    t.Fatal("player x position out of bounds.")
  }

  if player.Position.Y < 0 || player.Position.Y > e.BoardSize.Y {
    t.Fatal("player x position out of bounds.")
  }
}

func TextPlayerCantExceedBoundry(t * testing.T) {
  e := singularity.NewEnvironment()
  e.AddPlayer("garrydanger")
  player := e.Entities["garrydanger"]

  if player == nil {
    t.Fatal("null player found.")
  }

  player.Position.X = -10

  if e.PlayerIsOutside(player) == false {
    t.Fatal("player x position out of bounds")
  }

  player.Position.X = 10
  player.Position.Y = -10

  if e.PlayerIsOutside(player) == false {
    t.Fatal("player y position out of bounds")
  }

  player.Position.X = e.BoardSize.X + 10
  player.Position.Y = 10

  if e.PlayerIsOutside(player) == false {
    t.Fatal("player x position out of bounds")
  }

  player.Position.X = 10
  player.Position.Y = e.BoardSize.Y + 10

  if e.PlayerIsOutside(player) == false {
    t.Fatal("player y position out of bounds")
  }
}
