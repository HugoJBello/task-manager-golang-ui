package managers

import (
	"strconv"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type UpdateTaskManager struct {
	ApiManager ApiManager
}

func (m *UpdateTaskManager) GenerateUpdateTaskForm(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) (tview.Primitive, error) {

	var task = globalAppState.SelectedTask

	options := []string{"done", "doing", "blocked", "todo"}
	var selected = IndexOf(task.Status, options)
	if selected == -1 {
		selected = 0
	}

	var dificulty = "1"
	var priority = "1"

	if task.Dificulty != nil {
		dificulty = strconv.Itoa(*task.Dificulty)
	}
	if task.Priority != nil {
		priority = strconv.Itoa(*task.Priority)
	}

	form := tview.NewForm().
		AddDropDown("Status", options, selected, func(option string, index int) {
			task.Status = option
		}).
		AddInputField("title", task.TaskTitle, 20, nil, func(text string) {
			task.TaskTitle = text
		}).
		AddInputField("Dificulty", dificulty, 20, nil, func(text string) {
			marks, err := strconv.Atoi(text)
			var dificulty = 1
			if err == nil {
				dificulty = marks
			}
			task.Dificulty = &dificulty
		}).
		AddInputField("Priority", priority, 20, nil, func(text string) {
			marks, err := strconv.Atoi(text)
			var priority = 1
			if err == nil {
				priority = marks
			}
			task.Priority = &priority
		}).
		AddInputField("Tags", task.TaskBody, 20, nil, func(text string) {
			task.Tags = text
		}).
		AddTextArea("Body", task.Tags, 40, 0, 0, func(text string) {
			task.TaskBody = text
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
		TaskBody: task.TaskBody, Tags: task.Tags, Status: task.Status, BoardId: task.BoardId, DueDate: task.DueDate, Priority: task.Priority, Dificulty: task.Dificulty}
	m.ApiManager.UpdateTask(createTask)
}
func (m *UpdateTaskManager) CreateTaskChanges(task models.Task) {
	uuid := uuid.New().String()
	createTask := models.CreateTask{TaskId: uuid, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: task.Status, BoardId: task.BoardId, DueDate: task.DueDate, Priority: task.Priority, Dificulty: task.Dificulty}
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
