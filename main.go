package main

import (
	"fmt"
	"log"

	"github.com/GSabadini/go-value-object/vo"
)

type Info struct {
	Email vo.Email
	IPv4  vo.IPv4
}

func main() {
	e, err := vo.NewEmail("test@test.com")
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(e.String())
	//
	//e2, err := vo.NewEmail("test@test.com")
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(e.String())
	//
	//fmt.Println("\n", e.Equals(e2))

	ip, err := vo.NewIPv4("192.168.0.110")
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(e.String())
	//
	//ip2, err := vo.NewIPv4("192.168.0.110")
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(e.String())
	//
	//fmt.Println("\n", ip.Equals(ip2))

	p := Info{
		Email: e,
		IPv4:  ip,
	}

	fmt.Println(p.Email.String())
}
