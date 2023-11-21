package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ZabbixResponse struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
}

type AuthResponse struct {
	Token string `json:"result"`
}

type HostGroup struct {
	GroupId   string `json:"groupid"`
	GroupName string `json:"name"`
}

func main() {
	// 设置Zabbix服务器的URL、用户名和密码
	zabbixURL := "http://your-zabbix-server/zabbix/api_jsonrpc.php"
	username := "your-username"
	password := "your-password"

	// 发送身份验证请求
	authToken, err := authenticate(zabbixURL, username, password)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return
	}

	// 获取Zabbix中的所有主机组
	groups, err := getHostGroups(zabbixURL, authToken)
	if err != nil {
		fmt.Println("Failed to get host groups:", err)
		return
	}

	// 打印主机组信息
	for _, group := range groups {
		fmt.Printf("Group ID: %s, Group Name: %s\n", group.GroupId, group.GroupName)
	}
}

func authenticate(zabbixURL, username, password string) (string, error) {
	// 构建认证请求数据
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "user.login",
		"params": map[string]string{
			"user":     username,
			"password": password,
		},
		"id": 1,
	}

	// 发送认证请求
	response, err := sendRequest(zabbixURL, payload)
	if err != nil {
		return "", err
	}

	// 解析认证响应
	var authResponse ZabbixResponse
	err = json.Unmarshal(response, &authResponse)
	if err != nil {
		return "", err
	}

	// 检查认证是否成功，返回令牌
	if authResponse.Error != nil {
		return "", fmt.Errorf("authentication failed: %v", authResponse.Error)
	}

	if token, ok := authResponse.Result.(string); ok {
		return token, nil
	}

	return "", fmt.Errorf("invalid authentication response")
}

func getHostGroups(zabbixURL, authToken string) ([]HostGroup, error) {
	// 构建获取主机组请求数据
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "hostgroup.get",
		"params": map[string]interface{}{
			"output": "extend",
		},
		"auth": authToken,
		"id":   1,
	}

	// 发送获取主机组请求
	response, err := sendRequest(zabbixURL, payload)
	if err != nil {
		return nil, err
	}

	// 解析获取主机组响应
	var groupResponse ZabbixResponse
	err = json.Unmarshal(response, &groupResponse)
	if err != nil {
		return nil, err
	}

	// 检查获取主机组是否成功，返回主机组信息
	if groupResponse.Error != nil {
		return nil, fmt.Errorf("failed to get host groups: %v", groupResponse.Error)
	}

	groups, ok := groupResponse.Result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid host group response")
	}

	var hostGroups []HostGroup
	for _, group := range groups {
		groupMap, ok := group.(map[string]interface{})
		if !ok {
			continue
		}

		groupId, ok := groupMap["groupid"].(string)
		if !ok {
			continue
		}

		groupName, ok := groupMap["name"].(string)
		if !ok {
			continue
		}

		hostGroup := HostGroup{
			GroupId:   groupId,
			GroupName: groupName,
		}

		hostGroups = append(hostGroups, hostGroup)
	}

	return hostGroups, nil
}

func sendRequest(url string, payload map[string]interface{}) ([]byte, error) {
	// 构建请求
	// struct ===> []byte ===> string ===> io.Reader
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 发送请求
	response, err := http.Post(url, "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
