/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 17:42:58
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 17:49:59
 */

package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// http json 请求
//
// url 地址
//
// data 发送数据
//
// result 回复数据绑定
func HttpJsonPost(url string, data interface{}, result interface{}) error {
	var sendData []byte
	var err error

	if data != nil {
		tmp, err := json.Marshal(data)
		if err != nil {
			return err
		}
		sendData = tmp
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(sendData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if result != nil {
		err = json.Unmarshal(respData, result)
		if err != nil {
			return err
		}
	}
	return nil
}
