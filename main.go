package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "math/rand"

// import "time"
// import "fmt"

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	//rand.Seed(time.Now().UnixNano())
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera := rl.NewCamera2D(
		rl.NewVector2(0.0, 0.0),
		rl.NewVector2(0.0, 0.0),
		0,
		1.0,
	)

	state := map[rl.Vector2]bool{}

	for i := 0; i < 5000; i++ {
		state[rl.NewVector2(float32(randInt(0, 800/4)), float32(randInt(0, 450/4)))] = true
	}

	// iter := 0
	dt := 0.0

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		for pos := range state {
			rl.DrawRectangle(int32(pos.X)*4, int32(pos.Y)*4, 4, 4, rl.Red)
		}
		rl.EndMode2D()
		rl.DrawFPS(20, 20)
		// rl.DrawText(fmt.Sprintf("FPS: %v", rl.GetFPS()), 20, 20, 18, rl.Blue)
		rl.EndDrawing()

		if dt > 0.1 {
			checked := map[rl.Vector2]int8{}
			newState := map[rl.Vector2]bool{}

			for pos := range state {
				checked[rl.NewVector2(pos.X+1, pos.Y)]++
				checked[rl.NewVector2(pos.X-1, pos.Y)]++
				checked[rl.NewVector2(pos.X, pos.Y+1)]++
				checked[rl.NewVector2(pos.X, pos.Y-1)]++
				checked[rl.NewVector2(pos.X-1, pos.Y-1)]++
				checked[rl.NewVector2(pos.X+1, pos.Y-1)]++
				checked[rl.NewVector2(pos.X-1, pos.Y+1)]++
				checked[rl.NewVector2(pos.X+1, pos.Y+1)]++
			}

			for pos, count := range checked {
				if state[rl.NewVector2(pos.X, pos.Y)] { // Alive
					if count > 1 && count < 4 {
						newState[rl.NewVector2(pos.X, pos.Y)] = true
					}
				} else { // Dead
					if count == 3 {
						newState[rl.NewVector2(pos.X, pos.Y)] = true
					}
				}
			}

			state = newState

			dt = 0.0
		} else {
			dt += float64(rl.GetFrameTime())
		}

		/*
			if len(state) > 0 && false {
				fmt.Printf("state: %v\n", state)
				fmt.Printf("checked: %v\n", checked)
				fmt.Printf("iter: %d\n", iter)
				iter++
			}*/

	}
}
