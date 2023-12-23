package main

type context struct {
	topic            string
	message          string
	numberOfMessages int
}

func NewContext() context {
	return context{
		topic:            "",
		message:          "",
		numberOfMessages: 0,
	}
}
