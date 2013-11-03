package singularity

type TickData struct {
  Player * Entity
  VisableThings [](* Entity)
}

type Entity struct {
  Name string         `json:"name"`
  Position * Vector   `json:"position"`
  Direction * Vector  `json:"direction"`
  Speed int           `json:"-"`
  Action MoveAction   `json:"-"`
  Classifier string   `json:"classifier"`
}

type HostTable map[string] Movable

type MoveAction int

type Move struct {
  Direction * Vector
  Action MoveAction
}
