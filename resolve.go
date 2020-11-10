package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

func getDNSServerIP(server string) string {
	serverIPs, err := net.LookupHost(server)
	var serverIP string
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(serverIPs) > 0 {
			serverIP = serverIPs[0]
		}
	}
	return serverIP
}

func resolve(hostname, server string) []string {
	server = fmt.Sprintf("%s:53", server)
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", server)
		},
	}
	ips, err := r.LookupHost(context.Background(), hostname)
	if err != nil {
		if !strings.Contains(err.Error(), "no such host") {
			fmt.Println(err.Error())
		}
	}
	return ips
}
