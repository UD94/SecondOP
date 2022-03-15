package Function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/UD94/SecondOP/Common"
	"github.com/haccer/subjack/subjack"
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

func Subhackdomain() {
	sublist := []string{}
	outCh := make(chan string, 200)
	go Common.Read_file("domain.txt", outCh)

	for x := range outCh {
		sublist = append(sublist, x)
	}

	target := []string{}
	outCh1 := make(chan string, 200)
	go Common.Read_file("sublist.txt", outCh1)

	for x := range outCh1 {
		target = append(target, x)
	}

	for _, subdomain := range sublist {
		for _, domain := range target {
			go subhack(subdomain + domain)
		}
	}

}

func subhack(subdomain string) {
	var fingerprints []subjack.Fingerprints
	config, _ := ioutil.ReadFile("custom_fingerprints.json")
	json.Unmarshal(config, &fingerprints)

	/* Use subjack's advanced detection to identify
	if the subdomain is able to be taken over. */
	service := subjack.Identify(subdomain, false, false, 10, fingerprints)

	if service != "" {
		service = strings.ToLower(service)
		fmt.Printf("%s is pointing to a vulnerable %s service.\n", subdomain, service)
	}
}
