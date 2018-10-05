package parser

import (
    "depthLearn/goCrawler/engine"
	"regexp"
	"strings"
)
const housesRe = `<a href="(//short.58.com/[^"]*?)"[^>]*>([^<]+)</a>`
const nextPage = `<a class="next" href="([^>]+)"><span>下一页</span></a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(housesRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//格式化抓取的url
		fmtUrl := strings.Replace(string(m[1]), "/", "https:/", 1)
		result.Items = append(
			result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: fmtUrl,
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseRoomMsg(c, name)
				},
			})
	}

	nextRe := regexp.MustCompile(nextPage)
	linkMatch := nextRe.FindStringSubmatch(string(contents))
	if len(linkMatch) >= 2 {
	result.Requests = append(
		result.Requests, engine.Request{
			Url:linkMatch[1],
			ParserFunc:ParseCity,
		},
	)}


	return result
}
