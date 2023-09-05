package models

type GlobalAppState struct {
	SelectedBoardId *string
	SelectedTaskId  *string
	Boards          *[]Board
	TasksInBoard    *[]Task
	SelectedBoard   *Board
}
