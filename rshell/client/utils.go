package main

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
