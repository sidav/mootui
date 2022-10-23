package tui_client

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"time"
)

type consoleIO struct {
	offsetX, offsetY int
	style            tcell.Style
	screen           tcell.Screen
}

func (c *consoleIO) getConsoleSize() (int, int) {
	return c.screen.Size()
}

func (c *consoleIO) clearScreen() {
	io.screen.Clear()
	cw, ch := c.getConsoleSize()
	io.setStyle(tcell.ColorBlack, tcell.ColorBlack)
	io.drawFilledRect(' ', 0, 0, cw, ch)
}

func (c *consoleIO) debugPrint(str string, args ...interface{}) {
	io.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	c.putString(fmt.Sprintf(str, args...), 0, 0)
	io.screen.Show()
	time.Sleep(100 * time.Millisecond)
}

func (c *consoleIO) readKey() string {
	for {
		ev := c.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return "EXIT"
			}
			return eventToKeyString(ev)
		}
	}
}

func (c *consoleIO) setOffsets(x, y int) {
	c.offsetX = x
	c.offsetY = y
}

func eventToKeyString(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyUp:
		return "UP"
	case tcell.KeyRight:
		return "RIGHT"
	case tcell.KeyDown:
		return "DOWN"
	case tcell.KeyLeft:
		return "LEFT"
	case tcell.KeyEscape:
		return "ESCAPE"
	case tcell.KeyEnter:
		return "ENTER"
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		return "BACKSPACE"
	case tcell.KeyTab:
		return "TAB"
	case tcell.KeyDelete:
		return "DELETE"
	case tcell.KeyInsert:
		return "INSERT"
	case tcell.KeyEnd:
		return "END"
	case tcell.KeyHome:
		return "HOME"
	default:
		return string(ev.Rune())
	}
}

func (c *consoleIO) putChar(chr rune, x, y int) {
	c.screen.SetCell(x+c.offsetX, y+c.offsetY, c.style, chr)
}

func (c *consoleIO) putString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i+c.offsetX, y+c.offsetY, c.style, rune(str[i]))
	}
}

func (c *consoleIO) setStyle(fg, bg tcell.Color) {
	c.style = c.style.Background(bg).Foreground(fg)
}

func (c *consoleIO) resetStyle() {
	c.setStyle(tcell.ColorWhite, tcell.ColorBlack)
}

func (c *consoleIO) drawFilledRect(char rune, fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		for y := fy; y <= fy+h; y++ {
			c.putChar(char, x, y)
		}
	}
}

func (c *consoleIO) drawRect(fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		c.putChar(' ', x, fy)
		c.putChar(' ', x, fy+h)
	}
	for y := fy; y <= fy+h; y++ {
		c.putChar(' ', fx, y)
		c.putChar(' ', fx+w, y)
	}
}

func (c *consoleIO) drawStringCenteredAround(s string, x, y int) {
	c.putString(s, x-len(s)/2, y)
}
