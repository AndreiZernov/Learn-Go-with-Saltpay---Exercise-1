package error_handler

func HandlePanic(err error) {
	if err != nil {
		panic(err)
	}
}
