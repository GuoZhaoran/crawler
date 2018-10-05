package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
)

//根据网页链接获取到网页内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)

	return ioutil.ReadAll(bodyReader)
}
