package Common

import (
	"regexp"
	"strings"
)

func Loadini(filename string) map[string]map[string]string {
	configlist := make(map[string]map[string]string)
	sublist := []string{}
	var result [][]string
	var currentcfg string
	outCh := make(chan string, 200)
	go Read_file(filename, outCh)

	for x := range outCh {
		sublist = append(sublist, x)
	}
	Valuereg := regexp.MustCompile(`(.*)=(.*)`)
	Cfgreg := regexp.MustCompile(`\[(.*)\]`)

	for _, iniline := range sublist {
		if Valuereg.MatchString(iniline) {
			result = Valuereg.FindAllStringSubmatch(iniline, -1)

			configlist[currentcfg][strings.TrimSpace(result[0][1])] = strings.TrimSpace(result[0][2])

		} else {
			result = Cfgreg.FindAllStringSubmatch(iniline, -1)
			currentcfg = result[0][1]
			configlist[currentcfg] = map[string]string{}
		}
	}

	return configlist
}

func Configini(filename string, cfg string, key string, values string) {
	x := Loadini(filename)
	x[cfg][key] = values
}
