package main

const port = ":8040"

func main() {
	srv := NewServer()

	srv.Run(port)
}
