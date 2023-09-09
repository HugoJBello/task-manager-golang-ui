package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type UpdateTaskManager struct {
	ApiManager ApiManager
}

func (m *UpdateTaskManager) GenerateUpdateTaskForm(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) (tview.Primitive, error) {

	task := globalAppState.SelectedTask

	options := []string{"done", "doing", "blocked"}
	var selected = IndexOf(task.Status, options)
	if selected == -1 {
		selected = 0
	}
	form := tview.NewForm().
		AddDropDown("Status", options, selected, func(option string, index int) {
			task.Status = option
			m.UpdateTaskChanges(*task)
		}).
		AddInputField("title", task.TaskTitle, 20, nil, func(text string) {
			task.TaskTitle = text
			m.UpdateTaskChanges(*task)
		}).
		AddInputField("Body", task.TaskBody, 20, nil, func(text string) {
			task.TaskBody = text
			m.UpdateTaskChanges(*task)
		}).
		AddTextArea("Tags", task.Tags, 40, 0, 0, func(text string) {
			task.Tags = text
			m.UpdateTaskChanges(*task)
		}).
		AddButton("Save", func() {
			go func() {
				app.Stop()
				*updatedSelectedBoard <- task.BoardId
			}()
			pages.SwitchToPage("list")
		}).
		AddButton("Quit", func() {
			pages.SwitchToPage("list")
		})

	return form, nil

}

func (m *UpdateTaskManager) UpdateTaskChanges(task models.Task) {
	createTask := models.CreateTask{TaskId: task.TaskId, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: task.Status, BoardId: task.BoardId, DueDate: task.DueDate}
	m.ApiManager.UpdateTask(createTask)
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
