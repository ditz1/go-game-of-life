package main

import "github.com/gen2brain/raylib-go/raylib"
import "fmt"

const WINDOW_WIDTH = 1280
const WINDOW_HEIGHT = 720

func draw_grid(g_den int) { 
	w_spacing := (WINDOW_WIDTH - 20) / g_den 
	h_spacing := (WINDOW_HEIGHT - 20) / g_den 
	for i := 0; i <= g_den; i++ { 
		//vertical lines 
		rl.DrawLine(int32(i * w_spacing) + 10, 10, int32(i * w_spacing) + 10, WINDOW_HEIGHT - 10, rl.Gray) 
		//horizontal lines 
		rl.DrawLine(10, int32(i * h_spacing) + 10, WINDOW_WIDTH - 10, int32(i * h_spacing) + 10, rl.Gray) 
	} 
} 


func fill_grid(g_den int, live_cells []rl.Vector2) {
    w_spacing := (WINDOW_WIDTH - 20) / g_den
    h_spacing := (WINDOW_HEIGHT - 20) / g_den

    for i := 0; i < g_den; i++ {
        for j := 0; j < g_den; j++ {
            x := int32(i*w_spacing) + 10
            y := int32(j*h_spacing) + 10
            width := int32(w_spacing) - 10
            height := int32(h_spacing) - 10

            cell := rl.Vector2{float32(i), float32(j)}
            if contains(live_cells, cell) {
                rl.DrawRectangle(x + 5, y + 5, width, height, rl.Magenta)
            } else {
                rl.DrawRectangle(x, y, width, height, rl.Beige)
            }
        }
    }
}

func init_conway(g_den int) []rl.Vector2 {
    live_cells := make([]rl.Vector2, 0)
    // Randomly initialize live cells
    for i := 0; i < g_den; i++ {
        for j := 0; j < g_den; j++ {
            if rl.GetRandomValue(0, 1) == 1 {
                cell := rl.Vector2{float32(i), float32(j)}
                live_cells = append(live_cells, cell)
            }
        }
    }
    return live_cells
}

func conway_loop(g_den int, live_cells []rl.Vector2) []rl.Vector2 {
    new_live_cells := make([]rl.Vector2, 0)

    for i := 0; i < g_den; i++ {
        for j := 0; j < g_den; j++ {
            cell := rl.Vector2{float32(i), float32(j)}
            live_neighbors := count_live_neighbors(g_den, live_cells, cell)

            if contains(live_cells, cell) {
                // Any live cell with fewer than two live neighbors dies, as if by underpopulation.
                // Any live cell with two or three live neighbors lives on to the next generation.
                // Any live cell with more than three live neighbors dies, as if by overpopulation.
                if live_neighbors == 2 || live_neighbors == 3 {
                    new_live_cells = append(new_live_cells, cell)
                }
            } else {
                // Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
                if live_neighbors == 3 {
                    new_live_cells = append(new_live_cells, cell)
                }
            }
        }
    }

    return new_live_cells
}

func count_live_neighbors(g_den int, live_cells []rl.Vector2, cell rl.Vector2) int {
    count := 0
    for _, neighbor := range get_neighbors(g_den, cell) {
        if contains(live_cells, neighbor) {
            count++
        }
    }
    return count
}

func get_neighbors(g_den int, cell rl.Vector2) []rl.Vector2 {
    neighbors := make([]rl.Vector2, 0)
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if i == 0 && j == 0 {
                continue
            }
            x := int(cell.X) + i
            y := int(cell.Y) + j
            if x >= 0 && x < g_den && y >= 0 && y < g_den {
                neighbor := rl.Vector2{float32(x), float32(y)}
                neighbors = append(neighbors, neighbor)
            }
        }
    }
    return neighbors
}

func contains(cells []rl.Vector2, cell rl.Vector2) bool {
    for _, c := range cells {
        if c.X == cell.X && c.Y == cell.Y {
            return true
        }
    }
    return false
}

func main() {
    rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "game of life in go")
    defer rl.CloseWindow()

    rl.SetTargetFPS(20)
    g_den := 8
    live_cells := make([]rl.Vector2, 0)
    simulation_started := false
    generation := 0

    for !rl.WindowShouldClose() {
        rl.BeginDrawing()
        rl.ClearBackground(rl.Beige)
        draw_grid(g_den)

        if !simulation_started {
            if rl.IsKeyPressed(rl.KeyOne) {
				g_den = 8
			}
			if rl.IsKeyPressed(rl.KeyTwo) {
				g_den = 16
			}
			if rl.IsKeyPressed(rl.KeyThree) {
				g_den = 32
			}
			if rl.IsKeyPressed(rl.KeySpace) {
                live_cells = init_conway(g_den)
                simulation_started = true
                generation = 0
                print("simulation started")
            }
        } else {
            fill_grid(g_den, live_cells)
            live_cells = conway_loop(g_den, live_cells)
            generation++

            if rl.IsKeyPressed(rl.KeySpace) {
                simulation_started = false
                live_cells = make([]rl.Vector2, 0)
                print("simulation stopped")
            }
        }

        if simulation_started {
            rl.DrawText(fmt.Sprintf("Generation: %d", generation), 10, 30, 25, rl.Gray)
            rl.DrawText(fmt.Sprintf("Grid Size: %dx%d", g_den, g_den), 10, 60, 18, rl.Gray)

			} else {

			rl.DrawText("press 1, 2, or 3 for small, medium, or large grids", WINDOW_WIDTH/2 - 275, WINDOW_HEIGHT/2 - 50 , 25, rl.Gray)
			rl.DrawText("press space to start/stop simulation", WINDOW_WIDTH/2 - 300, WINDOW_HEIGHT/2 , 35, rl.Gray)
            rl.DrawText(fmt.Sprintf("Selected Grid Size: %dx%d", g_den, g_den), WINDOW_WIDTH/2 - 100, WINDOW_HEIGHT/2 + 50, 22, rl.Gray)

		}


        rl.EndDrawing()
    }
}
	
