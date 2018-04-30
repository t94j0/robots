// Package robots parses robots.txt files from: http://www.robotstxt.org/orig.html.

package robots

import (
	"bufio"
	"io"
	"reflect"
	"strings"
)

// Robot is a struct that holds an array of strings that has the different type
// of items that can be in a /robots.txt file
type Robot struct {
	// Standard Directives
	Disallow   []string `target:"disallow" name:"Disallow"`
	UserAgents []string `target:"user-agent" name:"User-Agent"`

	// Non-standard Directives
	Extra []string
}

// Parse uses regular expressions on each line to get the useful data
// from the robots.txt file given as a string
func Parse(inputFile io.Reader) (robot Robot, err error) {
	inputBuf := bufio.NewScanner(inputFile)

	robotValue := reflect.ValueOf(&robot).Elem()
	robotT := robotValue.Type()

	for inputBuf.Scan() {
		line := inputBuf.Text()
		foundTag := false
		// Find which field has the correct target
		for i := 0; i < robotT.NumField(); i++ {
			field := robotValue.Field(i)
			target := robotT.Field(i).Tag.Get("target") + ": "
			if len(line) < len(target) {
				continue
			}
			actual := strings.ToLower(line[0:len(target)])
			if actual == target {
				value := strings.TrimSpace(line[len(target):])
				newSlice := reflect.Append(field, reflect.ValueOf(value))
				field.Set(newSlice)
				foundTag = true
			}
		}
		if !foundTag {
			robot.Extra = append(robot.Extra, line)
		}
	}
	return
}
