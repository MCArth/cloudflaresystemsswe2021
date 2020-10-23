package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
)

type profileResults struct {
	time     int
	success  bool
	numBytes int
	err      string
}

func main() {
	URL, numRepeats := getCommandLineArgs()

	timeList := make([]int, 0)
	successCount := 0
	failCount := 0
	errorList := make([]string, 0)
	smallestByteCount := math.MaxInt32
	largestByteCount := 0

	var resp string
	for repeatNo := 0; repeatNo < numRepeats; repeatNo++ {
		resp, result := httpGETURL(URL)
		timeList = append(timeList, result.time)
		if result.success {
			successCount++
		} else {
			failCount++
		}
		if result.err != "" {
			errorList = append(errorList, result.err)
		}
		smallestByteCount = minInt(smallestByteCount, result.numBytes)
		largestByteCount = maxInt(largestByteCount, result.numBytes)
	}
	percentSuccessful := (successCount / (successCount + failCount)) * 100

	fastestTime := math.MaxInt32
	slowestTime := 0
	timeSum := 0
	for _, time := range timeList {
		fastestTime = minInt(fastestTime, time)
		slowestTime = maxInt(slowestTime, time)
		timeSum += time
	}
	var medianTime float32
	if len(timeList)%2 == 1 {
		medianTime = float32(timeList[(len(timeList)-1)/2])
	} else {
		medianTime = float32(timeList[(len(timeList)/2)-1]+timeList[len(timeList)/2]) / 2
	}
	if numRepeats > 0 {
		fmt.Println("Requests completed:", numRepeats, "\n")
		fmt.Println("Fastest request roundtrip:", fastestTime, "\n")
		fmt.Println("Slowest request roundtrip:", slowestTime, "\n")
		fmt.Println("Mean request roundtrip time:", timeSum/numRepeats, "\n")
		fmt.Println("Median request roundtrip:", medianTime, "\n")
		fmt.Println("Percentage successful requests: ", percentSuccessful, "\n")
		fmt.Println("HTTP 4xx responses encountered:")
		for _, err := range errorList {
			fmt.Println(err)
		}
		fmt.Println()
		fmt.Println("Smallest request response:", smallestByteCount, "(bytes\n")
		fmt.Println("Largest request response:", largestByteCount, "(bytes)\n")
	} else {
		fmt.Println(resp)
	}
}

func httpGETURL(URL string) (response string, result profileResults) {
	// todo: check for non success (e.g. 400 response)
	address, path := requestParamsFromURL(URL)

	conn, err := net.Dial("tcp", address)
	// todo: defer conn.Close()
	checkForError(err)
	_, err = conn.Write([]byte("GET " + path + " HTTP/1.1\r\nHost:" + address + "\r\nConnection: close\r\n\r\n"))
	checkForError(err)
	var buff bytes.Buffer
	numBytes, err = io.Copy(&buff, conn)
	checkForError(err)

	return string(buff.Bytes())), profileResults(time, success, numBytes, HTTP4xxResponse)
}

func getCommandLineArgs() (URL string, numRepeats int) {
	urlPtr := flag.String(
		"URL",
		"cloudflaregeneralswe2021.mcarth.workers.dev:80/links",
		"URL string in the form host:port/path, defaults to cloudflaregeneralswe2021.mcarth.workers.dev:80/links")
	numRepeatsPtr := flag.Int(
		"numRepeats",
		0,
		"Number of times to repeat GET request, if 0 then profiler will not run and will simply output response, value defaults to 0")
	flag.Parse()

	numRepeats = *numRepeatsPtr
	URL = *urlPtr

	if numRepeats < 0 {
		log.Fatalln("Invalid argument passed to --profile. Correct usage is '--profile [x]' where x is a positive integer.")
	}

	return
}

func checkForError(e error) {
	if e != nil {
		log.Fatalln(e.Error()) // prints to stderr and exits
	}
}

func requestParamsFromURL(URL string) (address string, path string) {
	pathIndex := indexOfChar(URL, '/')
	if pathIndex == -1 {
		path = "/"
	} else {
		path = URL[pathIndex:]
	}
	address = URL[:pathIndex]
	return
}

func indexOfChar(str string, ch rune) int {
	for i, v := range str {
		if v == ch {
			return i
		}
	}
	return -1
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a int, b int) int {
	if b < a {
		return b
	}
	return a
}
