package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter url: ")

	url, _ := reader.ReadString('\n')
	if len(url) == 0 {
		log.Fatal("No url provided")
	}
	url = strings.TrimSuffix(url, "\n")
	url = strings.TrimSuffix(url, "\r")

	prepend := "https://"
	var buffer bytes.Buffer

	if len(url) <= len(prepend) {
		buffer.WriteString(prepend)
		buffer.WriteString(url)
	} else if url[:len(prepend)] != prepend {
		buffer.WriteString(prepend)
		buffer.WriteString(url)
	} else {
		buffer.WriteString(url)
	}

	fmt.Printf("HTML code of %s ...\n", buffer.String())
	resp, err := http.Get(buffer.String())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	regex, err := regexp.Compile("\"(https://.*?)\"|\"(http://.*?)\"")
	if err != nil {
		panic(err)
	}
	regResult := regex.FindAllString(string(html), -1)
	// iterate over the regex results
	for _, result := range regResult {
		fmt.Printf("%s \n", result)
	}
}
