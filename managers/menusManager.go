package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenusManager struct {
	ApiManager     ApiManager
	UiTasksManager UiTasksManager
}

func (m *MenusManager) LoadMenus(listBoards *tview.List, app *tview.Application, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState) {

	tasksInBoard, _ := m.ApiManager.GetTasksInBoard(*globalAppState.SelectedBoardId)
	globalAppState.TasksInBoard = tasksInBoard

	tasksList, _ := m.UiTasksManager.GetTasksListUi(app, updatedSelectedBoard, globalAppState)

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(listBoards, 0, 1, true)

	for index, _ := range tasksList {
		flex.AddItem(tasksList[index], 0, 1, true)
	}

	inputs := []tview.Primitive{
		listBoards,
	}
	inputs = append(inputs, tasksList...)

	AddCycleFocus(flex, app, inputs, globalAppState)

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

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
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
