package service

import (
	"fmt"
	"time"
)

type Options struct {
	IP             string
	Port           int
	ServiceName    string
	ServiceVersion string
	MetricsAddr    string
	Uuid           string
	TTL            time.Duration
}

func (o Options) laddr() string {
	return fmt.Sprintf("%s:%d", o.IP, o.Port)
}
