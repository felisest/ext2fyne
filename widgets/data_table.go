package widgets

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DataTable[T any] struct {
	widget.BaseWidget

	container *fyne.Container
}

func NewDataTable[T any]() *DataTable[T] {
	table := &DataTable[T]{container: container.NewVBox()}
	table.ExtendBaseWidget(table)

	return table
}

func (t *DataTable[T]) CreateRenderer() fyne.WidgetRenderer {

	return widget.NewSimpleRenderer(t.container)
}

func (t *DataTable[T]) Fill(rows []T) {
	t.container.RemoveAll()

	scroll := container.NewVScroll(container.NewVBox())
	t.container.Add(scroll)

	if len(rows) <= 0 {
		return
	}

	titles := getTitles(rows[0])
	t.container.Add(header(titles))

	inverted := invertColor(theme.Color(theme.ColorNameBackground))

	for _, row := range rows {
		t.container.Add(container.NewBorder(canvas.NewRectangle(inverted), nil, nil, nil, nil))
		t.container.Add(newRow(row))
	}

	t.container.Add(container.NewBorder(canvas.NewRectangle(inverted), nil, nil, nil, nil))
}

func newRow[T any](row T) *fyne.Container {
	objects := getFields(row)
	cont := container.NewGridWithColumns(len(objects))

	for _, object := range objects {
		cont.Add(object)
	}

	return cont
}

func getFields(s any) []fyne.CanvasObject {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return []fyne.CanvasObject{}
	}

	typ := val.Type()
	objects := []fyne.CanvasObject{}

	for i := range val.NumField() {
		value := val.Field(i)

		title_tag := typ.Field(i).Tag.Get("title")
		if title_tag == "-" {
			continue
		}

		type_tag := typ.Field(i).Tag.Get("type")

		objects = append(objects, fabric(type_tag, fmt.Sprint((value.Interface()))))
	}

	return objects
}

func fabric(typ string, text string) fyne.CanvasObject {
	switch typ {
	case "clipboard_label":
		return NewClipboardLabel(text)
	default:
		l := widget.NewLabel(text)
		l.Truncation = fyne.TextTruncateEllipsis
		return l
	}
}

func header(cols []string) *fyne.Container {
	cont := container.NewGridWithColumns(len(cols))
	for _, col := range cols {
		cont.Add(widget.NewLabel(col))
	}

	return cont
}

func getTitles(row any) []string {
	val := reflect.ValueOf(row)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	names := []string{}
	typ := val.Type()

	for i := range val.NumField() {
		field := typ.Field(i)

		fieldName := field.Name
		tag := field.Tag.Get("title")

		switch tag {
		case "-":
		case "":
			names = append(names, fieldName)
		default:
			names = append(names, tag)
		}
	}

	return names
}
