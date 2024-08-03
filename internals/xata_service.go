package internals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var xataAPIKey = "xau_Gj7z32BSKPjRXfS4tWJNyyhJz4v6Ga1u3"
var baseURL = "https://Himanshu-Singh-s-workspace-aa7ln8.us-east-1.xata.sh/db/todo"

func createRequest(method, url string, bodyData *bytes.Buffer) (*http.Request, error) {
	var req *http.Request
	var err error

	if method == "GET" || method == "DELETE" || bodyData == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bodyData)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	return req, nil
}


func (app *Config) CreateTodoService(newTodo *TodoRequest) (*TodoResponse,error){

	createtodo:=TodoResponse{}
	jsonData:=Todo{
		Description: newTodo.Description,
	}

	bodyData:=new(bytes.Buffer)
	json.NewEncoder(bodyData).Encode(jsonData)

	fullUrl:=fmt.Sprintf("%s:main/tables/Todo/data",baseURL)
	req,err:=createRequest("POST",fullUrl,bodyData)

	if err!=nil{
		return nil,err
	}

	client:=&http.Client{}
	resp,err:=client.Do(req)

	if err!=nil{
		return nil,err
	}

	defer resp.Body.Close()

	if err:=json.NewDecoder(resp.Body).Decode(&createtodo); err!=nil{
		return nil,err
	}

	return &createtodo,nil

}

func (app *Config) getAllTodosService() ([]*Todo, error) {
	var todos []*Todo

	fullURL := fmt.Sprintf("%s:main/tables/Todo/query", baseURL)
	client := &http.Client{}
	req, err := createRequest("POST", fullURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Records []*Todo `json:"records"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	todos = response.Records

	return todos, nil
}

/*func (app *Config) DeleteTodosService(id string) ([]*Todo, error) {
	var todos []*Todo

	fullUrl:=fmt.Sprintf("%s:main/tables/Todo/data/%s",baseURL,id)
	client := &http.Client{}
	req, err := createRequest("POST", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Records []*Todo `json:"records"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	todos = response.Records

	return todos, nil
}*/








