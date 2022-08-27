package command

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
