package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

type ClipboardLabel struct {
	widget.BaseWidget
	label *widget.Label
	btn   *widget.Button

	text string
}

func NewClipboardLabel(text string) *ClipboardLabel {
	btn := widget.NewButton("📋", func() {
		clipboard.Write(clipboard.FmtText, []byte(text))
	},
	)

	clip_label := &ClipboardLabel{label: widget.NewLabel(text), btn: btn, text: text}
	clip_label.label.Truncation = fyne.TextTruncateEllipsis

	clip_label.ExtendBaseWidget(clip_label)

	return clip_label
}

func (l *ClipboardLabel) CreateRenderer() fyne.WidgetRenderer {
	cont := container.NewBorder(nil, nil, nil, l.btn, l.label)

	return widget.NewSimpleRenderer(cont)
}
