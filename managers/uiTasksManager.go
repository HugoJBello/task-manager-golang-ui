package managers

import (
	"sort"

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
	var statuses = make([]string, 0, len(tasksStatusMap))
	for k := range tasksStatusMap {
		statuses = append(statuses, k)
	}

	sort.Strings(statuses)

	globalAppState.Statuses = &statuses

	return m.generateFrameListsFromTasks(tasksStatusMap, statuses, app, updatedSelectedBoard, globalAppState)

}

func (m *UiTasksManager) generateFrameListsFromTasks(tasksStatusMap map[string][]models.Task, statuses []string, app *tview.Application, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) ([]tview.Primitive, error) {

	updateTaskManager := UpdateTaskManager{ApiManager: m.ApiManager}

	inputs := []tview.Primitive{}

	globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
	for _, status := range statuses {
		pages := tview.NewPages()
		tasks := tasksStatusMap[status]

		list := generateListFromTasks(&tasks, pages, updatedSelectedBoard, globalAppState, status)

		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

			//if list.HasFocus() && *globalAppState.FocusedElement != 0 {
			if list.HasFocus() {
				current := list.GetCurrentItem()
				status = list.GetTitle()
				//statuses := *globalAppState.Statuses
				//status := statuses[*globalAppState.FocusedElement-1]
				tasks := tasksStatusMap[status]
				task := tasks[current]
				globalAppState.SelectedTask = &task

				form, _ := updateTaskManager.GenerateUpdateTaskForm(app, pages, updatedSelectedBoard, globalAppState)
				pages.RemovePage("modal")
				pages.AddPage("modal", form, true, false)

				if event.Key() == tcell.KeyCtrlU && *globalAppState.FocusedElement != 0 {
					m.createNewTaskWithStatus("doing", task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyCtrlD && *globalAppState.FocusedElement != 0 {
					m.createNewTaskWithStatus("done", task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyCtrlB && *globalAppState.FocusedElement != 0 {
					m.createNewTaskWithStatus("blocked", task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyDelete && *globalAppState.FocusedElement != 0 {
					m.deleteTask(task, app, updatedSelectedBoard)
				} else if event.Key() == tcell.KeyCtrlN && *globalAppState.FocusedElement != 0 {
					globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
					form, _ := updateTaskManager.GenerateUpdateTaskForm(app, pages, updatedSelectedBoard, globalAppState)
					pages.RemovePage("modal")
					pages.AddPage("modal", form, true, true)
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
	var taskId = task.TaskId
	if taskId == "" {
		taskId = task.TaskTitle
	}
	createTask := models.CreateTask{TaskId: taskId, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: newStatus, BoardId: task.BoardId, DueDate: task.DueDate}
	m.ApiManager.UpdateTask(createTask)
	go func() {
		app.Stop()
		*updatedSelectedBoard <- task.BoardId
	}()
}

func (m *UiTasksManager) deleteTask(task models.Task, app *tview.Application, updatedSelectedBoard *chan string) {

	m.ApiManager.DeleteTask(task.TaskId)
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
	if (*tasks)[0].Status == "doing" {
		list.SetBackgroundColor(tcell.Color100)
	}
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
