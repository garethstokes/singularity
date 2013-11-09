package singularity

import (
  "sync"
  "errors"
  "github.com/garethstokes/singularity/log"
)

type HostTable struct {
  sync.Mutex
  rpc map[string] * RpcHost
  all map[string] Movable
}

func NewHostTable() * HostTable {
  ht := new( HostTable )

  ht.all = make( map[string] Movable, 0 )
  ht.rpc = make( map[string] * RpcHost, 0 )

  return ht
}

func (ht * HostTable) AddRpcHost(host * RpcHost) error {
  for name, h := range ht.rpc {

    // do a simple check if someone else is already using 
    // that port
    if host.Address == h.Address {
      return errors.New("ClientAddress is already in use.")
    }

    // maybe the user is already registered and just wants
    // to update their information?
    if name == host.Name {
      log.Infof( "Register Update :: %s", host.Name )
      host.resetErrors()

      ht.rpc[host.Name] = host
      ht.all[host.Name] = host
    }

    return nil
  }

  log.Infof( "Register New :: %s", host.Name )

  ht.rpc[host.Name] = host
  ht.all[host.Name] = host

  return nil
}

func (ht * HostTable) AddMovableHostOnly(host Movable) {
  ht.all[host.getName()] = host
}

func (ht * HostTable) RemoveHostByName(name string) {
  log.Info("acquiring lock :: HostTable")
  ht.Lock()

  log.Infof( "Removing %s from hosts table", name )
  delete(ht.all, name)

  _, ok := ht.rpc[name]
  if ok {
    log.Infof( "Removing %s from rpc hosts table", name )
    delete(ht.rpc, name)
  }

  log.Info("releasing lock :: HostTable")
  ht.Unlock()
}
