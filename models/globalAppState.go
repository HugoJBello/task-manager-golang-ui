package models

import "time"

type GlobalAppState struct {
	SelectedBoardId    *string
	SelectedTaskId     *string
	SelectedStatus     *string
	Statuses           []string
	FocusedElement     *int
	Boards             *[]Board
	TasksInBoard       *[]Task
	SelectedTask       *Task
	SelectedBoard      *Board
	RefreshApp         *chan string
	RefreshBlocked     bool
	DisplayCurrentTime string
	CurrentTime        time.Time
}

func (g *GlobalAppState) UpdateDisplayTime() {
	g.DisplayCurrentTime = g.CurrentTime.Format("Mon Jan 02 15:04:05 -0700 2006")
}

var Statuses = []string{"done", "doing", "blocked", "todo"}
