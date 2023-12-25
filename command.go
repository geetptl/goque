package main

type command struct {
	topic string
}

type addCommand_ struct {
	command
	message string
}

type readCommand_ struct {
	command
	number int
}

type removeCommand_ struct {
	command
}

func AddCommand() addCommand_ {
	return addCommand_{
		command: command{
			topic: "",
		},
		message: "",
	}
}

func ReadCommand() readCommand_ {
	return readCommand_{
		command: command{
			topic: "",
		},
		number: 0,
	}
}

func RemoveCommand() removeCommand_ {
	return removeCommand_{
		command: command{
			topic: "",
		},
	}
}

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
