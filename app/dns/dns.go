package dns

//go:generate go run $GOPATH/src/github.com/whatedcgveg/v2ray-core/tools/generrorgen/main.go -pkg dns -path App,DNS

import (
	"net"

	"github.com/whatedcgveg/v2ray-core/app"
)

// A Server is a DNS server for responding DNS queries.
type Server interface {
	Get(domain string) []net.IP
}

// FromSpace fetches a DNS server from context.
func FromSpace(space app.Space) Server {
	app := space.GetApplication((*Server)(nil))
	if app == nil {
		return nil
	}
	return app.(Server)
}
