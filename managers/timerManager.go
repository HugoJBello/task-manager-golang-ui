package managers

import (
	"fmt"
	"time"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/rivo/tview"
)

type TimerManager struct {
}

func (m *TimerManager) SetTimer(app *tview.Application, globalAppState *models.GlobalAppState) {

	for now := range time.Tick(5 * 60 * time.Second) {
		if globalAppState.RefreshBlocked == false {
			fmt.Println("refresh started")
			app.Stop()
			*&globalAppState.CurrentTime = now
			globalAppState.UpdateDisplayTime()
			*globalAppState.RefreshApp <- "refresh"
		}
	}

}
