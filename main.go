package main

import (
	"fmt"
	"github/felisest/fyne_widgets/dto"
	"github/felisest/fyne_widgets/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne widgets")
	w.Resize(fyne.NewSize(1024, 400))

	cont := container.NewVBox()

	table := widgets.NewDataTable[dto.Model]()
	table.Fill(rows)

	cont.Add(table)

	pagin := widgets.NewPagintator(20, 1200,
		func(limit, total int) {
			fmt.Println("Click left")
		},
		func(limit, total int) {
			fmt.Println("Click right")
		},
	)

	cont.Add(pagin)

	w.SetContent(cont)
	w.ShowAndRun()
}

var rows = []dto.Model{
	{
		Id:   1,
		Name: "Александр Пушкин",
		Body: "Я помню чудное мгновенье: Передо мной явилась ты, Как мимолетное виденье, Как гений чистой красоты...",
	},
	{
		Id:   2,
		Name: "John Keats",
		Body: "A thing of beauty is a joy for ever: Its loveliness increases; it will never Pass into nothingness...",
	},
	{
		Id:   3,
		Name: "William Blake",
		Body: "Tyger Tyger, burning bright, In the forests of the night; What immortal hand or eye, Could frame thy fearful symmetry?",
	},
	{
		Id:   4,
		Name: "Dylan Thomas",
		Body: "Do not go gentle into that good night, Old age should burn and rave at close of day; Rage, rage against the dying of the light.",
	},
	{
		Id:   5,
		Name: "William Shakespeare",
		Body: "Shall I compare thee to a summer's day? Thou art more lovely and more temperate...",
	},
}
