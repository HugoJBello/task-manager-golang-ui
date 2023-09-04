package managers

import (
	"encoding/json"
	"errors"
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"io/ioutil"
	"bytes"
	"net/http"
)


const version string = "/v1"

const GetTaskRoute = version + "/task/last"
const CreateTaskRoute = version + "/task/new"
const UpdateTaskRoute = version + "/task/save"
const DeleteTaskRoute = version + "/task/delete"

const GetBoardRoute = version + "/board/last"
const CreateBoardRoute = version + "/board/new"
const UpdateBoardRoute = version + "/board/save"
const DeleteBoardRoute = version + "/board/delete"

type ApiManager struct {
	Url string
}

func (m *ApiManager) GetBoards() (*[]models.Board, error){

		currentUrl := m.Url + GetBoardRoute+"?limit=10&skip=0"

		
		resp, err := http.Get(currentUrl)
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var boardResponse models.BoardResponse
		
		json.Unmarshal(bodyGetResp, &boardResponse)
	 
		return &boardResponse.Data, nil
	}

	func (m *ApiManager) CreateBoard(board models.CreateBoard) (*[]models.Board, error){

		currentUrl := m.Url + CreateBoardRoute 
		jsonBody, _ := json.Marshal(board)

		
		resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var boardResponse models.BoardResponse
		
		json.Unmarshal(bodyGetResp, &boardResponse)
	 
		return &boardResponse.Data, nil
	}

	
	func (m *ApiManager) UpdateBoard(board models.CreateBoard) (*[]models.Board, error){

		currentUrl := m.Url + UpdateBoardRoute 
		jsonBody, _ := json.Marshal(board)

		
		resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var boardResponse models.BoardResponse
		
		json.Unmarshal(bodyGetResp, &boardResponse)
	 
		return &boardResponse.Data, nil
	}


func (m *ApiManager) GetTasksInBoard(boardId string) (*[]models.Task, error){

		currentUrl := m.Url + GetTaskRoute+"?limit=100&skip=0&boardId=" + boardId

	
		resp, err := http.Get(currentUrl)
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var taskResponse models.TaskResponse
		
		json.Unmarshal(bodyGetResp, &taskResponse)
	 
		return &taskResponse.Data, nil
	}

	func (m *ApiManager) CreateTask(task models.CreateTask) (*[]models.Task, error){

		currentUrl := m.Url + CreateTaskRoute 
		jsonBody, _ := json.Marshal(task)

		
		resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var taskResponse models.TaskResponse
		
		json.Unmarshal(bodyGetResp, &taskResponse)
	 
		return &taskResponse.Data, nil
	}

	func (m *ApiManager) UpdateTask(task models.CreateTask) (*[]models.Task, error){

		currentUrl := m.Url + UpdateTaskRoute 
		jsonBody, _ := json.Marshal(task)

		
		resp, err := http.Post(currentUrl, "application/json", bytes.NewBuffer(jsonBody))
	
		if err != nil {
			return nil, err
		}
	
		if resp.StatusCode != 200 {
			return nil, errors.New("error in api")
		}
	
		bodyGetResp, err := ioutil.ReadAll(resp.Body)
		var taskResponse models.TaskResponse
		
		json.Unmarshal(bodyGetResp, &taskResponse)
	 
		return &taskResponse.Data, nil
	}