package singularity

type Host struct {
  Name string
  errCount int
}

type Movable interface {
  getName() string
  PerformMoveOn(* Server) (* Move, error)
}
