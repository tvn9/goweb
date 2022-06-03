package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n it was always cold no sunshine\n Yeah running' down a dream\n That never would come to me\n Workin' on a mestery\n goin' wherever it leads\n Running' down a dream"
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
