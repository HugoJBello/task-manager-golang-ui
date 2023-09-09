package managers

import (
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

	updateTaskManager := UpdateTaskManager{ApiManager: m.ApiManager}

	inputs := []tview.Primitive{}

	globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
	for status, tasks := range tasksStatusMap {
		pages := tview.NewPages()

		list := generateListFromTasks(&tasks, pages, updatedSelectedBoard, globalAppState, status)

		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

			if list.HasFocus() && *globalAppState.FocusedElement != 0 {
				current := list.GetCurrentItem()

				statuses := *globalAppState.Statuses
				status := statuses[*globalAppState.FocusedElement-1]

				tasks := tasksStatusMap[status]
				task := tasks[current]
				globalAppState.SelectedTask = &task

				form, _ := updateTaskManager.GenerateUpdateTaskForm(app, pages, updatedSelectedBoard, globalAppState)
				pages.RemovePage("modal")
				pages.AddPage("modal", form, true, false)

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

		pages.AddPage("list", list, true, true)

		inputs = append(inputs, pages)
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

func generateListFromTasks(tasks *[]models.Task, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState, title string) *tview.List {
	list := tview.NewList()
	for index, _ := range *tasks {
		br := (*tasks)[index]
		list.AddItem(br.TaskTitle, br.TaskBody, GetRune(index), func() {
			go func() {
				pages.SwitchToPage("modal")
			}()
		})
	}

	list.SetBorder(true).SetTitle(title)

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		//fmt.Println(index, shortcut)
	})
	list.SetHighlightFullLine(true)
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
