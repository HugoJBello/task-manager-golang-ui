package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type UiTasksManager struct {
}

func (m *UiTasksManager) GetTasksListUi(app *tview.Application, updatedSelectedTask *chan string, globalAppState *models.GlobalAppState) (*tview.List, error) {
	list := tview.NewList()

	tasks := globalAppState.TasksInBoard
	for index, _ := range *tasks {
		br := (*tasks)[index]
		list.AddItem(br.TaskTitle, br.TaskBody, GetRune(index), func() {
			go func() {
				*updatedSelectedTask <- br.TaskId
			}()
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		go func() {
			*updatedSelectedTask <- "none"
			close(*updatedSelectedTask)
		}()

		app.Stop()
	})
	list.SetBorder(true).SetTitle(*globalAppState.SelectedBoardId)

	return list, nil
}

func findCurrentBoardTitle(globalAppState *models.GlobalAppState) {

}
