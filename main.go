// Demo code for the List primitive.
package main

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/HugoJBello/task-manager-golang-ui/managers"
	"github.com/subosito/gotenv"
	"os"
)

func init() {
	gotenv.Load()
}

func main() {
	apiManager := managers.ApiManager{Url:os.Getenv("API_URL")}
	result, _ := apiManager.GetBoards()
	fmt.Println(result)

	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}