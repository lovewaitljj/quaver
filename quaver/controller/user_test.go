package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {

	type test struct { // 定义test结构体
		param string
		want  ResCode
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple": {param: `{"username": "blingder"}`, want: CodeInvalidParam},
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/douyin/user/register/"
	r.POST(url, Register)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(tc.param)))
			req.Header.Set("Content-Type", "application/json") // 告诉服务端消息主体是序列化后的 JSON 字符串
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			// 判断响应的内容是不是按预期返回了请求参数的错误

			// 将响应的内容反序列化到ResponseData 然后判断字段与预期是否一致
			res := new(ResponseData)
			if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
				t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
			}
			assert.Equal(t, res.Code, tc.want)
		})
	}

}
