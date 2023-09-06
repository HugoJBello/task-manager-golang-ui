package managers

import (
	"fmt"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UiTasksManager struct {
	ApiManager ApiManager
}

func (m *UiTasksManager) GetTasksListUi(app *tview.Application, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) ([]tview.Primitive, error) {

	tasks := globalAppState.TasksInBoard
	tasksStatusMap := m.organizeTasksUsingStatus(*tasks)
	statuses := make([]string, 0, len(tasksStatusMap))
	for k := range tasksStatusMap {
		statuses = append(statuses, k)
	}

	globalAppState.Statuses = &statuses

	return m.generateFrameListsFromTasks(tasksStatusMap, app, updatedSelectedBoard, globalAppState)

}

func (m *UiTasksManager) generateFrameListsFromTasks(tasksStatusMap map[string][]models.Task, app *tview.Application, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) ([]tview.Primitive, error) {
	inputs := []tview.Primitive{}

	for status, tasks := range tasksStatusMap {
		list := generateListFromTasks(&tasks, updatedSelectedBoard, globalAppState, status)

		list.Focus(func(p tview.Primitive) {
			fmt.Println("----", status)
		})

		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if list.HasFocus() {
				current := list.GetCurrentItem()

				statuses := *globalAppState.Statuses
				status := statuses[*globalAppState.FocusedElement-1]
				task := tasksStatusMap[status][current]

				if event.Key() == tcell.KeyCtrlU {
					m.createNewTaskWithStatus("doing", task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyCtrlD {
					m.createNewTaskWithStatus("done", task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyCtrlB {
					m.createNewTaskWithStatus("blocked", task, app, updatedSelectedBoard)
				}
			}
			return event
		})

		inputs = append(inputs, list)
	}

	//AddCycleFocus(flex, app, inputs)
	return inputs, nil
}

func (m *UiTasksManager) createNewTaskWithStatus(newStatus string, task models.Task, app *tview.Application, updatedSelectedBoard *chan string) {
	createTask := models.CreateTask{TaskId: task.TaskId, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: newStatus, BoardId: task.BoardId, DueDate: task.DueDate}
	m.ApiManager.UpdateTask(createTask)
	go func() {
		app.Stop()
		*updatedSelectedBoard <- task.BoardId
	}()
}

func generateListFromTasks(tasks *[]models.Task, updatedSelectedTask *chan string, globalAppState *models.GlobalAppState, title string) *tview.List {
	list := tview.NewList()
	for index, _ := range *tasks {
		br := (*tasks)[index]
		list.AddItem(br.TaskTitle, br.TaskBody, GetRune(index), func() {
			go func() {
				*updatedSelectedTask <- br.TaskId
			}()
		})
	}

	list.SetBorder(true).SetTitle(title)

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		//fmt.Println(index, shortcut)
	})
	return list
}

func (m *UiTasksManager) organizeTasksUsingStatus(tasks []models.Task) map[string][]models.Task {
	result := make(map[string][]models.Task)

	for i := 0; i < len(tasks); i++ {
		tk := tasks[i]
		if result[tk.Status] == nil {
			result[tk.Status] = []models.Task{tk}
		} else {
			result[tk.Status] = append(result[tk.Status], tk)
		}
	}

	return result
}
