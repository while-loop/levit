package registry

import (
	"strconv"
	"testing"

	"fmt"

	"github.com/stretchr/testify/require"
)

func TestStubImpl(t *testing.T) {
	s := new(stubReg)
	require.Implements(t, (*Registry)(nil), s)
	require.Equal(t, "stub", s.Name())
}

func TestStubReg_Register(t *testing.T) {
	a := require.New(t)
	Use(NewStub())

	srvc := createService("0")
	a.NoError(Register(srvc))

	actual, err := GetService("service0", "0")
	a.NoError(err)
	a.Equal(srvc, actual)

	for i := 1; i <= 5; i++ {
		idx := fmt.Sprintf("%d", i)
		_, err := GetService("service"+idx, idx)
		a.EqualError(err, ErrServiceDNE.Error())
	}
}

func TestStubReg_Deregister(t *testing.T) {
	a := require.New(t)
	s := &stubReg{services: map[string]Service{}}

	srvc := createService("0")
	a.NoError(s.Register(srvc))
	a.NoError(s.Deregister(srvc))

	srvcs, err := s.GetServices()
	a.NoError(err)
	a.Len(srvcs, 0)
	a.Len(s.services, 0)
	a.NoError(s.Deregister(srvc))
	a.NoError(s.Deregister(srvc))
}

func TestStubRemovesMapOnLastInstance(t *testing.T) {
	a := require.New(t)
	Use(NewStub())

	srvc0 := createService("0")
	a.NoError(Register(srvc0))
	createInstance(&srvc0, 1)
	createInstance(&srvc0, 2)
	a.NoError(Register(srvc0))

	srvc1 := createService("1")
	a.NoError(Register(srvc1))

	srvcs, err := GetServices()
	a.NoError(err)
	a.Len(srvcs, 2)
	a.Len(srvcs[0].Instances, 3)

	a.NoError(Deregister(srvc1))
	srvcs, err = GetServices()
	a.NoError(err)
	a.Len(srvcs, 1)
	a.Len(srvcs[0].Instances, 3)

	a.NoError(Deregister(srvc0))
	srvcs, err = GetServices()
	a.NoError(err)
	a.Len(srvcs, 0)

	a.Equal("stub", Name())
}

func createService(index string) Service {
	idx, err := strconv.ParseInt(index, 10, 32)
	if err != nil {
		panic(err)
	}

	return Service{
		Version: index,
		Name:    "service" + index,
		Instances: map[string]Instance{
			"uuid" + index: {
				IP:   "ip" + index,
				Port: int(idx),
				UUID: "uuid" + index,
			},
		},
	}
}

func createInstance(srvc *Service, idx int) {
	index := fmt.Sprintf("%d", idx)
	srvc.AddInstance(Instance{UUID: "uuid" + index, Port: idx, IP: index})

}
