package singularity

type TickData struct {
  Player * Entity
  VisableThings [](* Entity)
}

type MoveAction int

type Move struct {
  Direction * Vector
  Action MoveAction
}

type Size struct {
  Height int          `json:"height"`
  Width int           `json:"width"`
}

type Entity struct {
  Rect
  Name string         `json:"name"`
  Direction * Vector  `json:"direction"`
  Speed int           `json:"-"`
  Action MoveAction   `json:"-"`
  Classifier string   `json:"classifier"`
  Health int          `json:"health"`
  Layer int           `json:"-"`
}
