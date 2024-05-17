package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/czcorpus/rexplorer/parser"
)

var (
	q1 = regexp.MustCompile(`a\[([^]]+)\](\[([^]]+)\])?`)
)

func query(q string, document *parser.Document) any {
	ans := q1.FindStringSubmatch(q)
	if len(ans) > 0 {
		attr := document.GetPosAttr(ans[1])
		//spew.Dump(document.Entries)
		v := attr.GetProperty(ans[3])
		return v
	}
	return nil
}

func main() {
	// rexplorer A[word][type]
	// rexplorer S[document][foo]
	// rexplorer S[document]A[id]

	flag.Parse()
	f, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	document, err := parser.ParseRegistryBytes(filepath.Base(flag.Arg(0)), f)
	if err != nil {
		errRn := []rune(err.Error())
		if len(errRn) > 140 {
			fmt.Println(string(errRn[:140]))

		} else {
			fmt.Println(string(errRn))
		}
		os.Exit(2)
	}
	ans := query(flag.Arg(1), document)
	fmt.Println(ans)
}
