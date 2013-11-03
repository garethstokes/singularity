package singularity

type MemoryHost struct {
  Host
}

func (host * MemoryHost) getName() string {
  return host.Name
}


func (host * MemoryHost) PerformMoveOn(s * Server) (* Move, error) {
  result := new(Move)
  result.Direction = &Vector{0,0}
  result.Action = ACTION_MOVE_FORWARD

  return result, nil
}
