/*
Use it like this:

func main() {
	for i := range GeneratePrimeNumbers() {
		fmt.Println(i)
	}
}
*/

package primeutil

func GeneratePrimeNumbers() chan int {
	channel := make(chan int)

	go func() {
		for i := 2; ; i++ {
			if IsPrime(i) {
				channel <- i
			}
		}
	}()

	return channel
}
