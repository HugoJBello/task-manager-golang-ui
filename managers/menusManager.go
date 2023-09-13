package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenusManager struct {
	ApiManager         ApiManager
	UiTasksManager     UiTasksManager
	HistoryViewManager HistoryViewManager
}

func (m *MenusManager) LoadMenus(listBoards *tview.List, app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState, horizontalView bool) {

	tasksInBoard, _ := m.ApiManager.GetTasksInBoard(*globalAppState.SelectedBoardId)
	globalAppState.TasksInBoard = tasksInBoard

	i := 0
	globalAppState.FocusedElement = &i
	tasksList, _ := m.UiTasksManager.GetTasksListUi(app, updatedSelectedBoard, globalAppState)

	flex := tview.NewFlex().
		AddItem(listBoards, 0, 1, true)

	tasksFlex := tview.NewFlex()

	if horizontalView == true {
		tasksFlex.SetDirection(tview.FlexRow)
	}

	for index, _ := range tasksList {
		tasksFlex.AddItem(tasksList[index], 0, 1, true)
	}

	flex.AddItem(tasksFlex, 0, 3, false)

	historicList := m.HistoryViewManager.AddHistoryPage(app, pages, globalAppState)

	pages.AddPage("historic", historicList, true, true)

	pages.AddPage("tasks_board", flex, true, true)

	inputs := []tview.Primitive{
		listBoards,
	}
	inputs = append(inputs, tasksList...)

	AddCycleFocus(flex, app, inputs, globalAppState)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func AddCycleFocus(flex *tview.Flex, app *tview.Application, inputs []tview.Primitive, globalAppState *models.GlobalAppState) {
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlSpace {
			CycleFocus(app, inputs, false, globalAppState)
		} else if event.Key() == tcell.KeyBacktab {
			CycleFocus(app, inputs, true, globalAppState)
		}
		return event

	})
}

func CycleFocus(app *tview.Application, elements []tview.Primitive, reverse bool, globalAppState *models.GlobalAppState) {
	for i, el := range elements {
		if !el.HasFocus() {
			continue
		}

		if reverse {
			i = i - 1
			if i < 0 {
				i = len(elements) - 1
			}
		} else {
			i = i + 1
			i = i % len(elements)
		}
		app.SetFocus(elements[i])
		globalAppState.FocusedElement = &i
		return
	}
}
