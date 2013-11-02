package singularity

type Host struct {
  Name string
  Address string
  errCount int
}

type Movable interface {
  PerformMoveOn(* Server) (* Move, error)
  getName() string
  getAddress() string
  resetErrors()
}

