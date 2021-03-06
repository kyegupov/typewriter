package template

import (
	"regexp"
	"strings"
	"text/template"
)

// Language is a parsable language
// stringer -type=Language
type Language int

// languages
const (
	Typescript Language = iota
	Flow
	Elm
)

// custom types
const (
	EmptyInterface = "emptyIface"
	NestedStruct   = "struct"
	TimeStruct     = "Date"
)

var funcMap = template.FuncMap{
	"updateFlowType":       updateTypes(conversions[Flow]),
	"updateElmType":        updateTypes(conversions[Elm]),
	"updateTSType":         updateTypes(conversions[Typescript]),
	"flowComment":          lineComment("//"),
	"elmComment":           lineComment("--"),
	"tsComment":            lineComment("//"),
	"flowMultilineComment": multilineComment("//"),
	"elmMultilineComment":  multilineComment("--"),
	"tsMultilineComment":   multilineComment("//"),
}

const goInt = "int64|int32|int16|int8|int|uint64|uint32|uint16|uint8|uint|byte|rune"
const goFloat = "float32|float64|complex64|complex128"
const goNumbers = goInt + "|" + goFloat

func asWord(baseRegex string) *regexp.Regexp {
	return regexp.MustCompile("\\b(" + baseRegex + ")\\b")
}

var conversions = map[Language]map[string]*regexp.Regexp{
	Flow: map[string]*regexp.Regexp{
		"any":     asWord(EmptyInterface),
		"Object":  asWord(NestedStruct),
		"Date":    asWord(TimeStruct),
		"number":  asWord(goNumbers),
		"boolean": asWord("bool"),
	},
	Typescript: map[string]*regexp.Regexp{
		"any":     asWord(EmptyInterface),
		"object":  asWord(NestedStruct),
		"Date":    asWord(TimeStruct),
		"number":  asWord(goNumbers),
		"boolean": asWord("bool"),
	},
	Elm: map[string]*regexp.Regexp{
		"string": asWord("string"),
		"Maybe":  asWord(EmptyInterface + "|" + NestedStruct),
		"Date":   asWord(TimeStruct),
		"Bool":   asWord("bool"),
		"Int":    asWord(goInt),
		"Float":  asWord(goFloat),
	},
}

// updateTypes takes a conversion slice and returns
// a function used as a string replacer
func updateTypes(replacements map[string]*regexp.Regexp) func(string) string {
	return func(s string) string {
		for target, regex := range replacements {
			s = regex.ReplaceAllString(s, target)
		}
		return s
	}
}

func multilineComment(prefix string) func(string, int) string {
	return func(c string, indent int) string {
		if c == "" {
			return c
		}
		lineStart := strings.Repeat("\t", indent) + prefix + " "
		return lineStart + strings.ReplaceAll(strings.Trim(c, "\n"), "\n", "\n"+lineStart) + "\n"
	}
}

func lineComment(prefix string) func(string) string {
	return func(c string) string {
		if c == "" {
			return c
		}
		return " " + prefix + " " + c
	}
}
