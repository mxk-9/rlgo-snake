package gui

import (
	"errors"
)

var (
	ContainerIsFull         = errors.New("ContainerIsFull")
	ContainerCellIsNotEmpty = errors.New("ContainerCellIsNotEmpty")
	PositionOutOfBounds     = errors.New("InsertPosIsOutOfBounds")
	RowsOrColsLessThanZero  = errors.New("RowsOrColsLessThanZero")
)
