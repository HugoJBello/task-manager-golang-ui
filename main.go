// Demo code for the List primitive.
package main

import (
	"fmt"
	"os"

	"github.com/HugoJBello/task-manager-golang-ui/managers"
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
	"github.com/subosito/gotenv"
)

var apiManager managers.ApiManager
var uiTasksManager = managers.UiTasksManager{}
var uiBoardsManager = managers.UiBoardsManager{}

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
}

func main() {
	boards, _ := apiManager.GetBoards()

	globalAppState := models.GlobalAppState{Boards: boards}

	updatedSelectedBoard := make(chan string)
	updatedSelectedTask := make(chan string)

	app := tview.NewApplication()
	fmt.Println(boards)
	listBoards, _ := uiBoardsManager.GetBoardsListUi(boards, app, &globalAppState, &updatedSelectedBoard)

	globalAppState.SelectedBoardId = &(*boards)[0].BoardId
	reloadMenus(listBoards, app, &updatedSelectedTask, &globalAppState)

	for selected := range updatedSelectedBoard {
		if selected != "none" {
			globalAppState.SelectedBoardId = &selected
			reloadMenus(listBoards, app, &updatedSelectedTask, &globalAppState)
		}
		app.Stop()
	}

}

func reloadMenus(listBoards *tview.List, app *tview.Application, updatedSelectedTask *chan string, globalAppState *models.GlobalAppState) {

	tasksInBoard, _ := apiManager.GetTasksInBoard(*globalAppState.SelectedBoardId)
	globalAppState.TasksInBoard = tasksInBoard

	tasksList, _ := uiTasksManager.GetTasksListUi(app, updatedSelectedTask, globalAppState)

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(listBoards, 0, 1, true).
		AddItem(tasksList, 0, 2, true)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
