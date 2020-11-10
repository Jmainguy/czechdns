package main

type expectedResult struct {
	Hostname string   `yaml:"hostname"`
	IPS      []string `yaml:"ips"`
}

type expectedResults []expectedResult
