package parse

import (
	"fmt"
	"regexp"
	"reptile/enginer"
	_ "reptile/model"
)


const citylistre = `<a href="([^"]+)" title="[^"]*" [^>]*>([^>]+)</a>`

func Parsecitylist(content []byte,url string) enginer.ParseResult {
	//城市连接
	//contents:=perfile.ConvertToString("������", "gb2312", "utf-8")
	//fmt.Println(contents)
	res := regexp.MustCompile(citylistre)
	submatch := res.FindAllSubmatch(content, -1)
	if submatch == nil {
		fmt.Println("server error")
	}
	parseresult := enginer.ParseResult{}
	for _, result := range submatch {
		parseresult.Request = append(parseresult.Request, enginer.Request{
			Url:string(result[1]),
			Parsefunc:Parsecity,
		})
	}

	return parseresult

}
