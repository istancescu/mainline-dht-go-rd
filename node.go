package dht

import (
	"github.com/istancescu/int160-go-rd"
	"net"
	"time"
)

type Node struct {
	ID        *int160.Int160
	IP        net.IP
	Port      uint16
	LastAlive time.Time
}
