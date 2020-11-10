package main

import (
	"flag"
)

func main() {
	serverPtr := flag.String("server", "8.8.8.8", "UDP DNS server")
	inputFilePtr := flag.String("hostsFile", "", "path to txt file, containing one hostname per line. Required for createTest")
	testFilePtr := flag.String("testFile", "", "path to yaml test file, if createTest=true, will overwrite this file")
	createTestPtr := flag.Bool("createTest", false, "if true, will create a test file")

	flag.Parse()
	if *createTestPtr {
		writeTest(*serverPtr, *inputFilePtr, *testFilePtr)
	} else {
		runTests(*serverPtr, *testFilePtr)
	}
}
