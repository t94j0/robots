package robots

import (
	"regexp"
	"strings"
)

// Robot is a struct that holds an array of strings that has the different type
// of items that can be in a /robots.txt file
type Robot struct {
	// Standard Directives
	UserAgents []string
	Disallow   []string

	// Non-standard Directives
}

// ParseRobots uses regular expressions on each line to get the useful data
// from the robots.txt file given as a string
func ParseRobots(robotFile string) (robot Robot) {
	disallowRegex := regexp.MustCompile("^(Disallow:) (.+)")
	userAgentRegex := regexp.MustCompile("^(User-Agent:) (.+)")

	for _, line := range strings.Split(robotFile, "\n") {
		disallowSubmatches := disallowRegex.FindStringSubmatch(line)
		if len(disallowSubmatches) == 3 {
			robot.Disallow = append(robot.Disallow, disallowSubmatches[2])
		}

		userAgentSubmatches := userAgentRegex.FindStringSubmatch(line)
		if len(userAgentSubmatches) == 3 {
			robot.UserAgents = append(robot.UserAgents, userAgentSubmatches[2])
		}
	}

	return
}
