package models

type GlobalAppState struct {
	SelectedBoardId *string
	SelectedTaskId  *string
	SelectedStatus  *string
	Statuses        *[]string
	FocusedElement  *int
	Boards          *[]Board
	TasksInBoard    *[]Task
	SelectedBoard   *Board
}
