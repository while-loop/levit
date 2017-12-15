package service

import "fmt"

type Options struct {
	IP             string
	Port           int
	ServiceName    string
	ServiceVersion string
	MetricsAddr    string
	Uuid           string
}

func (o Options) laddr() string {
	return fmt.Sprintf("%s:%d", o.IP, o.Port)
}
