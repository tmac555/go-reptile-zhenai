package parse

import (
	"regexp"
	"reptile/enginer"
)

var (
	re    = regexp.MustCompile(`<a href="(http://www.hongniang.com/user/member/id/[\d]+)" [^>]+> <img [^>]*> <p [^>]*>([^>]+)</p>`)
	urlre = regexp.MustCompile(`class="num" href="([^"]+)"`)

)

func Parsecity(content []byte, rurl string) enginer.ParseResult {
	parseresult := enginer.ParseResult{}
	//人物连接
	submatch := re.FindAllSubmatch(content, -1)

	for _, result := range submatch {
		parseresult.Request = append(parseresult.Request, enginer.Request{
			Url:       string(result[1]),
			Parsefunc:Parseperfile(string(result[2])),
		})
	}
	//下一页连接
	nextpage := urlre.FindAllSubmatch(content, -1)
	for _, url := range nextpage {
		nexturl := rurl + string(url[1])
		parseresult.Request = append(parseresult.Request, enginer.Request{
			Url:       nexturl,
			Parsefunc: Nextpage,
		})
	}

	return parseresult
}

func Parseperfile(name string)enginer.Parsefunc{
	return func(c []byte, url string) enginer.ParseResult {
		return Perfile(c,url,name)
	}

}
