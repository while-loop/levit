package registry

import (
	"fmt"
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEtcdRegistry_Deregister(t *testing.T) {
	reg := newEtcd()
	srvcs, err := reg.GetServices()
	require.NoError(t, err)
	fmt.Printf("%v", srvcs)
	for _, srv := range srvcs {
		_, err := reg.cli.Delete(context.TODO(), srv.Key())
		assert.NoError(t, err)
	}
}
