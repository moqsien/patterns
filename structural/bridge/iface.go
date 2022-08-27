package bridge

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}
