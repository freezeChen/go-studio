package registry

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"net"
	"time"
)

const (
	registerTTL = 90
)

type Etcd struct {
	ctx    context.Context
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease

	watch clientv3.Watcher
}

func NewEtcdRegistry(ctx context.Context, host []string) (Registry, error) {
	var registry = new(Etcd)
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   host,
		DialTimeout: 5 * time.Second})
	if err != nil {
		return nil, err
	}
	registry.ctx = ctx
	registry.client = client
	registry.kv = clientv3.NewKV(client)

	return registry, nil
}

func (e *Etcd) RegistryService(service Service) error {
	if e.lease != nil {
		e.lease.Close()
	}
	e.lease = clientv3.NewLease(e.client)

	if e.watch != nil {
		e.watch.Close()
	}
	e.watch = clientv3.NewWatcher(e.client)

	grant, err := e.lease.Grant(e.ctx, registerTTL)
	if err != nil {
		return err
	}

	_, err = e.kv.Put(e.ctx, "studio:"+service.Name, GetLocalIp()+":"+service.Port, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}

	return nil
}

func (e *Etcd) unRegistryService(service Service) error {
	_, err := e.client.Delete(context.TODO(), service.Name)
	return err
}

func (e *Etcd) ListServices() (map[string][]string, error) {
	resp, err := e.client.Get(context.TODO(), "studio", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var v =make(map[string][]string)
	for _, kv := range resp.Kvs {
		v[string(kv.Key)] = append(v[string(kv.Key)], string(kv.Value))
	}

	return v, nil
}

func (e *Etcd) WatchServices() {
	panic("implement me")
}

func GetLocalIp() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}
