package main
import (
	"fmt"
	"time"
)
// Regular function
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(500 * time.Millisecond)
 	}
}
func printLetters() {
	letters := []string{"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		fmt.Printf("Letter: %s\n", letter)
		time.Sleep(500 * time.Millisecond)
 	}
}

func counter(name string, max int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("Counter %s: %d\n", name, i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("Counter %s done!\n", name)
}

func fetchWebsite(name string, delayMs int) {
	fmt.Printf("Fetching %s...\n", name)
	time.Sleep(time.Duration(delayMs) * time.Millisecond)
	fmt.Printf("✓ Got data from %s\n", name)
}

func calculateSum(numbers []int, result chan int) {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
		time.Sleep(100 * time.Millisecond)
	}
	result <- sum
	close(result)
}

func calculateAverage(numbers []int, result chan float64) {
	average := 0.0
	for i := range numbers {
		average += float64(numbers[i])
		time.Sleep(100 * time.Millisecond)
	}
	average = average / float64(len(numbers))
	result <- average
	close(result)
}

func sender(name string, messages chan string, count int) {
	for i := range count {
		message := fmt.Sprintf("Message %d from %s", i+1 , name)
		messages <- message
		time.Sleep(150 * time.Millisecond)
	}
}



func main() {
 	fmt.Println("=== Message Queue ===")

	// Create a channel for messages
	messageChannel := make(chan string, 10)

	// Start senders
	go sender("Alice", messageChannel, 3)
	go sender("Bob", messageChannel, 2)
	go sender("Charlie", messageChannel, 4)

	
	// Receive messages
	for i := range 9 {
		message := <-messageChannel
		fmt.Printf("Message %d: %s\n", i+1, message)
	}
	
	fmt.Println("\nAll messages received!")

}