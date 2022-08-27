package command

/*
命令是一种行为设计模式，它可将请求或简单操作转换为一个对象。
此类转换让你能够延迟进行或远程执行请求，还可将其放入队列中。

创建独立命令对象的优势在于可将UI逻辑与底层业务逻辑解耦。
这样就无需为每个请求者开发不同的处理者了。
命令对象中包含执行所需的全部信息，所以也可用于延迟执行。
*/

func execute() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
