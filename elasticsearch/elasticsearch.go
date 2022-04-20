package elasticsearch

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"

	es "github.com/elastic/go-elasticsearch/v7"
)

type ClientConfig struct {
	Address  string
	Username string
	Password string
	CAPem    []byte
}

// 以逗号分隔的集群地址  http://node1:9200/,http://node2:9200/,http://node3:9200/
func (c ClientConfig) SplitAddress() []string {
	return strings.Split(c.Address, ",")
}

func NewClient(config ClientConfig) (client *es.Client, err error) {

	cfg := es.Config{
		Addresses: config.SplitAddress(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	if config.Username != "" && config.Password != "" {
		cfg.Username = config.Username
		cfg.Password = config.Password
	}
	if len(config.CAPem) != 0 {
		cfg.CACert = config.CAPem
	}

	return es.NewClient(cfg)
}
