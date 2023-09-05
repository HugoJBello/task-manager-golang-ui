// Demo code for the List primitive.
package main

import (
	"fmt"
	"os"

	"github.com/HugoJBello/task-manager-golang-ui/managers"
	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	apiManager := managers.ApiManager{Url: os.Getenv("API_URL")}
	uiBoardsManager := managers.UiBoardsManager{}

	globalAppState := models.GlobalAppState{}

	updatedSelectedBoard := make(chan string)

	boards, _ := apiManager.GetBoards()
	fmt.Println(boards)

	app := tview.NewApplication()

	list, _ := uiBoardsManager.GetBoardsListUi(boards, app, &globalAppState, &updatedSelectedBoard)

	box := tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:-] [::bu]title")

		// Create the layout.
	flex := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(box, 0, 1, true)


	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	
	for v := range updatedSelectedBoard {
		fmt.Println("----->")
		fmt.Println(v)
	}

	

}
