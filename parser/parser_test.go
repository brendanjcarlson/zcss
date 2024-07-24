package parser

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brendanjcarlson/zcss/lexer"
)

func Test_Parser(t *testing.T) {
	contents, err := os.ReadFile("../testdata/input.css")
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()

	l := lexer.New(string(contents))
	p := New(l)

	stylesheet := p.ParseStyleSheet()

	fmt.Println(stylesheet.CSS(false))
	fmt.Println(stylesheet.CSS(true))
	fmt.Println(stylesheet.Literal())

	elapsed := time.Since(now)
	fmt.Println("parsed in", elapsed.Milliseconds(), "ms")
}
