package lexer

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brendanjcarlson/zcss/token"
)

func Test_Lexer(t *testing.T) {
	contents, err := os.ReadFile("../testdata/input.css")
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()

	l := New(string(contents))
	for {
		tok := l.NextToken()
		// fmt.Println(tok)
		// fmt.Println()
		if tok.Kind() == token.END_OF_FILE {
			break
		}
	}

	elapsed := time.Since(now)
	fmt.Println("lexed in", elapsed.Milliseconds(), "ms")
}
