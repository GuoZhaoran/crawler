package parser

import (
	"depthLearn/goCrawler/engine"
	"depthLearn/goCrawler/model"
	"regexp"
	"strconv"
	"strings"
)

//预先编译好要提取信息的正则表达式
var priceRe  = regexp.MustCompile(`<span class="c_ff552e">[\s\S]*<b[\s\S]+class="f36 strongbox">(\d+)</b>[^<]+</span>`)      //租房价格
var leaseRe = regexp.MustCompile(`<span[\s\S]+class="c_888 mr_15">租赁方式：</span><span>([^<]+)</span>`)
var houseStyleRe = regexp.MustCompile(`<span[\s\S]+class="c_888 mr_15">房屋类型：</span><span>([^<]+)</span>`)
var communityRe = regexp.MustCompile(`<span class="c_888 mr_15">所在小区：</span><span><a[\s\S]+>([^<]+)</a></span>`)
var addressRe = regexp.MustCompile(`<span[\s\S]+class="c_888 mr_15">详细地址：</span><span[\s\S]+class="dz" >([^<]+)</span>`)

func ParseRoomMsg(
	contents []byte, name string) engine.ParseResult{
	profile := model.Profile{}
    profile.Title = name
	price, err := strconv.Atoi(
		extractString(contents, priceRe))
	if err == nil {
		profile.Price = price
	}

	profile.LeaseStyle = extractString(contents, leaseRe)
	profile.HouseStyle = extractString(contents, houseStyleRe)
	profile.Community = extractString(contents, communityRe)
	profile.Address = extractString(contents, addressRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

//封装正则匹配函数,提取字符串
func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return strings.Trim(string(match[1]), " ")
	} else {
		return ""
	}
}