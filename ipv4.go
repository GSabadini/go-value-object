package go_value_objects

import (
	"errors"
	"net"
)

var (
	ErrInvalidIPv4 = errors.New("invalid IPv4")
)

type IP string

type IPv4 struct {
	value string
}

func (ip IPv4) validate() bool {
	if net.ParseIP(ip.value) == nil {
		return false
	}

	return true
}

func NewIPv4(value string) (IPv4, error) {
	var IP = IPv4{value: value}

	if !IP.validate() {
		return IPv4{}, ErrInvalidIPv4
	}

	return IP, nil
}

func (ip IPv4) Value() IP {
	return IP(ip.value)
}