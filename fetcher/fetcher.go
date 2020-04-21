package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 获得url返回网页内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 把网页放到bufio里，交给determinEncoding函数去确认网页编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	// 网页内容转码存入一个ioreader
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 读取ioreader，把转码后的网页以byte[]返回
	return ioutil.ReadAll(utf8Reader)

}

// 通过放在bufio里的网页内容返回encoding
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//charset.DetermineEncoding读取bufio前1024字节
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
