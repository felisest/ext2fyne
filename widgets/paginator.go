package widgets

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const defaultLimit int = 20

type Paginator struct {
	widget.BaseWidget

	leftTap  func(int, int)
	rightTap func(int, int)

	limitStr string
	infoStr  string
	totalStr string
	total    int
}

func NewPagintator(limit, total int, left, right func(limit, total int)) *Paginator {
	pag := &Paginator{total: total, limitStr: strconv.Itoa(limit), leftTap: left, rightTap: right}
	pag.setTotal(total)

	pag.ExtendBaseWidget(pag)

	return pag
}

func (p *Paginator) CreateRenderer() fyne.WidgetRenderer {
	lBtn := widget.NewButton("◀",
		func() {
			p.leftTap(p.getLimit(), p.total)
		},
	)

	limEntry := NewNumericalEntryWithData(binding.BindString(&p.limitStr))

	rBtn := widget.NewButton("▶",
		func() {
			p.rightTap(p.getLimit(), p.total)
		},
	)

	infoLabel := widget.NewLabelWithData(binding.BindString(&p.infoStr))
	totalLabel := widget.NewLabelWithData(binding.BindString(&p.totalStr))

	cont := container.New(layout.NewHBoxLayout(), lBtn, limEntry, rBtn, layout.NewSpacer(), infoLabel, totalLabel)

	return widget.NewSimpleRenderer(cont)
}

func (p *Paginator) SetInfo(info string) {
	p.infoStr = info
}

func (p *Paginator) getLimit() int {
	lim, err := strconv.Atoi(p.limitStr)
	if err != nil {
		return defaultLimit
	}

	return lim
}

func (p *Paginator) setTotal(total int) {
	p.totalStr = fmt.Sprintf("Total: %d", total)
}
