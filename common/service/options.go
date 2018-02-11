package service

import (
	"fmt"

	"github.com/google/uuid"
	levnet "github.com/while-loop/levit/common/net"
)

type Options struct {
	IP             string
	Port           int
	ServiceName    string
	ServiceVersion string
	MetricsAddr    string
	UUID           string
}

func (o *Options) applyDefaults() {
	if o.IP == "" {
		o.IP = levnet.GetIP()
	}

	if o.Port <= 0 {
		o.Port = 8080
	}

	o.UUID = uuid.New().String()
}
func (o Options) laddr() string {
	return fmt.Sprintf("%s:%d", o.IP, o.Port)
}
