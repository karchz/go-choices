package choices

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

var selectRow = 1
var itemPerPage = 10
var currentPage = 0
var allItems = []string{}
var pageItems = []string{}
var header = "Choose Item ."
var page int

func Ask(items []string, title string) (int, bool) {

	var result int
	answered := false
	header = title
	allItems = items

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.HideCursor()

	refreshPage()
	setHighlight(1)
	termbox.Flush()

end:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			refreshPage()
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				selectRow = 0
				result = -1
				break end
			case termbox.KeyEnter:
				answered = true
				result = (itemPerPage * currentPage) + (selectRow - 1)
				break end
			case termbox.KeyArrowLeft:
				if currentPage != 0 {
					currentPage--
				}
				refreshPage()
			case termbox.KeyArrowRight:
				if page > currentPage+1 {
					currentPage++
				}
				refreshPage()
			case termbox.KeyArrowDown:
				if selectRow < len(pageItems) {
					selectRow++
				} else {
					selectRow = 1
				}
				setHighlight(selectRow)
			case termbox.KeyArrowUp:
				if selectRow == 1 {
					selectRow = len(pageItems)
				} else {
					selectRow--
				}
				setHighlight(selectRow)
			default:
				break
			}
		default:
			break
		}
	}

	return result, answered
}

func refreshPage() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	_, height := termbox.Size()

	itemPerPage = height - 2

	page = len(allItems) / itemPerPage
	if len(allItems)%itemPerPage != 0 {
		page++
	}

	start := currentPage * itemPerPage
	end := start + itemPerPage
	if len(allItems[start:]) < itemPerPage {
		end = start + len(allItems[start:])
	}

	pageItems = allItems[start:end]

	for index, item := range pageItems {
		writeLine(index+1, item)
	}
	resetHeader()
	resetFooter()
	setHighlight(1)
	termbox.Flush()
}

func resetHeader() {
	centerWrite(0, header, termbox.ColorMagenta, termbox.ColorCyan)
}

func resetFooter() {
	_, bottom := termbox.Size()
	page := len(allItems) / itemPerPage
	if len(allItems)%itemPerPage != 0 {
		page++
	}
	prevText := "Prev <- "
	if currentPage == 0 {
		prevText = "        "
	}
	nextText := " -> Next"
	if currentPage+1 == page {
		nextText = "        "
	}
	footer := prevText + strconv.Itoa(currentPage+1) + "/" + strconv.Itoa(page) + nextText
	// termbox.co
	centerWrite(bottom-1, footer, termbox.ColorBlack, termbox.ColorCyan)
}

func centerWrite(row int, text string, textColor termbox.Attribute, bgColor termbox.Attribute) {
	width, _ := termbox.Size()
	fixing := (width / 2) - (len(text) / 2)
	for col := 0; col < width; col++ {
		char := ' '
		if col >= fixing && col < fixing+(len(text)) {
			char = rune(text[col-fixing])
		}
		termbox.SetCell(col, row, char, textColor, bgColor)
	}
}
func writeLine(row int, str string) {

	width, _ := termbox.Size()
	for col := 0; col < width; col++ {
		char := ' '
		if col < len(str) {
			char = rune(str[col])
		}
		termbox.SetCell(col, row, char, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func setHighlight(targetRow int) {
	width, _ := termbox.Size()
	selectRow = targetRow
	for row := 1; row <= itemPerPage; row++ {
		bgColor := termbox.ColorDefault
		if row == targetRow {
			bgColor = termbox.ColorWhite
		}
		for col := 0; col < width; col++ {
			char := termbox.CellBuffer()[(width*row)+col].Ch
			termbox.SetCell(col, row, char, termbox.ColorDefault, bgColor)
		}
	}
	termbox.Flush()
}
