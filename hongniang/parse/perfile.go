package parse

import (
	"regexp"
	"reptile/enginer"
	"reptile/model"
)

var (
	idre        = regexp.MustCompile(`<div [^>]*>Loveid:([^>]+)</div>`)
	agere       = regexp.MustCompile(`<li><span>年龄：</span>([^>]+)</li>`)
	heire       = regexp.MustCompile(`<li><span>身高：</span>([^>]+)</li>`)
	incomere    = regexp.MustCompile(`<li><span>年收入：</span>([^>]+)</li>`)
	maritalre   = regexp.MustCompile(`<li><span>婚况：</span>([^>]+)</li>`)
	educationre = regexp.MustCompile(`<li><span>学历：</span>([^>]+)</li>`)
	addre       = regexp.MustCompile(`<li><span>工作地：</span>([^>]+)</li>`)
)

func Perfile(content []byte, url, name string) enginer.ParseResult {
	perfile := perfile.Archives{}
	id := matching(content, idre)
	perfile.Name = name
	perfile.Age = matching(content, agere)
	perfile.Height = matching(content, heire)
	perfile.Income = matching(content, incomere)
	perfile.Marital = matching(content, maritalre)
	perfile.Education = matching(content, educationre)
	perfile.Address = matching(content, addre)
	//id url parsrresult
	result := enginer.ParseResult{
		Item: []enginer.Perfileitem{
			{
				Id:    id,
				Type:  "hongniang",
				Url:   url,
			},
		},
	}
	return result
}
func matching(c []byte, re *regexp.Regexp) string {
	submatch := re.FindSubmatch(c)
	if len(submatch) >= 2 {
		return string(submatch[1])
	} else {
		return ""
	}

}
