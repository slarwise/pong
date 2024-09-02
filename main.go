package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BALL_SPEED_INCREASE = 0
	BALL_SPEED          = 8
	BALL_RADIUS         = 8

	PADDLE_SPEED  = 5
	PADDLE_WIDTH  = 10
	PADDLE_HEIGHT = 100

	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

var (
	BACKGROUND_COLOR = rl.NewColor(50, 50, 50, 255)
	PADDLE_COLOR     = rl.White
	BALL_COLOR       = rl.White
	GAME_OVER_COLOR  = rl.Green
)

type Ball struct{ x, y, dx, dy int32 }
type SIDE int32

const (
	LEFT_SIDE SIDE = iota
	RIGHT_SIDE
)

type Player struct {
	side SIDE
	y    int32
}

func main() {
	ball := Ball{x: 10 + BALL_RADIUS, y: 10 + BALL_RADIUS, dx: BALL_SPEED, dy: BALL_SPEED}
	playerLeft := Player{side: LEFT_SIDE, y: 100}
	playerRight := Player{side: RIGHT_SIDE, y: 100}
	result := CONTINUE

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Pong")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && result == CONTINUE {
		rl.BeginDrawing()

		rl.ClearBackground(BACKGROUND_COLOR)

		if rl.IsKeyDown(rl.KeyD) {
			movePaddleDown(&playerLeft)
		} else if rl.IsKeyDown(rl.KeyA) {
			movePaddleUp(&playerLeft)
		}
		if rl.IsKeyDown(rl.KeyJ) || rl.IsKeyDown(rl.KeyDown) {
			movePaddleDown(&playerRight)
		} else if rl.IsKeyDown(rl.KeyK) || rl.IsKeyDown(rl.KeyUp) {
			movePaddleUp(&playerRight)
		}
		result = moveBall(&ball, playerLeft, playerRight)

		drawPlayerLeft(playerLeft)
		drawPlayerRight(playerRight)
		drawBall(ball)

		rl.EndDrawing()
	}

	if result != CONTINUE {
		rl.BeginDrawing()
		rl.ClearBackground(BACKGROUND_COLOR)
		if result == LEFT_PLAYER_WON {
			rl.DrawText("Left player won", WINDOW_WIDTH/2-150, WINDOW_HEIGHT/2-40, 40, GAME_OVER_COLOR)
		} else if result == RIGHT_PLAYER_WON {
			rl.DrawText("Right player won", WINDOW_WIDTH/2-150, WINDOW_HEIGHT/2-40, 40, GAME_OVER_COLOR)
		}
		rl.EndDrawing()
		time.Sleep(1 * time.Second)
	}

	rl.CloseWindow()
}

func movePaddleDown(p *Player) {
	ny := p.y + PADDLE_SPEED
	if ny < WINDOW_HEIGHT-PADDLE_HEIGHT {
		p.y = ny
	}
}

func movePaddleUp(paddle *Player) {
	ny := paddle.y - PADDLE_SPEED
	if ny >= 0 {
		paddle.y = ny
	}
}

type RESULT int

const (
	CONTINUE RESULT = iota
	LEFT_PLAYER_WON
	RIGHT_PLAYER_WON
)

func moveBall(b *Ball, pLeft, pRight Player) RESULT {
	increaseSpeed := false
	result := CONTINUE
	switch {
	case b.x-BALL_RADIUS <= 0:
		b.dx *= -1
		increaseSpeed = true
		result = RIGHT_PLAYER_WON
	case b.x-BALL_RADIUS <= PADDLE_WIDTH-1:
		if b.y+BALL_RADIUS >= pLeft.y && b.y-BALL_RADIUS <= pLeft.y+PADDLE_HEIGHT {
			b.dx *= -1
			increaseSpeed = true
		}
	case b.x+BALL_RADIUS >= WINDOW_WIDTH-1:
		result = LEFT_PLAYER_WON
	case b.x+BALL_RADIUS >= WINDOW_WIDTH-PADDLE_WIDTH-1:
		if b.y+BALL_RADIUS >= pRight.y && b.y-BALL_RADIUS <= pRight.y+PADDLE_HEIGHT {
			b.dx *= -1
			increaseSpeed = true
		}
	}
	if b.y-BALL_RADIUS <= 0 || b.y+BALL_RADIUS >= WINDOW_HEIGHT-1 {
		b.dy *= -1
		increaseSpeed = true
	}
	if increaseSpeed {
		if b.dx > 0 {
			b.dx += BALL_SPEED_INCREASE
		}
	}
	b.x += b.dx
	b.y += b.dy
	return result
}

func drawBall(b Ball) {
	rl.DrawCircle(b.x, b.y, BALL_RADIUS, BALL_COLOR)
}

func drawPlayerLeft(p Player) {
	rl.DrawRectangle(0, p.y, PADDLE_WIDTH, PADDLE_HEIGHT, PADDLE_COLOR)
}

func drawPlayerRight(p Player) {
	rl.DrawRectangle(WINDOW_WIDTH-PADDLE_WIDTH-1, p.y, PADDLE_WIDTH, PADDLE_HEIGHT, PADDLE_COLOR)
}
