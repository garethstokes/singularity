package singularity

type MemoryHost struct {
  Host
}

func (host * MemoryHost) getName() string {
  return host.Name
}


func (host * RpcHost) PerformMoveOn(s * Server) (* Move, error) {
  return nil, nil
}