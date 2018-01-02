package server

import (
	"os"
)

// NewServerConnection create connection for server api
func NewServerConnection(accessKey string, secretKey string) *ServerConn {
	ServerConn := &ServerConn{
		accessKey: accessKey,
		secretKey: secretKey,
		apiURL:    "https://api.ncloud.com/server/",
	}

	// for other phase(dev, test, beta ...) test
	if os.Getenv("NCLOUD_API_GW") != "" {
		ServerConn.apiURL = os.Getenv("NCLOUD_API_GW")
	}

	return ServerConn
}
