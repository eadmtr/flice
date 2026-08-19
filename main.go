package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var testFile = "./test-file.txt"

func main() {
	fmt.Println(testFile)
	// freaderV1()

	testGetURL()
}

func freaderV1() {
	src := strings.NewReader("1234")
	dst := make([]byte, 3)

	cnt, err := src.Read(dst)
	println(cnt, err)
	println(string(dst[:cnt]))
	println("---")

	cnt1, err1 := src.Read(dst)
	println(cnt1, err1)
	println(string(dst[:cnt1]))
	println("---")
	// Printfln("Read %v bytes: %v", count, string(b[0:count]))
}

func testGetURL() {
	dst := getRequestDestination()
	for _, url := range dst {
		status := getReqStatus(url)
		println(url, "\n\t", status, "\n")
	}
}

func getRequestDestination() []string {
	r := []string{}
	rU := getURL()
	rP := getPATH()

	for _, u := range rU {
		for _, p := range rP {
			r = append(r, "https://"+u+"/"+p)
		}
	}

	return r
}

func getURL() []string {
	return []string{"google.com", "ya.ru", "bing.com"}
}

func getPATH() []string {
	return []string{"index.html", "robot.txt", "favicon.ico"}
}

func getReqStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Status
}
