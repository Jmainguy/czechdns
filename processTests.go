package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func writeTest(server, inputFile, outputFile string) {
	var ers expectedResults
	hostnames := readFile(inputFile)
	for _, hostname := range hostnames {
		var er expectedResult
		server = getDNSServerIP(server)
		ips := resolve(hostname, server)
		if len(ips) > 0 {
			er.Hostname = hostname
			er.IPS = ips
			ers = append(ers, er)
		}
	}
	data, err := yaml.Marshal(&ers)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func runTests(server, testFile string) {
	var hostnames []string
	var ers expectedResults
	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &ers)
	if err != nil {
		fmt.Println(err)
	}
	var msgs []string
	for _, er := range ers {
		hostname := er.Hostname
		hostnames = append(hostnames, hostname)
		expectedIPS := er.IPS
		server = getDNSServerIP(server)
		ips := resolve(hostname, server)
		sort.Strings(expectedIPS)
		sort.Strings(ips)
		match := reflect.DeepEqual(expectedIPS, ips)
		if !match {
			msg := "Comparison of " + hostname + " Failed: expected [" + strings.Join(expectedIPS, " ") + "] got [" + strings.Join(ips, " ") + "]"
			msgs = append(msgs, msg)
		}
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			fmt.Println(msg)
		}
	} else {
		msg := fmt.Sprintf("Tests against all %v hostnames have passed", len(hostnames))
		fmt.Println(msg)
	}
}
