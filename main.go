package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"sort"
	"time"
)

type profileResults struct {
	time     int
	success  bool
	numBytes int
	err      string
}

func main() {
	URL, numRepeats := getCommandLineArgs()
	profilerEnabled := true
	if numRepeats == 0 {
		profilerEnabled = false
		numRepeats = 1
	}

	timeList := make([]int, 0)
	successCount := 0
	errorList := make([]string, 0)

	smallestByteCount := math.MaxInt32
	largestByteCount := 0

	var resp string
	var result profileResults
	for repeatNo := 0; repeatNo < numRepeats; repeatNo++ {
		resp, result = httpGETURL(URL)
		timeList = append(timeList, result.time)
		if result.success {
			successCount++
		}
		if result.err != "" {
			errorList = append(errorList, result.err)
		}
		smallestByteCount = minInt(smallestByteCount, result.numBytes)
		largestByteCount = maxInt(largestByteCount, result.numBytes)
	}
	percentSuccessful := (successCount / numRepeats) * 100

	fastestTime := math.MaxInt32
	slowestTime := 0
	timeSum := 0
	for _, time := range timeList {
		fastestTime = minInt(fastestTime, time)
		slowestTime = maxInt(slowestTime, time)
		timeSum += time
	}
	sort.Ints(timeList)
	var medianTime float32
	if len(timeList)%2 == 1 {
		medianTime = float32(timeList[(len(timeList)-1)/2])
	} else {
		medianTime = float32(timeList[(len(timeList)/2)-1]+timeList[len(timeList)/2]) / 2
	}
	if profilerEnabled {
		fmt.Println("Requests attempted:", numRepeats, "\n")
		if fastestTime > 0 { // a request was completed
			fmt.Println("Fastest request roundtrip:", fastestTime, "ms")
			fmt.Println("Slowest request roundtrip:", slowestTime, "ms")
			fmt.Println("Mean request roundtrip time:", timeSum/len(timeList), "ms")
			fmt.Println("Median request roundtrip:", medianTime, "ms", "\n")
		} else {
			fmt.Println("No requests were completed\n")
		}
		fmt.Println("Percentage successful requests: ", percentSuccessful, "\n")
		if len(errorList) != 0 {
			fmt.Println("Error or non-2xx code responses encountered:")
			for _, err := range errorList {
				fmt.Println(err)
			}
			fmt.Println()
		} else {
			fmt.Println("No errors or non-2xx responses encountered\n")
		}

		if fastestTime > 0 {
			fmt.Println("Smallest request response:", smallestByteCount, "(bytes)")
			fmt.Println("Largest request response:", largestByteCount, "(bytes)")
		}

	} else if !profilerEnabled {
		fmt.Println(resp)
	}
}

func httpGETURL(URL string) (response string, result profileResults) {
	results := profileResults{}
	address, path := requestParamsFromURL(URL)

	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	if err != nil {
		results.err = "Error encountered while establishing connection: " + err.Error()
	}

	start := time.Now()
	_, err = conn.Write([]byte("GET " + path + " HTTP/1.1\r\nHost:" + address + "\r\nConnection: close\r\n\r\n"))
	if err != nil {
		results.err = "Error encountered while writing request to remote: " + err.Error()
		return
	}

	var buff bytes.Buffer
	numBytes, err := io.Copy(&buff, conn)
	if err != nil {
		results.err = "Error encountered while reading response from remote: " + err.Error()
		return
	}

	results.numBytes = int(numBytes)
	results.time = int(time.Since(start).Milliseconds())

	response = string(buff.Bytes())
	if response[9] != '2' {
		results.err = "Non 2xx http response code: " + response[9:indexOfChar(response, '\r')]
	} else {
		results.success = true
	}

	return response, results
}

func getCommandLineArgs() (URL string, numRepeats int) {
	urlPtr := flag.String(
		"url",
		"cloudflaregeneralswe2021.mcarth.workers.dev:80/links",
		"URL string in the form host:port/path")
	numRepeatsPtr := flag.Int(
		"profile",
		0,
		"Number of times to repeat GET request while profiling remote responses, if 0 then a single request is sent and the response is printed, value defaults to 0")
	flag.Parse()

	numRepeats = *numRepeatsPtr
	URL = *urlPtr

	if numRepeats < 0 {
		log.Fatalln("Invalid argument passed to --profile. Correct usage is '--profile [x]' where x is a positive integer.")
	}

	return
}

func requestParamsFromURL(URL string) (address string, path string) {
	var autoPort string
	if URL[:7] == "http://" {
		URL = URL[7:]
		autoPort = ":80"
	} else if URL[:8] == "https://" {
		URL = URL[8:]
		autoPort = ":80"
	}

	pathIndex := indexOfChar(URL, '/')
	if pathIndex == -1 {
		address, path = URL, "/"
	} else {
		address, path = URL[:pathIndex], URL[pathIndex:]
	}
	address += autoPort
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
