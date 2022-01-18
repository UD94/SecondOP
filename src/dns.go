package src

import (
	"end"
	"fmt"
	"net"
	"os"
	"sync"
)

func Google_domain(domain_name string) {

}

func Dns_thread(domain_name string) {
	var file_use end.F_CONTROL

	var concontrolset = []string{}

	set_domain := "ud94iscreater." + domain_name
	ns, err := net.LookupHost(set_domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())

	} else {
		concontrolset = ns
	}

	sublist := []string{}
	outCh := make(chan string, 200)
	go file_use.Read_file("domain.txt", outCh)

	for x := range outCh {
		sublist = append(sublist, x)
	}

	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	for _, s := range sublist {

		wait.Add(1)
		domain := s + "." + domain_name
		go func() {

			ns, err := net.LookupHost(domain)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
				return
			}
			if file_use.Equal(concontrolset, ns) {
				return
			} else {
				mutex.Lock()
				file_use.write_result(domain+",", "log.txt")
				for _, n := range ns {
					fmt.Fprintf(os.Stdout, "--%s\n", n)
					file_use.write_result(n+",", "log.txt")
				}
				file_use.write_result("\n", "log.txt")
				mutex.Unlock()
			}

			defer wait.Done()
		}()
	}
	wait.Wait()

}
