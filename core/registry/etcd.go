package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"net"
	"strings"
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

func NewEtcdRegistry(ctx context.Context, host []string) (Registrar, error) {
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

func (e *Etcd) RegistryService(service *Service) error {
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

	value, err := json.Marshal(service)
	if err != nil {
		return err
	}

	_, err = e.kv.Put(e.ctx, fmt.Sprintf("studio/%s/%s", service.Name, service.Id), string(value), clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}

	hb, err := e.client.KeepAlive(e.ctx, clientv3.LeaseID(grant.ID))

	go func() {
		for {
			select {
			case _, ok := <-hb:
				if !ok {
					return
				}
			case <-e.ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (e *Etcd) UnRegistryService(service *Service) error {
	_, err := e.client.Delete(context.TODO(), fmt.Sprintf("sutdio/%s/%s", service.Name, service.Id))
	return err
}

func (e *Etcd) ListServices() (map[string][]*Service, error) {
	resp, err := e.client.Get(context.TODO(), "studio", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var v = make(map[string][]*Service)

	for _, kv := range resp.Kvs {
		var s Service
		err := json.Unmarshal(kv.Value, &s)
		if err != nil {
			return nil, err
		}
		k := string(kv.Key)[:strings.LastIndex(string(kv.Key), "/")]
		v[k] = append(v[k], &s)
	}

	return v, nil
}

func (e *Etcd) GetService(serviceName string) ([]*Service, error) {
	resp, err := e.client.Get(context.TODO(), fmt.Sprintf("studio/%s", serviceName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var svcs = make([]*Service, 0)
	for _, kv := range resp.Kvs {
		var svc Service
		err := json.Unmarshal(kv.Value, &svc)
		if err != nil {
			return nil, err
		}
		svcs = append(svcs, &svc)
	}
	return svcs, nil
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
