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

func (m *MenusManager) LoadMenus(listBoards *tview.List, app *tview.Application, pages *tview.Pages, updatedSelectedBoard *chan string, globalAppState *models.GlobalAppState, horizontalView bool) {

	tasksInBoard, _ := m.ApiManager.GetTasksInBoard(*globalAppState.SelectedBoardId)
	globalAppState.TasksInBoard = tasksInBoard

	tasksList, _ := m.UiTasksManager.GetTasksListUi(app, updatedSelectedBoard, globalAppState)

	selectedSideMenu := (*globalAppState.SelectedStatus) == "none"
	tasksWithSideMenuflex := tview.NewFlex().
		AddItem(listBoards, 18, 0, selectedSideMenu)

	tasksFlex := tview.NewFlex()

	if horizontalView == true {
		tasksFlex.SetDirection(tview.FlexRow)
	}

	for index, _ := range tasksList {
		selected := (*globalAppState.SelectedStatus) == globalAppState.Statuses[index]
		fmt.Println(selected, globalAppState.Statuses[index], *globalAppState.SelectedStatus)
		tasksFlex.AddItem(tasksList[index], 0, 1, selected)
	}

	tasksWithSideMenuflex.AddItem(tasksFlex, 0, 3, true)

	historicList := m.HistoryViewManager.AddHistoryPage(app, pages, globalAppState)

	pages.AddPage("historic", historicList, true, false)

	actionsList := m.ActionsViewManager.AddActionsPage(app, pages, updatedSelectedBoard, globalAppState)

	pages.AddPage("actions", actionsList, true, false)

	pages.AddPage("tasks_board", tasksWithSideMenuflex, true, true)

	inputs := []tview.Primitive{
		listBoards,
	}
	inputs = append(inputs, tasksList...)

	AddCycleFocus(tasksWithSideMenuflex, app, inputs, globalAppState)

	lowerBarFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	lowerBarMenu := m.ButtonBarViewManager.CreateButtonBarWithPoints(globalAppState)

	lowerBarFlex.AddItem(pages, 0, 1, true)
	lowerBarFlex.AddItem(lowerBarMenu, 2, 0, false)

	if err := app.SetRoot(lowerBarFlex, true).EnableMouse(true).Run(); err != nil {
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
