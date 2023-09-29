package main

const port = ":9940"

func main() {
	srv := NewServer()

	srv.Run(port)
}
