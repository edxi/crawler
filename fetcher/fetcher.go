package fetcher

import (
	"bufio"
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
	// resp, err := http.Get(url)

	//返回值是403的时候，尝试实用https获取页面内容
	//if resp.StatusCode > 200 {
	//	fmt.Println("**********************************")
	//	u := strings.Replace(url, "http", "https", 1)
	//	fmt.Println(strings.Replace(url, "http", "https", 1))
	//	fmt.Printf("fetching %v\n", u)
	//	resp1, err1 := http.Get(u)
	//	//user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36
	//	fmt.Println(resp1.StatusCode)
	//	if err1 != nil {
	//		panic(err1)
	//	}
	//	defer resp1.Body.Close()

	//}

	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	// }

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
