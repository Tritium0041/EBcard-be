package main

func checkerr(err error) {
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
