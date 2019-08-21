package enginer

type Parsefunc func(content []byte, url string) ParseResult
type Request struct {
	Url       string
	Parsefunc Parsefunc
}
type ParseResult struct {
	Request []Request
	Item    []Perfileitem
}
type Perfileitem struct {
	Id    string
	Type  string
	Url   string
	Pitem interface{}
}


