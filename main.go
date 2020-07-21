package main

import (
	"fmt"
	"log"

	"github.com/GSabadini/go-value-object/vo"
)

type Info struct {
	email    vo.Email
	ipv4     vo.IPv4
	filepath vo.FilePath
}

func main() {
	e, err := vo.NewEmail("test@test.com")
	if err != nil {
		log.Println(err)
	}

	e2, err := vo.NewEmail("test@test.com")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("\nemail equals:", e.Equals(e2))

	ip, err := vo.NewIPv4("192.168.0.110")
	if err != nil {
		log.Println(err)
	}

	ip2, err := vo.NewIPv4("192.168.0.110")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("\nip equals:", ip.Equals(ip2))

	f, err := vo.NewFilePath("/vo/file_path_test/")
	if err != nil {
		log.Println(err)
	}

	f1, err := vo.NewFilePath("/vo/file_path_test/")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("\nfilepath equals:", f.Equals(f1))

	p := Info{
		email:    e,
		ipv4:     ip,
		filepath: f,
	}

	fmt.Println("\nvalues:")
	fmt.Println(p.email.Value())
	fmt.Println(p.ipv4.Value())
	fmt.Println(p.filepath.Value())
}
