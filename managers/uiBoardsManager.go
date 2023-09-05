package managers

import (
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type UiBoardsManager struct {
}

func (m *UiBoardsManager) GetBoardsListUi(boards *[]models.Board, app *tview.Application, globalAppState *models.GlobalAppState, updatedSelectedBoard *chan string) (*tview.List, error) {
	list := tview.NewList()

	for index, _ := range *boards {
		br := (*boards)[index]
		list.AddItem(br.BoardTitle, br.BoardBody, GetRune(index), func() {
			go func() {
				*updatedSelectedBoard <- br.BoardId
			}()
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		go func() {
			*updatedSelectedBoard <- "none"
			close(*updatedSelectedBoard)
		}()

		app.Stop()
	})

	return list, nil
}

func GetRune(index int) rune {
	return rune('a' + index)
}