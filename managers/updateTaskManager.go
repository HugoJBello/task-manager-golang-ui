package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type UpdateTaskManager struct {
	ApiManager ApiManager
}

func (m *UpdateTaskManager) GenerateUpdateTaskForm(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) (tview.Primitive, error) {

	var task = globalAppState.SelectedTask

	options := []string{"done", "doing", "blocked"}
	var selected = IndexOf(task.Status, options)
	if selected == -1 {
		selected = 0
	}
	form := tview.NewForm().
		AddDropDown("Status", options, selected, func(option string, index int) {
			task.Status = option
		}).
		AddInputField("title", task.TaskTitle, 20, nil, func(text string) {
			task.TaskTitle = text
		}).
		AddInputField("Body", task.TaskBody, 20, nil, func(text string) {
			task.TaskBody = text
		}).
		AddTextArea("Tags", task.Tags, 40, 0, 0, func(text string) {
			task.Tags = text
		}).
		AddButton("Save", func() {
			go func() {
				if task.Id != 0 {
					m.UpdateTaskChanges(*task)
				} else {
					m.CreateTaskChanges(*task)
				}
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
func (m *UpdateTaskManager) CreateTaskChanges(task models.Task) {
	uuid := uuid.New().String()
	createTask := models.CreateTask{TaskId: uuid, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: task.Status, BoardId: task.BoardId, DueDate: task.DueDate}
	m.ApiManager.CreateTask(createTask)
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
