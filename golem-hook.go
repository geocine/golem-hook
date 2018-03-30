package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const snipLine = "------------------------ >8 ------------------------"
const commentChar = '#'

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You must provide a path to a file containing"+
			" the commit message.")
		os.Exit(1)
	}

	path := os.Args[1]
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file \"%s\".\n", path)
		os.Exit(1)
	}
	msg := string(bytes)

	cleaned := cleanMsg(msg)
	subject, body := parseMsg(cleaned)
	fmt.Printf("subject: %s \n", subject)
	fmt.Printf("body: %s \n", body)
}

func cleanMsg(msg string) string {
	remComments := bytes.Buffer{}
	split := strings.SplitAfter(msg, "\n")
	for _, line := range split {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, string(commentChar)+" "+snipLine) {
			break
		}
		if strings.HasPrefix(trim, string(commentChar)) {
			continue
		}

		remComments.WriteString(line)
	}
	return strings.TrimSpace(remComments.String())
}

func parseMsg(cleanMsg string) (subject string, body string) {
	split := strings.SplitN(strings.TrimSpace(cleanMsg), "\n\n", 2)
	subject = split[0]
	if len(split) > 1 {
		body = split[1]
	}
	return
}
