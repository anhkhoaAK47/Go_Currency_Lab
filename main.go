package main

import (
	"fmt"
	"math/rand"
	"sync"
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

// Task 1.1
func counter(name string, max int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("Counter %s: %d\n", name, i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("Counter %s done!\n", name)
}

// Task 1.2
func fetchWebsite(name string, delayMs int) {
	fmt.Printf("Fetching %s...\n", name)
	time.Sleep(time.Duration(delayMs) * time.Millisecond)
	fmt.Printf("✓ Got data from %s\n", name)
}

// Task 2.1
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

// Task 2.2
func sender(name string, messages chan string, count int) {
	for i := 0; i < count; i++ {
		message := fmt.Sprintf("Message %d from %s", i+1, name)
		messages <- message
		time.Sleep(150 * time.Millisecond)
	}
}

// Task 3.1
func downloadFile(filename string, sizeMB int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading %s (%dMB)...\n", filename, sizeMB)
	time.Sleep(time.Duration(sizeMB*100) * time.Millisecond)
	fmt.Printf("✓ %s complete!\n", filename)
}

// Task 3.2
func findEvens(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
		time.Sleep(50 * time.Millisecond)
	}
	result <- evens
}

func findOdds(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var odds []int
	for _, num := range numbers {
		if num%2 != 0 {
			odds = append(odds, num)
		}
		time.Sleep(50 * time.Millisecond)
	}
	result <- odds
}

func findSquares(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var squares []int
	for _, num := range numbers {
		squares = append(squares, num*num)
		time.Sleep(50 * time.Millisecond)
	}
	result <- squares
}

// Task 4.1
func searchEngineA(query string, ch chan string) {
	time.Sleep(300 * time.Millisecond)
	result := fmt.Sprintf("Results from Engine A for '%s'", query)
	ch <- result
}

func searchEngineB(query string, ch chan string) {
	time.Sleep(200 * time.Millisecond)
	result := fmt.Sprintf("Results from Engine B for '%s'", query)
	ch <- result
}

// Task 5.1
type Student struct {
	ID         int
	StudyHours int
}

func study(student Student, library chan bool, wg *sync.WaitGroup, waitChan chan float64, peak *int, mu *sync.Mutex) {
	defer wg.Done()

	arrival := time.Now()

	if len(library) == cap(library) {
		fmt.Printf("Student %d waiting... (library full)\n", student.ID)
	}

	// Enter library
	library <- true

	waitTime := time.Since(arrival).Seconds()
	waitChan <- waitTime

	mu.Lock()
	if len(library) > *peak {
		*peak = len(library)
	}
	mu.Unlock()

	fmt.Printf("Student %d entered library, will study for %d hours\n",
		student.ID, student.StudyHours)

	time.Sleep(time.Duration(student.StudyHours) * time.Second)

	fmt.Printf("Student %d left library after %d hours\n",
		student.ID, student.StudyHours)

	<-library
}

func main() {

	fmt.Println("=== Library Simulation ===")
	fmt.Println("Library capacity: 30 students")
	fmt.Println("Total students today: 100")
	fmt.Println("Simulation: 1 second = 1 hour")

	start := time.Now()

	library := make(chan bool, 30)
	waitChan := make(chan float64, 100)

	var wg sync.WaitGroup
	var mu sync.Mutex

	peakOccupancy := 0

	for i := 1; i <= 100; i++ {

		wg.Add(1)

		student := Student{
			ID:         i,
			StudyHours: rand.Intn(4) + 1,
		}

		go study(student, library, &wg, waitChan, &peakOccupancy, &mu)

		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()
	close(waitChan)

	totalWait := 0.0
	count := 0

	for w := range waitChan {
		totalWait += w
		count++
	}

	averageWait := totalWait / float64(count)

	fmt.Println("=== Simulation Complete ===")
	fmt.Println("Total students served: 100")
	fmt.Printf("Library was open for: %d hours\n", int(time.Since(start).Seconds()))
	fmt.Printf("Average wait time: %.2f hours\n", averageWait)
	fmt.Printf("Peak occupancy: %d students\n", peakOccupancy)
}
