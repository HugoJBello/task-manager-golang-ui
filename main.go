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
var uiTasksManager managers.UiTasksManager
var uiBoardsManager = managers.UiBoardsManager{}
var menusManager managers.MenusManager
var buttonBarViewManager managers.ButtonBarViewManager
var actionsViewManager managers.ActionsViewManager
var updateTaskManager managers.UpdateTaskManager

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
	uiTasksManager = managers.UiTasksManager{ApiManager: apiManager}
	historyViewManager := managers.HistoryViewManager{ApiManager: apiManager}
	buttonBarViewManager = managers.ButtonBarViewManager{ApiManager: apiManager}
	updateTaskManager = managers.UpdateTaskManager{ApiManager: apiManager}
	actionsViewManager = managers.ActionsViewManager{ApiManager: apiManager, UpdateTaskManager: updateTaskManager}

	menusManager = managers.MenusManager{ApiManager: apiManager,
		UiTasksManager: uiTasksManager, HistoryViewManager: historyViewManager,
		ButtonBarViewManager: buttonBarViewManager,
		ActionsViewManager:   actionsViewManager}
}

func main() {
	boards, _ := apiManager.GetBoards()

	selectedStatus := "doing"
	focusedElement := 0
	globalAppState := models.GlobalAppState{Boards: boards, SelectedStatus: &selectedStatus, FocusedElement: &focusedElement, Statuses: models.Statuses}

	updatedSelectedBoard := make(chan string)

	app := tview.NewApplication()
	pages := tview.NewPages()

	fmt.Println(boards)

	globalAppState.SelectedBoardId = &(*boards)[0].BoardId
	menusManager.LoadMenus(app, pages, &updatedSelectedBoard, &globalAppState, false)

	for selected := range updatedSelectedBoard {
		if selected != "none" {
			globalAppState.SelectedBoardId = &selected
			menusManager.LoadMenus(app, pages, &updatedSelectedBoard, &globalAppState, false)
		}
		app.Stop()
	}

}
