package dto

type Model struct {
	Id   int    `title:"ID" type:"label"`
	Name string `title:"Name" type:"label"`
	Body string `title:"Body" type:"clipboard_label"`
}
