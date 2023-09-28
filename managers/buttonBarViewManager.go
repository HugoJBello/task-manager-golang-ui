package managers

import (
	"strconv"

	"github.com/HugoJBello/task-manager-golang-ui/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ButtonBarViewManager struct {
	ApiManager ApiManager
}

func (m *ButtonBarViewManager) CreateButtonBarWithPoints(globalAppState *models.GlobalAppState) *tview.Frame {
	boardId := globalAppState.SelectedBoardId
	pointsReport, _ := m.ApiManager.GetPointsCurrentWeek(*boardId)
	lowerBarMenu := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 4, 4).
		AddText(m.PointsCurrentWeekText((*pointsReport)[0]), true, tview.AlignLeft, tcell.ColorWhite).
		AddText("_", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("_", true, tview.AlignRight, tcell.ColorWhite)
	return lowerBarMenu
}

func (m *ButtonBarViewManager) PointsCurrentWeekText(pointsReport models.PointsReport) string {
	points := pointsReport.Points
	week := pointsReport.Week
	return "WEEK: " + strconv.Itoa(week) + "  POINTS: " + strconv.Itoa(points)
}
