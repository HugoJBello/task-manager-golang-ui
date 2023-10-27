package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type ActionsViewManager struct {
	ApiManager        ApiManager
	UpdateTaskManager UpdateTaskManager
}

func (m *ActionsViewManager) AddActionsPage(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) *tview.Frame {

	list := tview.NewList()

	for index, _ := range *globalAppState.Boards {
		br := (*globalAppState.Boards)[index]
		list.AddItem(br.BoardTitle, br.BoardBody, GetRune(index), func() {
			go func() {
				app.Stop()
				*globalAppState.RefreshApp <- br.BoardId
			}()
		})
	}

	list.AddItem("History", "Access history", 'h', func() {
		globalAppState.RefreshBlocked = true
		pages.SwitchToPage("historic")
	})

	list.AddItem("Create New Task", "adds a new task", 'c', func() {
		globalAppState.RefreshBlocked = true
		globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
		form, _ := m.UpdateTaskManager.GenerateUpdateTaskForm(app, pages, globalAppState)
		pages.AddPage("modal", form, true, true)
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
			*globalAppState.RefreshApp <- (*tasks)[0].BoardId
		}()

		pages.SwitchToPage("tasks_board")

	})

	list.AddItem("Quit", "Press to exit", 'q', func() {
		pages.SwitchToPage("tasks_board")
	})

	frame := tview.NewFrame(list).SetBorders(2, 2, 2, 2, 4, 4)

	return frame
}
