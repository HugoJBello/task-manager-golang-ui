package managers

import (
	"sort"
	"strconv"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UiTasksManager struct {
	ApiManager ApiManager
}

func (m *UiTasksManager) GetTasksListUi(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) (taskLists []tview.List, err error) {

	tasks := globalAppState.TasksInBoard
	tasksStatusMap := m.organizeTasksUsingStatus(*tasks)
	statuses := make([]string, 0, len(tasksStatusMap))
	for k := range tasksStatusMap {
		statuses = append(statuses, k)
	}

	sort.Strings(statuses)

	taskLists, err = m.generateFrameListsFromTasks(pages, tasksStatusMap, statuses, app, globalAppState)
	return taskLists, err

}

func (m *UiTasksManager) generateFrameListsFromTasks(pages *tview.Pages, tasksStatusMap map[string][]models.Task, statuses []string,
	app *tview.Application, globalAppState *models.GlobalAppState) ([]tview.List, error) {

	updateTaskManager := UpdateTaskManager{ApiManager: m.ApiManager}

	inputs := []tview.List{}
	priority := 1
	dificulty := 1

	globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", Priority: &priority, Dificulty: &dificulty, TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
	for _, status := range statuses {
		tasks := tasksStatusMap[status]

		list := m.generateListFromTasks(&tasks, tasksStatusMap, pages, globalAppState, status, app)

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

				if event.Key() == tcell.KeyCtrlU {
					m.updateNewTaskWithStatus("doing", task, globalAppState, app)
				} else if event.Key() == tcell.KeyCtrlD {
					m.updateNewTaskWithStatus("done", task, globalAppState, app)
				} else if event.Key() == tcell.KeyCtrlB {
					m.updateNewTaskWithStatus("blocked", task, globalAppState, app)
				} else if event.Key() == tcell.KeyDelete {
					m.deleteTask(task, globalAppState, app)
				} else if event.Key() == tcell.KeyCtrlN {
					globalAppState.SelectedTask = &models.Task{Id: 0, TaskId: "", TaskTitle: "", TaskBody: "", Tags: "", Status: "", BoardId: *globalAppState.SelectedBoardId}
					form, _ := updateTaskManager.GenerateUpdateTaskForm(app, pages, globalAppState)
					globalAppState.RefreshBlocked = true
					pages.AddPage("modal", form, true, true)
				} else if event.Key() == tcell.KeyCtrlA {
					pages.SwitchToPage("actions")
				}

			}
			return event
		})

		inputs = append(inputs, *list)
	}
	return inputs, nil
}

func (m *UiTasksManager) updateNewTaskWithStatus(newStatus string, task models.Task, globalAppState *models.GlobalAppState, app *tview.Application) {
	var taskId = task.TaskId
	if taskId == "" {
		taskId = task.TaskTitle
	}
	createTask := models.CreateTask{TaskId: taskId, TaskTitle: task.TaskTitle,
		TaskBody: task.TaskBody, Tags: task.Tags, Status: newStatus, BoardId: task.BoardId, DueDate: task.DueDate}
	m.ApiManager.UpdateTask(createTask)
	go func() {
		app.Stop()
		*globalAppState.RefreshApp <- task.BoardId
	}()
}

func (m *UiTasksManager) deleteTask(task models.Task, globalAppState *models.GlobalAppState, app *tview.Application) {

	m.ApiManager.DeleteTask(task.TaskId)
	go func() {
		app.Stop()
		*globalAppState.RefreshApp <- task.BoardId
	}()
}

func (m *UiTasksManager) generateListFromTasks(tasks *[]models.Task, tasksStatusMap map[string][]models.Task, pages *tview.Pages, globalAppState *models.GlobalAppState, title string, app *tview.Application) *tview.List {
	updateTaskManager := UpdateTaskManager{ApiManager: m.ApiManager}

	list := tview.NewList()
	for index, _ := range *tasks {
		br := (*tasks)[index]
		var subtext = ""
		if br.Dificulty != nil {
			subtext = subtext + "[::bl] DIF: " + strconv.Itoa(*br.Dificulty) + "[-:-:-:-]"
		} else {
			subtext = subtext + "[::bl] DIF: 1[-:-:-:-]"
		}
		if br.Priority != nil {
			subtext = subtext + " PRIO: " + strconv.Itoa(*br.Priority)
		} else {
			subtext = subtext + " PRIO: 1"
		}

		subtext = subtext + " - " + br.TaskBody
		list.AddItem(br.TaskTitle, subtext, GetRune(index), func() {
			globalAppState.RefreshBlocked = true
			if pages.HasPage("modal") {
				pages.SwitchToPage("modal")
			}
		})
	}

	list.SetBorder(true).SetTitle(title)

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		current := index
		status := list.GetTitle()

		tasks := tasksStatusMap[status]
		task := tasks[current]
		globalAppState.SelectedTask = &task

		form, _ := updateTaskManager.GenerateUpdateTaskForm(app, pages, globalAppState)
		pages.AddPage("modal", form, true, true)

	})
	list.SetHighlightFullLine(true)
	if (*tasks)[0].Status == "doing" {
		list.SetBackgroundColor(tcell.ColorDarkSlateGray)
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

	for status := range result {
		sort.SliceStable(result[status], func(i, j int) bool {
			var prioi = 1
			var prioj = 1
			if result[status][i].Priority != nil {
				prioi = *result[status][i].Priority
			}
			if result[status][j].Priority != nil {
				prioj = *result[status][j].Priority
			}
			return prioi > prioj
		})
	}

	return result
}
