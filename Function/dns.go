package Function

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/UD94/SecondOP/Common"
)

func Google_domain(domain_name string) {

}

func Dns_thread(domain_name string) {

	var concontrolset = []string{}
	CacheFileName := Common.RandStringRunes(5)

	set_domain := "ud94iscreater." + domain_name
	ns, err := net.LookupHost(set_domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())

	} else {
		concontrolset = ns

	}

	sublist := []string{}
	outCh := make(chan string, 200)
	go Common.Read_file("domain.txt", outCh)

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

			for _, n := range ns {

				if !Common.In(n, concontrolset) {
					mutex.Lock()
					Common.Write_result(domain+",", "Cache\\"+CacheFileName)
					fmt.Fprintf(os.Stdout, "--%s\n", n)
					Common.Write_result(n+",", "Cache\\"+CacheFileName)
					Common.Write_result("\n", "Cache\\"+CacheFileName)
					mutex.Unlock()
				}

			}

			defer wait.Done()
		}()
	}
	wait.Wait()

}
