package gui

import (
	"embed"

	"fmt"
	"rltest/pkg/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Container struct {
	BaseWidget

	colStep float32 // scale_x/cols
	rowStep float32 // scale_y/rows

	rows int
	cols int

	items [][]Widget
}

func NewConrainer(
	pos, scale rl.Vector2, rows, cols int,
) (c *Container, err error) {
	if rows < 0 || cols < 0 {
		err = RowsOrColsLessThanZero
		return
	}

	// c = &Container{
	// 	Pos:   pos,
	// 	scale: scale,
	// 	size: rl.Vector2{
	// 		X: pos.X + scale.X,
	// 		Y: pos.Y + scale.Y,
	// 	},
	// 	rows:          rows,
	// 	cols:          cols,
	// 	showWidget: true,
	// }
	c = &Container{}
	c.Pos = pos
	c.Scale = scale

	c.colStep = c.Scale.X / float32(cols)
	c.rowStep = c.Scale.Y / float32(rows)

	c.items = make([][]Widget, cols)

	for i := range cols {
		c.items[i] = make([]Widget, rows)
	}
	return
}

func (c *Container) SetTexture(
	content *embed.FS, imagePath string, frames int32, rotateDegree int32,
) (err error) {
	if frames <= 0 {
		rl.SetTraceLogLevel(rl.LogError)
		rl.TraceLog(rl.LogError,
			fmt.Sprintf("Frame count cannot be less than 1! Got: '%d'", frames),
		)
		err = fmt.Errorf("FramesLessThanOne")
		return
	}

	image, err := assets.LoadImage(content, imagePath)
	if err != nil {
		rl.SetTraceLogLevel(rl.LogError)
		errMsg := fmt.Sprintf("Failed to create button: %v", err)
		rl.TraceLog(rl.LogError, errMsg)
	}
	defer rl.UnloadImage(image)

	c.SpriteFrames = frames
	rl.ImageResizeNN(image, int32(c.Scale.X), int32(c.Scale.Y)*frames)
	c.Texture = rl.LoadTextureFromImage(image)
	c.FrameHeight = float32(c.Texture.Height) / float32(frames)
	c.DrawRec = rl.NewRectangle(0, 0, float32(c.Texture.Width), c.FrameHeight)

	return
}

func (c *Container) replaceItemWithConfirm(
	w Widget, row, col int, replace bool,
) (err error) {
	if row < 0 || col < 0 || row >= c.rows || col >= c.cols {
		err = PositionOutOfBounds
		fmt.Printf(
			"%v:\nMax: %dx%d\nGot: %dx%d\n",
			err, c.rows-1, c.cols-1, row, col,
		)
		return
	}

	if !replace && c.items[row][col] != nil {
		err = ContainerCellIsNotEmpty
		return
	}

	newPos := rl.Vector2{
		X: c.Pos.X + c.colStep*float32(col),
		Y: c.Pos.Y + c.rowStep*float32(row),
	}

	newSize := rl.Vector2{
		X: c.Pos.X + c.colStep*float32(col+1),
		Y: c.Pos.Y + c.rowStep*float32(row+1),
	}

	w.SetPos(newPos)
	w.SetSize(newSize)

	c.items[row][col] = w
	return
}

func (c *Container) InsertItem(w Widget, row, col int) (err error) {
	err = c.replaceItemWithConfirm(w, row, col, false)
	return
}

func (c *Container) ReplaceItem(w Widget, row, col int) (err error) {
	err = c.replaceItemWithConfirm(w, row, col, true)
	return
}

func (c *Container) Draw() {
	if c.ShowWidget {
		rl.DrawTextureRec(c.Texture, c.DrawRec, c.Pos, rl.White)
		for _, row := range c.items {
			for _, item := range row {
				if item != nil {
					item.Draw()
				}
			}
		}
	}
}
