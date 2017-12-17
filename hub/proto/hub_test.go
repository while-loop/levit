package proto

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	m := &HubMessage{Event: &HubMessage_EventChannelSeen{}}

	var handlers = map[reflect.Type]bool{
		reflect.TypeOf(&HubMessage_EventChannelSeen{}): true,
	}

	u := reflect.TypeOf(m.Event)
	assert.True(t, handlers[u])
}
