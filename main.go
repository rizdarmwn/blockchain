package main

func main() {
	bc := NewBlockchain("0xRizdar")
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
