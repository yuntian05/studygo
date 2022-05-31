package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

type NilParse struct {
}

func (n NilParse) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParse) Serialize() (name string, args interface{}) {
	return "NilParse", nil
}

type FuncParser struct {
	parser ParseFunc
	Name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		Name:   name,
	}
}
