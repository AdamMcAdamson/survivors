package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Wall struct {
	Id       int
	Rect     rl.Rectangle
	Origin   rl.Vector2
	Rotation float32
}

func (w Wall) draw() {
	rl.DrawRectanglePro(w.Rect, w.Origin, w.Rotation, rl.Magenta)
	// var sourceRect = rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 400}
	// var destRect = rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: 80, Height: 80}
	// var angle = -rl.Vector2LineAngle(p.FaceDir, rl.Vector2{X: -1, Y: 0}) * 360 / rl.Pi
	// rl.DrawTexturePro(p.Texture, sourceRect, destRect, rl.Vector2{X: 25, Y: 40}, angle, rl.White)
}

func placeWall(posX float32, posY float32, width float32, height float32, rotation float32) {
	entities.Walls = append(entities.Walls, &Wall{Id: entities.NextId, Rect: rl.NewRectangle(posX, posY, width, height), Origin: rl.Vector2{X: width / 2, Y: height / 2}, Rotation: rotation})
	entities.NextId++
}
