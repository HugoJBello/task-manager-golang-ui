// Demo code for the List primitive.
package main

import (
	"fmt"
	"os"
	"time"

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
var timerManager managers.TimerManager

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
	uiTasksManager = managers.UiTasksManager{ApiManager: apiManager}
	historyViewManager := managers.HistoryViewManager{ApiManager: apiManager}
	buttonBarViewManager = managers.ButtonBarViewManager{ApiManager: apiManager}
	updateTaskManager = managers.UpdateTaskManager{ApiManager: apiManager}
	actionsViewManager = managers.ActionsViewManager{ApiManager: apiManager, UpdateTaskManager: updateTaskManager}
	timerManager = managers.TimerManager{}

	menusManager = managers.MenusManager{ApiManager: apiManager,
		UiTasksManager: uiTasksManager, HistoryViewManager: historyViewManager,
		ButtonBarViewManager: buttonBarViewManager,
		ActionsViewManager:   actionsViewManager}
}

func main() {
	boards, _ := apiManager.GetBoards()

	selectedStatus := "doing"
	focusedElement := 0
	refreshApp := make(chan string)
	globalAppState := models.GlobalAppState{RefreshApp: &refreshApp, Boards: boards, SelectedStatus: &selectedStatus, FocusedElement: &focusedElement, Statuses: models.Statuses,
		CurrentTime: time.Now(), RefreshBlocked: false}

	updatedSelectedBoard := make(chan string)

	app := tview.NewApplication()

	go timerManager.SetTimer(app, &globalAppState)

	pages := tview.NewPages()

	fmt.Println(boards)

	globalAppState.SelectedBoardId = &(*boards)[0].BoardId
	menusManager.LoadMenus(app, pages, &updatedSelectedBoard, &globalAppState, false)

	go func() {
		for selected := range updatedSelectedBoard {
			fmt.Println("updated selected board")
			if selected != "none" {
				globalAppState.SelectedBoardId = &selected
				menusManager.LoadMenus(app, pages, &updatedSelectedBoard, &globalAppState, false)
			}
			app.Stop()
		}
	}()

	go func() {
		for refresh := range *globalAppState.RefreshApp {
			fmt.Println("-----", refresh)
			menusManager.LoadMenus(app, pages, &updatedSelectedBoard, &globalAppState, false)

		}
	}()

	<-(chan int)(nil)

}
