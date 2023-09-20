package managers

import (
	"fmt"
	"strconv"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type HistoryViewManager struct {
	ApiManager ApiManager
}

func (m *HistoryViewManager) AddHistoryPage(app *tview.Application, pages *tview.Pages, globalAppState *models.GlobalAppState) *tview.List {

	tasksHistory, _ := m.ApiManager.GetTasksHistory(100)
	fmt.Println("----->", tasksHistory)

	list := tview.NewList()

	if tasksHistory != nil {
		for _, taskHistory := range *tasksHistory {
			list.AddItem(m.getHistoryText(taskHistory), m.getHistorySubText(taskHistory), '-', nil)
		}
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		go func() {
			pages.SwitchToPage("tasks_board")
		}()
	})

	return list
}

func (m *HistoryViewManager) getHistoryText(taskHistory models.TaskHistory) string {
	title := taskHistory.TaskTitle
	oldStatus := taskHistory.OldStatus
	newStatus := taskHistory.NewStatus
	return title + oldStatus + " --> " + newStatus
}

func (m *HistoryViewManager) getHistorySubText(taskHistory models.TaskHistory) string {
	var dificulty = "1"
	if taskHistory.Dificulty != nil {
		dificulty = strconv.Itoa(*taskHistory.Dificulty)
	}

	var priority = "1"
	if taskHistory.Priority != nil {
		priority = strconv.Itoa(*taskHistory.Priority)
	}

	date := taskHistory.EditedAt.Format("2006-01-02 15:04:05")

	return date + " DIF: " + dificulty + " PRIOR: " + priority
}
