package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)
	if err != nil || !want.MatchString(message) {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	message, err := Hello(name)
	if message != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
