//go:generate pigeon -o ./parser.go ./grammar.peg
package parser

import (
	"fmt"
	"strings"
)

func mknumtabs(n int) string {
	var ans strings.Builder
	for i := 0; i < n; i++ {
		ans.WriteString("\t")
	}
	return ans.String()
}

func fromIdxOfUntypedSlice(arr any, idx int) any {
	if arr == nil {
		return nil
	}
	v := arr.([]any)
	return v[idx]
}

type property string

// Value returns a representation
// without additional characters which
// is useful for further processing.
func (p property) Value() string {
	return string(p)
}

func (p property) String() string {
	return p.IndentedString(0)
}

// IndentedString produces the value as a registry
// file entry with some indentation and a newline
// character at the end of the line.
func (p property) IndentedString(ntab int) string {
	return mknumtabs(ntab) + string(p) + "\n"
}

func (p property) IsEmpty() bool {
	return p == ""
}

type KeyVals []*KeyVal

func (kv KeyVals) Get(name string) property {
	for _, v := range kv {
		if v.Name == name {
			return property(v.Value)
		}
	}
	return property("")
}

type Document struct {
	Entries    KeyVals      `json:"entries"`
	PosAttrs   []*Attr      `json:"posAttrs"`
	Structures []*Structure `json:"stuctures"`
}

func (doc *Document) String() string {
	var ans strings.Builder

	for _, v := range doc.Entries {
		ans.WriteString(v.IndentedString(0))
	}
	ans.WriteString("\n")
	for _, v := range doc.PosAttrs {
		ans.WriteString(v.IndentedString(0))
	}
	ans.WriteString("\n")
	for _, v := range doc.Structures {
		ans.WriteString(v.IndentedString(0))
	}
	ans.WriteString("\n")
	return ans.String()
}

func (doc *Document) GetProperty(name string) property {
	return doc.Entries.Get(name)
}

// GetStaticPosattrs returns positional attributes
// which are not DYNAMIC. I.e. the attributes returned
// by this method (including their order) should be
// exactly matching data in a corresponding vertical file.
func (doc *Document) GetStaticPosattrs() []*Attr {
	ans := make([]*Attr, 0, len(doc.PosAttrs))
	for _, v := range doc.PosAttrs {
		dn := v.Entries.Get("DYNAMIC")
		if dn.IsEmpty() {
			ans = append(ans, v)
		}
	}
	return ans
}

func (doc *Document) GetPosAttr(name string) *Attr {
	for _, p := range doc.PosAttrs {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func (doc *Document) GetStructure(name string) *Structure {
	for _, s := range doc.Structures {
		if s.Name == name {
			return s
		}
	}
	return nil
}

type KeyVal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (kv KeyVal) String() string {
	return kv.IndentedString(0)
}

func (kv KeyVal) IndentedString(ntab int) string {
	return fmt.Sprintf("%s%s\t%s\n", mknumtabs(ntab), kv.Name, kv.Value)
}

type Attr struct {
	Name    string  `json:"name"`
	Entries KeyVals `json:"entries"`
}

func (attr *Attr) String() string {
	return attr.IndentedString(0)
}

func (attr *Attr) IndentedString(ntab int) string {
	if len(attr.Entries) == 0 {
		return fmt.Sprintf("%s%s\n", mknumtabs(ntab), attr.Name)
	}
	var ans strings.Builder
	ans.WriteString(fmt.Sprintf("%s%s {\n", mknumtabs(ntab), attr.Name))
	for _, v := range attr.Entries {
		ans.WriteString(mknumtabs(ntab) + v.IndentedString(ntab+1))
	}
	ans.WriteString(mknumtabs(ntab) + "}\n")
	return ans.String()
}

func (attr *Attr) GetProperty(name string) property {
	return attr.Entries.Get(name)
}

type Structure struct {
	Name    string  `json:"name"`
	Attrs   []*Attr `json:"attr"`
	Entries KeyVals `json:"entries"`
}

func (st *Structure) String() string {
	return st.IndentedString(0)
}

func (st *Structure) IndentedString(ntab int) string {
	var ans strings.Builder
	ans.WriteString(fmt.Sprintf("%s%s {\n", mknumtabs(ntab), st.Name))
	for _, v := range st.Entries {
		ans.WriteString(mknumtabs(ntab) + v.IndentedString(ntab+1))
	}
	if len(st.Attrs) > 0 {
		for _, v := range st.Attrs {
			ans.WriteString(mknumtabs(ntab) + v.IndentedString(ntab+1))
		}

	} else {
		ans.WriteString(mknumtabs(ntab) + "\n")
	}
	ans.WriteString(mknumtabs(ntab) + "}\n")
	return ans.String()
}

func (st *Structure) GetAttribute(name string) *Attr {
	for _, a := range st.Attrs {
		if a.Name == name {
			return a
		}
	}
	return nil
}

func (st *Structure) GetProperty(name string) property {
	return st.Entries.Get(name)
}

func NewDocument() *Document {
	return &Document{
		Entries:    make(KeyVals, 0, 20),
		PosAttrs:   make([]*Attr, 0, 20),
		Structures: make([]*Structure, 0, 20),
	}
}
