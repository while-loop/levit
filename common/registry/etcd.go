package registry

import (
	"context"
	"time"

	"encoding/json"

	client "github.com/coreos/etcd/clientv3"
	"github.com/while-loop/levit/common/log"
)

const (
	MasterAddress = "http://anthonyalves.science:2379"
)

type EtcdRegistry struct {
	cli *client.Client
}

func NewEtcd() Registry {
	return newEtcd()
}

func newEtcd() *EtcdRegistry {
	cfg := client.Config{
		Endpoints:            []string{MasterAddress},
		AutoSyncInterval:     30 * time.Second,
		DialKeepAliveTimeout: 5 * time.Second,
		DialKeepAliveTime:    5 * time.Second,
		DialTimeout:          5 * time.Second,
	}
	cli, err := client.New(cfg)
	retries := 0
	for err != nil && retries <= 20 {
		log.Error("Failed to start Etcd registry client ", err)
		time.Sleep(2 * time.Second)
		cli, err = client.New(cfg)
		retries++
	}

	if cli == nil {
		log.Fatal("Failed to start Etcd registry client")
	}

	log.Info("Starting etcd registry")
	return &EtcdRegistry{
		cli: cli,
	}
}

func (e *EtcdRegistry) Register(srvc Service) error {
	log.Debug("registering service ", srvc.Key())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := e.cli.Grant(ctx, int64(srvc.TTL.Seconds()))
	if err != nil {
		return err
	}
	leaseID := resp.ID

	err = nil
	bs, _ := json.Marshal(srvc)
	_, err = e.cli.Put(ctx, srvc.Key(), string(bs), client.WithLease(leaseID))

	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdRegistry) Deregister(srvc Service) error {
	log.Debug("deregistering service ", srvc.Key())
	_, err := e.cli.Delete(context.Background(), srvc.Key())
	return err
}

func (e *EtcdRegistry) GetServices() ([]Service, error) {
	resp, err := e.cli.Get(context.Background(), "", client.WithPrefix())
	if err != nil {
		return nil, err
	}

	srvcs := make([]Service, 0)
	for _, kv := range resp.Kvs {
		var srvc Service
		err := json.Unmarshal(kv.Value, &srvc)
		if err != nil {
			log.Errorf("Failed to unmarshal service struct %s: %v\n%s", string(kv.Key), err, string(kv.Value))
			continue
		}

		srvcs = append(srvcs, srvc)
	}

	return srvcs, nil
}

func (e *EtcdRegistry) GetService(serviceName, version string) ([]Service, error) {
	resp, err := e.cli.Get(context.Background(), serviceName+"-"+version, client.WithPrefix())
	if err != nil {
		return nil, err
	}

	if resp.Count <= 0 {
		return nil, ErrServiceDNE
	}

	srvcs := make([]Service, 0)
	var srvc Service
	for _, kv := range resp.Kvs {
		err := json.Unmarshal(kv.Value, &srvc)
		if err != nil {
			log.Errorf("Failed to unmarshal service struct %s: %v\n%s", string(kv.Key), err, string(kv.Value))
			continue
		}

		srvcs = append(srvcs, srvc)
	}

	return srvcs, nil
}

func (e *EtcdRegistry) Name() string {
	return "etcd"
}
