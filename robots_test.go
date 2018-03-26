package robots_test

import (
	"strings"
	"testing"

	"github.com/t94j0/robots"
)

func TestParse_excludeall(t *testing.T) {
	test := `User-agent: *
Disallow: /test
`

	testReader := strings.NewReader(test)
	output, err := robots.Parse(testReader)
	if err != nil {
		t.Fatal(err)
	}

	ua := output.UserAgents
	d := output.Disallow

	if len(ua) != 1 || ua[0] != "*" {
		t.Fatal("User agent '*' not found")
	}

	if len(d) != 1 || d[0] != "/test" {
		t.Fatal("Disallow '/test' not found")
	}
}

func TestParse_extra(t *testing.T) {
	target := "// This might be a comment or something"
	test := `User-Agent: *
Disallow: /`
	test = target + "\n" + test

	testReader := strings.NewReader(test)
	output, err := robots.Parse(testReader)
	if err != nil {
		t.Fatal(err)
	}

	ua := output.UserAgents
	d := output.Disallow
	e := output.Extra

	if len(ua) != 1 || ua[0] != "*" {
		t.Fatal("User agent parsing broken")
	}

	if len(d) != 1 || d[0] != "/" {
		t.Fatal("Disallow parsing broken")
	}

	if len(e) != 1 || e[0] != target {
		t.Fatal("Extra parsing broken")
	}
}
