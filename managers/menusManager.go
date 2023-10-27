package managers

import (
	"fmt"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenusManager struct {
	ApiManager           ApiManager
	UiTasksManager       UiTasksManager
	HistoryViewManager   HistoryViewManager
	ButtonBarViewManager ButtonBarViewManager
	ActionsViewManager   ActionsViewManager
}

func (m *MenusManager) LoadMenus(app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState, horizontalView bool) {

	tasksInBoard, _ := m.ApiManager.GetTasksInBoard(*globalAppState.SelectedBoardId)
	globalAppState.TasksInBoard = tasksInBoard

	tasksList, _ := m.UiTasksManager.GetTasksListUi(app, pages, globalAppState)

	tasksFlex := tview.NewFlex()

	if horizontalView == true {
		tasksFlex.SetDirection(tview.FlexRow)
	}

	for index, _ := range tasksList {
		selected := (*globalAppState.SelectedStatus) == tasksList[index].GetTitle()
		fmt.Println(*globalAppState.SelectedStatus)
		var primitive tview.Primitive = &tasksList[index]
		tasksFlex.AddItem(primitive, 0, 1, selected)
	}

	historicList := m.HistoryViewManager.AddHistoryPage(app, pages, globalAppState)

	pages.AddPage("historic", historicList, true, false)

	actionsList := m.ActionsViewManager.AddActionsPage(app, pages, globalAppState)

	pages.AddPage("actions", actionsList, true, false)

	pages.AddPage("tasks_board", tasksFlex, true, true)

	inputs := []tview.List{}
	inputs = append(inputs, tasksList...)

	AddCycleFocus(tasksFlex, app, inputs, globalAppState)

	lowerBarFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	lowerBarMenu := m.ButtonBarViewManager.CreateButtonBarWithPoints(globalAppState)

	lowerBarFlex.AddItem(pages, 0, 1, true)
	lowerBarFlex.AddItem(lowerBarMenu, 2, 0, false)

	if err := app.SetRoot(lowerBarFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func AddCycleFocus(flex *tview.Flex, app *tview.Application, inputs []tview.List, globalAppState *models.GlobalAppState) {
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlSpace || event.Key() == tcell.KeyCtrlK {
			CycleFocus(app, inputs, false, globalAppState)
		} else if event.Key() == tcell.KeyBacktab || event.Key() == tcell.KeyCtrlJ {
			CycleFocus(app, inputs, true, globalAppState)
		}
		return event

	})
}

func CycleFocus(app *tview.Application, elements []tview.List, reverse bool, globalAppState *models.GlobalAppState) {
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
		var primitive tview.Primitive = &elements[i]
		app.SetFocus(primitive)
		globalAppState.FocusedElement = &i

		var selected string = "doing"

		if i == 0 {
			selected = "none"
		} else {
			selected = globalAppState.Statuses[i-1]
		}

		globalAppState.SelectedStatus = &selected
		return
	}
}
