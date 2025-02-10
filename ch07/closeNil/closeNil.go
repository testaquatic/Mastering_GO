package main

// nil 채널을 닫으려고 하므로 패닉이 발생한다.
func main() {
	var c chan string

	close(c)
}
