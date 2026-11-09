package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	reg = `NAME     "syn2020"
PATH     /var/opt/kontext/indexed/syn2020
LANGUAGE "Czech"
LOCALE   "cs_CZ.UTF-8"
ENCODING "utf-8"
INFO     "Synchronní reprezentativní korpus"
VERTICAL "/home/tomas/work/data/corpora/vertikaly/syn2020/syn2020_chunk3m"  # this is a comment

ATTRIBUTE lc {
	LABEL "lc [lowercase word]"
	DYNAMIC utf8lowercase
	DYNLIB internal  # and another comment
	ARG1 "C"
	FUNTYPE s
	FROMATTR word
	TYPE index
	TRANSQUERY yes
} # and comment in the end of attr

# and top level comment

ATTRIBUTE wordcount         # some comment

ATTRIBUTE sforma {  # comment on the first block line
	TYPE "FD_FGD"
	MULTIVALUE y
	MULTISEP "|"
}

ATTRIBUTE p_pos {
	LABEL "p_pos [parent part of speech]"
	DYNAMIC getnchar
	DYNLIB  internal
	ARG1    1
	FUNTYPE i
	FROMATTR p_tag
	TYPE index
}

STRUCTURE text {   # comment on the first block line
	TYPE "file64"
	ATTRIBUTE author {
			LOCALE "sk_SK"
	}
	ATTRIBUTE section
	ATTRIBUTE section_orig
	ATTRIBUTE id {
			TYPE "UNIQUE"
			LOCALE "en_US"
	}
}
`

	reg2 = `MAINTAINER "CLARIN.SI <info@clarin.si>"
NAME "CLASSLAWiki-sl (Slovenian Wikipedia)"
INFO "Slovenian Wikipedia corpus CLASSLAWiki-sl 1.0"
INFOHREF "http://hdl.handle.net/11356/1427"
ENCODING "UTF-8"
LOCALE "sl_SI.UTF-8"
DEFAULTLOCALE "sl_SI.UTF-8"
LANGUAGE "Slovenian"
TAGSETDOC "https://universaldependencies.org/guidelines.html"

PATH "/path/to/classlawiki_sl/"
VERTICAL "| zcat /path/to/vert/classlawiki-sl.vert.gz"

DOCSTRUCTURE text
MAXDETAIL 10000

SUBCORPATTRS "text.id|text.title|text.url"
SHORTREF "=text.title"
FULLREF "text.id,text.title,text.url,p.id,s.id"

STRUCTURE text {
    ATTRIBUTE id {
      LABEL "Text ID"
      TYPE "UNIQUE"
      TYPE "FD_FGD"
    }
    ATTRIBUTE title
    ATTRIBUTE url
}
STRUCTURE p {
    ATTRIBUTE id {
      TYPE "UNIQUE"
      TYPE "FD_FGD"
    }
    DISPLAYTAG 0
    DISPLAYBEGIN ""
    DISPLAYEND " ¶ "
}
ATTRIBUTE tag {
        LABEL   "MULTEXT-East MSD"
        TYPE "FD_FGD"
}`
)

func TestRegExample(t *testing.T) {
	doc, err := ParseRegistry("test1", reg)
	assert.NoError(t, err)

	sText := doc.GetStructure("text")
	assert.NotNil(t, sText)

	lcAttr := doc.GetPosAttr("lc")
	assert.NotNil(t, lcAttr)

	assert.Equal(t, "/var/opt/kontext/indexed/syn2020", string(doc.Entries.Get("PATH")))

	assert.Equal(t, "Synchronní reprezentativní korpus", string(doc.Entries.Get("INFO")))

}

func TestRegExampleClarinSL(t *testing.T) {
	doc, err := ParseRegistry("test1", reg2)
	assert.NoError(t, err)

	sText := doc.GetStructure("text")
	assert.NotNil(t, sText)

	lcAttr := doc.GetPosAttr("tag")
	assert.NotNil(t, lcAttr)

	assert.Equal(t, "/path/to/classlawiki_sl/", string(doc.Entries.Get("PATH")))

	assert.Equal(t, "Slovenian Wikipedia corpus CLASSLAWiki-sl 1.0", string(doc.Entries.Get("INFO")))
}

func TestParseIntercorpV13udCs(t *testing.T) {
	data, err := os.ReadFile("testdata/intercorp_v13ud_cs")
	assert.NoError(t, err)

	_, err = ParseRegistryBytes("intercorp_v13ud_cs", data)
	assert.NoError(t, err)
}
