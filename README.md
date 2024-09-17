# czechdns

## Usage
```
# Create a file with a list of hostnames
oc get routes --all-namespaces --no-headers | awk '{print $3}' > tests/example_routes

# Create test file from a list of hostnames
./czechdns -createTest=true -hostsFile ./tests/example_routes -testFile tests/example_tests

# Run test against a nameserver you wish to czech
./czechdns -testFile tests/example_tests -server ns-1684.awsdns-18.co.uk
```

## Install

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/czechdns/releases)

## Build
```/bin/bash
go build
```
