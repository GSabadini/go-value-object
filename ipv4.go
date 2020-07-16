package vo

import (
	"errors"
	"net"
)

var (
	// ErrInvalidIPv4 return invalid IPv4
	ErrInvalidIPv4 = errors.New("invalid IPv4")
)

// IP represents an IP type
type IP string

// IPv4 structure
type IPv4 struct {
	value string
}

// NewIPv4 create new IPv4
func NewIPv4(value string) (IPv4, error) {
	var IP = IPv4{value: value}

	if !IP.validate() {
		return IPv4{}, ErrInvalidIPv4
	}

	return IP, nil
}

func (ip IPv4) validate() bool {
	return net.ParseIP(ip.value) != nil
}

// Value return value IPv4
func (ip IPv4) Value() IP {
	return IP(ip.value)
}
