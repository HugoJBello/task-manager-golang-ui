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

func init() {
	gotenv.Load()
	apiManager = managers.ApiManager{Url: os.Getenv("API_URL")}
	uiTasksManager = managers.UiTasksManager{ApiManager: apiManager}
	menusManager = managers.MenusManager{ApiManager: apiManager, UiTasksManager: uiTasksManager}
}

func main() {
	boards, _ := apiManager.GetBoards()

	globalAppState := models.GlobalAppState{Boards: boards}

	updatedSelectedBoard := make(chan string)

	app := tview.NewApplication()
	pages := tview.NewPages()

	fmt.Println(boards)
	sideMenu, _ := uiBoardsManager.GetBoardsListUi(boards, app, pages, &globalAppState, &updatedSelectedBoard)

	globalAppState.SelectedBoardId = &(*boards)[0].BoardId
	menusManager.LoadMenus(sideMenu, app, pages, &updatedSelectedBoard, &globalAppState, false)

	for selected := range updatedSelectedBoard {
		if selected != "none" {
			globalAppState.SelectedBoardId = &selected
			menusManager.LoadMenus(sideMenu, app, pages, &updatedSelectedBoard, &globalAppState, false)
		}
		app.Stop()
	}

}
