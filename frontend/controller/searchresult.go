package controller

import (
	"crawler/frontend/view"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"net/http"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elasticsearch.Client
}

// url 格式： http://localhost:8888/search?q=男 已购房&from=10
func (s SearchResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 获取 url 中参数 q 后面的内容
	query := strings.TrimSpace(request.FormValue("q"))
	// 获取 url 中的分页值，也就是放在 from 后面的数值
	from, err := strconv.Atoi(request.FormValue("from"))
	if err != nil {
		from = 0
	}

	fmt.Fprintf(writer, "query=%s, from=%d", query, from)

}
