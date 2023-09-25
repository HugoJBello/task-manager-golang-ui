package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type ActionsViewManager struct {
	ApiManager ApiManager
}

func (m *ActionsViewManager) AddActionsPage(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) *tview.List {

	list := tview.NewList()

	list.AddItem("Create New Task", "adds a new task", 'c', func() {
	})

	list.AddItem("Archive all done tasks", "They will disappear from view", 'd', func() {

		tasks := globalAppState.TasksInBoard

		filteredTasks := []models.Task{}

		for index, task := range *tasks {
			if task.Status == "done" {
				filteredTasks = append(filteredTasks, (*tasks)[index])
			}
		}

		m.ApiManager.ArchiveTasks(&filteredTasks)
		go func() {
			app.Stop()
			*updatedSelectedBoard <- (*tasks)[0].BoardId
		}()

		pages.SwitchToPage("tasks_board")

	})

	return list
}
