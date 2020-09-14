package tools

import (
	_struct "emafs/struct"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeremywohl/flatten"
	"net/http"
	"strings"
)

func HttpReqGetMap(url string) map[string]interface{} {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	formJson := make(gin.H)
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&formJson)
	jsonM, _ := flatten.Flatten(formJson, "", flatten.DotStyle)
	return jsonM
}

func HttpReqPostPush(argMap _struct.Push) {
	fmt.Println(argMap)
	client := &http.Client{}
	str := "user_account=" + argMap.UserAccount + "&title=" + argMap.Title + "&description=" + argMap.Description +
		"&stricted_package_name=com.linkft.androidapp&pass_through=0&notify_type=-1"
	req, _ := http.NewRequest("POST", "https://api.xmpush.xiaomi.com/v2/message/user_account", strings.NewReader(str))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "key=22tJ/P7DGB9nWEAHj+L9sw==")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
}
