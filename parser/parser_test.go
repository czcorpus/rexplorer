package parser

import (
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
VERTICAL "/home/tomas/work/data/corpora/vertikaly/syn2020/syn2020_chunk3m"

ATTRIBUTE lc {
	LABEL "lc [lowercase word]"
	DYNAMIC utf8lowercase
	DYNLIB internal
	ARG1 "C"
	FUNTYPE s
	FROMATTR word
	TYPE index
	TRANSQUERY yes
}

ATTRIBUTE sforma {
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

STRUCTURE text {
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
)

func TestRegExample(t *testing.T) {
	_, err := ParseRegistry("test1", reg)
	assert.NoError(t, err)
}
