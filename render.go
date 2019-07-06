// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"
	"log"
	"os"
	"sync"

	tb "github.com/hawklithm/termbox-go"
)

type Drawable interface {
	GetRect() image.Rectangle
	SetRect(int, int, int, int)
	Draw(*Buffer)
	sync.Locker
}

func Render(items ...Drawable) {
	fileName := "log.txt"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	wxLogger := log.New(logFile, "[*]", log.LstdFlags)
	defer logFile.Close()
	for _, item := range items {
		buf := NewBuffer(item.GetRect())
		item.Lock()
		item.Draw(buf)
		item.Unlock()
		for point, cell := range buf.CellMap {
			if point.In(buf.Rectangle) {
				if os.Getenv("WECHAT_TERM") == "iterm" {
					if cell.Bytes != nil {
						tb.SetImageCell(point.X, point.Y, cell.Bytes)
						wxLogger.Println("image point", point)
						continue
					}
				}
				tb.SetCell(
					point.X, point.Y,
					cell.Rune,
					tb.Attribute(cell.Style.Fg+1)|tb.Attribute(cell.Style.Modifier), tb.Attribute(cell.Style.Bg+1),
				)
			}
		}
	}
	tb.Flush()
}
