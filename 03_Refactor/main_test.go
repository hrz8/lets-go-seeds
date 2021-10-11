package main

import (
	"testing"
)

func TestFindFirstStringInBracket(t *testing.T) {
	res := FindFirstStringInBracket("Nuruddin (Zayn Malik)")
	if res != "Zayn" {
		t.Fatalf(`FindFirstStringInBracket("Nuruddin (Zayn Malik)) = %s want "Zayn"`, res)
	}
}
