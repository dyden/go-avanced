package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

//GET OPPEN PORTS
var site = flag.String("site", "scanme.nmap.org", "Url to scan ports")

func main() {
	flag.Parse()
	fmt.Println("Scanning ports for: ", *site)
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Port", port, "is open")

		}(i)
	}
	wg.Wait()
	fmt.Println("Scanner finished")

}
