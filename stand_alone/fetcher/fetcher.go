package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 创建错误的方式有两种，一种是使用 errors.New(text string)，另一种如下所示
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 获取网页编码，并将其转换成 UTF-8 编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(nil)
	}

	return all, nil
}

// determineEncoding 使用 golang.org/x/net/html 库检测网页编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 调用 peek 函数可使流能重复读，检测网页编码只需前 1024 个字节即可
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	// TODO 返回值含义有待确认
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
