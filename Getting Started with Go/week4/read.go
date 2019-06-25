package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Record struct {
	fname, lname string
}

func main() {
	var filename string

	fmt.Print("Enter a filename: ")
	fmt.Scan(&filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	records := RecordsFromBytes(content)

	for i, rec := range records {
		fmt.Printf("%v. %s %s\n", i+1, rec.fname, rec.lname)
	}
}

// ParseStringToRecord creates the Record structure from the string s.
// The s string should contains two words (first name and last name)
// separated by a space.
func ParseStringToRecord(s string) Record {
	data := strings.Split(s, " ")
	fname, lname := data[0], data[1]
	return Record{fname, lname}
}

// RecordsFromBytes creates slice of the Record structures
// from a byte array.
func RecordsFromBytes(content []byte) []Record {
	var records []Record

	rows := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, row := range rows {
		records = append(records, ParseStringToRecord(row))
	}
	return records
}
