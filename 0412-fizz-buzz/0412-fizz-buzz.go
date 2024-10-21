func isDivisibleBy3(n int) bool {
	return (n % 3) == 0
}

func isDivisibleBy5(n int) bool {
	return (n % 5) == 0
}

func isDivisibleBy3And5(n int) bool {
	return ((n % 3) == 0) && ((n % 5) == 0)
}

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
		if isDivisibleBy3And5(i) {
			answer[i-1] = "FizzBuzz"
		} else if isDivisibleBy3(i) {
			answer[i-1] = "Fizz"
		} else if isDivisibleBy5(i) {
			answer[i-1] = "Buzz"
		} else {
			str := fmt.Sprintf("%d", i)
			answer[i-1] = str
		}
	}
	return answer
}