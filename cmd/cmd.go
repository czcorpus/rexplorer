package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/czcorpus/rexplorer/parser"
)

const (
	outTypeJSON outType = "json"
	outTypeTXT  outType = "txt"
	outTypeNone outType = "none"
)

type outType string

func (ot outType) validate() error {
	if ot != outTypeJSON && ot != outTypeNone && ot != outTypeTXT {
		return fmt.Errorf("invalid output type: %s", ot)
	}
	return nil
}

var (
	q1 = regexp.MustCompile(`^a\[([^]]+)\](\[([^]]+)\])?$`)
	q2 = regexp.MustCompile(`^s\[([^]]+)\]a\[([^]]+)\](\[([^]]+)\])?$`)
	q3 = regexp.MustCompile(`^s\[([^]]+)\](\[([^]]+)\])?$`)
)

func outObject(v fmt.Stringer, outType outType) string {
	switch outType {
	case outTypeJSON:
		ans, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
		return string(ans)
	case outTypeTXT:
		return v.String()
	case outTypeNone:
		return ""
	}
	return ""
}

func queryDocument(q string, document *parser.Document) fmt.Stringer {
	ans := q1.FindStringSubmatch(q)
	if len(ans) > 0 {
		attr := document.GetPosAttr(ans[1])
		if attr == nil {
			return nil
		}
		if ans[3] != "" {
			v := attr.GetProperty(ans[3])
			return v
		}
		return attr
	}
	ans = q2.FindStringSubmatch(q)
	if len(ans) > 0 {
		str := document.GetStructure(ans[1])
		if str == nil {
			return nil
		}
		attr := str.GetAttribute(ans[2])
		if attr == nil {
			return nil
		}
		if ans[4] != "" {
			return attr.GetProperty(ans[4])
		}
		return attr
	}
	ans = q3.FindStringSubmatch(q)
	if len(ans) > 0 {
		str := document.GetStructure(ans[1])
		if str == nil {
			return nil
		}
		if ans[3] != "" {
			return str.GetProperty(ans[3])
		}
		return str
	}
	return nil
}

func main() {
	// rexplorer A[word][type]
	// rexplorer S[document][foo]
	// rexplorer S[document]A[id]
	var outTypeSel string
	flag.StringVar(&outTypeSel, "out-type", "txt", "Output type (json, txt, none)")
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
	tOutType := outType(outTypeSel)
	if err := tOutType.validate(); err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
	query := flag.Arg(1)
	if query != "" {
		ans := queryDocument(query, document)
		if ans != nil {
			fmt.Println(outObject(ans, tOutType))

		} else {
			fmt.Println("object not found")
		}

	} else {
		fmt.Println(outObject(document, tOutType))
	}
}
