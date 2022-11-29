package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"
)

var (
	Debug = false
	ServerDns = "proxies.gay"
)

type Task struct {
	TaskID   string
	Token    string
	OutToken string
}

type CreateTaskResp struct {
	Token string
}

func HandleError(Err error) bool {
	if Err != nil {
		if Debug {
			fmt.Println(Err.Error())
		}
		return true
	}

	return false
}

func doReq(Method string, Path string, Content map[string]string) string {
	jsonData, _ := json.Marshal(Content)

	request, error := http.NewRequest(Method, fmt.Sprintf("https://%s%s", ServerDns, Path), bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}

func getTask() (Task, error) {
	var Response CreateTaskResp
	T := Task{
		TaskID: strings.Split(fmt.Sprintf("%s", js.Global().Get("location").Get("href")), "task=")[1],
	}

	Err := json.Unmarshal([]byte(doReq("POST", "/create_task", map[string]string{
		"task_id": T.TaskID,
	})), &Response)

	if HandleError(Err) {
		return T, Err
	}

	T.Token = Response.Token
	return T, nil
}

func (T *Task) Pow() {
	x := func(input, key string) (output string) {
		for i := 0; i < len(input); i++ {
			output += string(input[i] ^ key[i%len(key)])
		}
		return hex.EncodeToString([]byte(output))
	}

	T.OutToken = base64.RawURLEncoding.EncodeToString([]byte(x(T.Token, T.TaskID)))
}

func (T *Task) Submit() {
	doReq("POST", "/submit_captcha", map[string]string{
		"response": T.OutToken,
		"task_id":  T.TaskID,
	})
}

func main() {
	T, Err := getTask()
	if HandleError(Err) {
		return
	}

	T.Pow()
	T.Submit()
}
