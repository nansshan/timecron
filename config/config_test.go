package config

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

//写入json文件
func Test_writeFile(t *testing.T) {
	data := `{
    "default_sbox_drive_id": "97951021",
    "role": "user",
    "user_name": "155***810",
    "need_link": false,
    "expire_time": "2021-08-30T10:02:45Z",
    "pin_setup": false,
    "need_rp_verify": false,
    "avatar": "https://ccp-bj29-bj-1592982087.oss-cn-beijing.aliyuncs.com/pds%2Favatar%2F7f409f0041ee489ea34075542c37c1b4?x-oss-access-key-id=LTAIsE5mAn2F493Q&x-oss-expires=1630314165&x-oss-signature=hb4YSCwu1gi7MkfECGll15W2x7z21PJ3BqvzweJ%2F76U%3D&x-oss-signature-version=OSS2",
    "token_type": "Bearer",
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI1MTJhNzA2YTY4OWU0ZWU3YWFmZTAzODJkMjk4ZjkzYSIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiU1RPUkFHRS5BTExcIixcIlNUT1JBR0VGSUxFLkxJU1RcIixcIkJBVENIXCIsXCJPQVVUSC5BTExcIixcIklNQUdFLkFMTFwiLFwiSU5WSVRFLkFMTFwiLFwiQUNDT1VOVC5BTExcIl0sXCJyb2xlXCI6XCJ1c2VyXCIsXCJyZWZcIjpcImh0dHBzOi8vd3d3LmFsaXl1bmRyaXZlLmNvbS9cIn0iLCJleHAiOjE2MzAzMTc3NjUsImlhdCI6MTYzMDMxMDUwNX0.ZD1nF1o9OsyamZYveBGUzfMSLUN5BjMfiykLdUwEBz0Np-iSiBqkQBdQ--tI6OVBdEYEOs67OKGIkO-mUoFIPFxdXMrpaOaoxWO8aCvXJ-qIm8EtIhhllV_oNsieu2qpHOu392p3AeRloe3BibiOU8z7TuwjcnnQIMZENa_8L-Y",
    "default_drive_id": "87951021",
    "refresh_token": "e60a4522c06c4fbea953b8d729936cb4",
    "is_first_login": false,
    "user_id": "512a706a689e4ee7aafe0382d298f93a",
    "nick_name": "lnl",
    "expires_in": 7200,
    "status": "enabled"
	}`
	var dat map[string]interface{}
	bytdata := []byte(data)
	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(bytdata, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	f, err := os.Create("./test.json")
	if err != nil {
		// 创建文件失败处理
		fmt.Println("文件创建失败0")
	} else {
		_, err = f.Write([]byte(data))
		if err != nil {
			// 写入失败处理
			fmt.Println("文件写入失败")
		}
	}
	defer f.Close()
}
