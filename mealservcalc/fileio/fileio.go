package fileio

import (
	"bufio"
	"log"
	"os"
)

func ReadDataFromFile(name string) (ss []string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewScanner(file)
	if err := br.Err(); err != nil {
		log.Fatal(err)
	}
	for br.Scan() {
		s := br.Text()
		ss = append(ss, s)
	}
	return ss
}
