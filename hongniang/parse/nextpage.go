package parse

import (
	"regexp"
	"reptile/enginer"
)

//nexpage url
func Nextpage(content []byte, nexturl string) enginer.ParseResult {
	parseresult := enginer.ParseResult{}
	//人物连接
	pre := regexp.MustCompile(`<a href="(http://www.hongniang.com/user/member/id/[\d]+)" [^>]+> <img [^>]*> <p [^>]*>([^>]+)</p>`)
	submatch := pre.FindAllSubmatch(content, -1)
	for _, result := range submatch {
		parseresult.Request = append(parseresult.Request, enginer.Request{
			Url:       string(result[1]),
			Parsefunc: Parseperfile(string(result[2])),
		})
	}
	return parseresult
}
