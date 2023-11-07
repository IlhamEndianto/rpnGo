package main

import "testing"

var (
	answer1 = "abc*+"
	answer2 = "ab+zx+*"
	answer3 = "at+bac++cd+^*"
)

func TestCase1(t *testing.T) {
	text := "(a+(b*c))"
	t.Logf("Case:%s", text)

	result := ComposeRPNFromText(text)
	if result != answer1 {
		t.Errorf("Error! Result should be: %s", answer1)
	}
}

func TestCase2(t *testing.T) {
	text := "((a+b)*(z+x))"
	t.Logf("Case:%s", text)

	result := ComposeRPNFromText(text)
	if result != answer2 {
		t.Errorf("Error! Result should be: %s", answer2)
	}
}
func TestCase3(t *testing.T) {
	text := "((a+t)*((b+(a+c))^(c+d)))"
	t.Logf("Case:%s", text)

	result := ComposeRPNFromText(text)
	if result != answer3 {
		t.Errorf("Error! Result should be: %s", answer3)
	}
}
