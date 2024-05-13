package news_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type Params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}
type NewResponse struct {
	Code int `json:"code"`
	Data []struct {
		Index    string `json:"index"`
		Title    string `json:"title"`
		HotValue string `json:"hotValue"`
		Link     string `json:"link"`
	}
	Msg string `json:"msg"`
}

func TestNews(t *testing.T) {
	var params = Params{
		ID:   "mproPpoq6O",
		Size: 1,
	}
	reqParam, _ := json.Marshal(params)
	reqBody := bytes.NewReader([]byte(reqParam))
	httpReq, err := http.NewRequest("POST", "https://api.codelife.cc/api/top/list", reqBody)
	if err != nil {
		fmt.Println(err)
	}
	// httpReq.Header.Add("Content-Type", "application/json")
	// httpReq.Header.Add("signaturekey", "U2FsdGVkX1+tE85PyG35mvlEK28qp7ivsr8AQNAuRsk=")
	// httpReq.Header.Add("version", "1.3.56")

	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(httpRsp, err)

}
