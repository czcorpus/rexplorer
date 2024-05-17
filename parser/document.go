//go:generate pigeon -o ./parser.go ./grammar.peg
package parser

func fromIdxOfUntypedSlice(arr any, idx int) any {
	if arr == nil {
		return nil
	}
	v := arr.([]any)
	return v[idx]
}

type KeyVals []*KeyVal

func (kv KeyVals) Get(name string) string {
	for _, v := range kv {
		if v.Name == name {
			return v.Value
		}
	}
	return ""
}

type Document struct {
	Entries    KeyVals
	PosAttrs   []*Attr
	Structures []*Structure
}

func (doc *Document) GetProperty(name string) string {
	return doc.Entries.Get(name)
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
	Name  string
	Value string
}

type Attr struct {
	Name    string
	Entries KeyVals
}

func (attr *Attr) GetProperty(name string) string {
	return attr.Entries.Get(name)
}

type Structure struct {
	Name    string
	Attrs   []*Attr
	Entries KeyVals
}

func (st *Structure) GetAttribute(name string) *Attr {
	for _, a := range st.Attrs {
		if a.Name == name {
			return a
		}
	}
	return nil
}

func NewDocument() *Document {
	return &Document{
		Entries:    make(KeyVals, 0, 20),
		PosAttrs:   make([]*Attr, 0, 20),
		Structures: make([]*Structure, 0, 20),
	}
}
