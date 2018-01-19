package doc

import (
	"fmt"
	"regexp"
	"text/template"
)

type ParameterType int

const (
	Number ParameterType = iota
	String
	Boolean
	EnumString
	EnumNumber
)

const (
	numberRe = `^[0-9\.]+$`
	boolRe   = `^(?:[tT][rR][uU][eE]|[fF][aA][lL][sS][eE])$`
)

var (
	parameterTmpl *template.Template
	parameterFmt  = `    + {{.Name}}: {{.Value.Quote}} ({{.Type.String}}){{with .Description}} - {{.}}{{end}}
{{with .AdditionalDescription}}{{.}}{{end}}
{{with .DefaultValue}}+ Default: {{.}}{{end}}
{{if .Members}}
    + Members{{range .Members}}
		+ {{.}}{{end}}
{{end}}
`
)

func init() {
	parameterTmpl = template.Must(template.New("parameter").Parse(parameterFmt))
}

type Parameter struct {
	Name                  string
	Description           string
	AdditionalDescription string
	Value                 ParameterValue
	Type                  ParameterType
	IsRequired            bool

	// String representation of the parameters
	DefaultValue string
	Members      []string
}

func MakeParameter(key, val string, p Parameter) Parameter {
	p.Name = key
	p.Value = ParameterValue(val)
	if p.Type == ParameterType(0) {
		p.Type = paramType(val)
	}
	p.IsRequired = true // assume anything in route URL is required
	return p
}

func (p *Parameter) Render() string {
	return render(parameterTmpl, p)
}

type ParameterValue string

func (val ParameterValue) Quote() (qval string) {
	if len(val) > 0 {
		qval = fmt.Sprintf("`%s`", string(val))
	}

	return
}

func paramType(val string) ParameterType {
	if isBool(val) {
		return Boolean
	} else if isNumber(val) {
		return Number
	} else {
		return String
	}
}

func isBool(str string) bool {
	re := regexp.MustCompile(boolRe)
	return re.MatchString(str)
}

func isNumber(str string) bool {
	re := regexp.MustCompile(numberRe)
	return re.MatchString(str)
}

func (pt ParameterType) String() string {
	switch pt {
	case Number:
		return "number"
	case Boolean:
		return "boolean"
	case EnumString:
		return "enum[string]"
	case EnumNumber:
		return "enum[number]"
	default:
		return "string"
	}
}
