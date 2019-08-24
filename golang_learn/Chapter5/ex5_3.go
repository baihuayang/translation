package main


func main(){

	for i := 1; i <= 100; i++ {
		switch {
			case i%3 ==0 && i%5 != 0:
				println("Fizz")
			case i%3 != 0 && i%5 == 0:
				println("Buzz")
			case i%15 == 0:
				println("FizzBuzz")
			default:
				println(i)
		}
		
	}
}