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

func study(id int, studyHours int, library chan bool, stats *LibraryStats, start time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	arrival := time.Now()

	if len(library) == cap(library) {
		fmt.Printf("Student %d waiting... (library full)\n", id)
	}

	// Enter library
	library <- true

	waitTime := time.Since(arrival).Seconds()
	stats.AddWaitTime(time.Duration(waitTime))

	hour := int(time.Since(start).Seconds())
	if hour >= len(stats.HourlyActivity) {
		hour = len(stats.HourlyActivity)
	}

	stats.RecordEntry(hour)



	fmt.Printf("Student %d entered library, will study for %d hours\n",
		id, studyHours)

	time.Sleep(time.Duration(studyHours) * time.Second)

	<-library

	stats.RecordExit()

	fmt.Printf("Student %d left library after %d hours\n",
		id, studyHours)
}

// Task 5.2
type LibraryStats struct {
	TotalStudents int
	TotalWaitTime time.Duration
	PeakOccupancy int
	CurrentOccupancy int
	HourlyActivity []int // tracks how many students enter each hour
	mu sync.Mutex
}

func (stats *LibraryStats) RecordEntry(hour int) {
 	// TODO: Update statistics when student enters
	stats.mu.Lock()
	defer stats.mu.Unlock()

	stats.CurrentOccupancy++
	stats.HourlyActivity[hour]++

	if stats.CurrentOccupancy > stats.PeakOccupancy {
		stats.PeakOccupancy = stats.CurrentOccupancy
	}
}


func (stats *LibraryStats) RecordExit() {
 	// TODO: Update statistics when student exits
	stats.mu.Lock()
	defer stats.mu.Unlock()

	stats.TotalStudents++
	stats.CurrentOccupancy--

}

func (stats *LibraryStats) AddWaitTime(wait time.Duration) {
	// TODO: Add Waiting Time when student enters
	stats.mu.Lock()
	defer stats.mu.Unlock()

	stats.TotalWaitTime += wait
}


func (stats *LibraryStats) PrintReport(totalTime time.Duration) {
 	// TODO: Print detailed statistics
	avgWait := float64(stats.TotalWaitTime) / float64(time.Duration(stats.TotalStudents))

	quietHour := 0
	minStudents := stats.HourlyActivity[0]

	totalStudentHours := 0

	for h, students := range stats.HourlyActivity {
		totalStudentHours += students

		if students < minStudents {
			minStudents = students
			quietHour = h
		}
	}

	avgStudy := float64(totalStudentHours) / float64(stats.TotalStudents)
	fmt.Println("\n=== Simulation Complete ===")
	fmt.Printf("Total students served: %d\n", stats.TotalStudents)
	fmt.Printf("Library was open for: %.0f hours\n", totalTime.Seconds())
	fmt.Printf("Average wait time: %.2f hours\n", avgWait)
	fmt.Printf("Peak occupancy: %d students\n", stats.PeakOccupancy)
	fmt.Printf("Quietest hour: Hour %d (%d students)\n", quietHour, minStudents)
	fmt.Printf("Total student-hours: %d hours\n", totalStudentHours)
	fmt.Printf("Average study duration: %.2f hours per student\n", avgStudy)
}


func main() {

	fmt.Println("=== Library Simulation ===")
	fmt.Println("Library capacity: 30 students")
	fmt.Println("Total students today: 100")
	fmt.Println("Simulation: 1 second = 1 hour")

	library := make(chan bool, 30)
	stats := LibraryStats {
		HourlyActivity: make([]int, 24),
	}

	var wg sync.WaitGroup

	start := time.Now()

	students := make([]Student, 100)
	for i := range 100 {

		students[i] = Student{
			ID:         i + 1,
			StudyHours: rand.Intn(4) + 1,
		}
	}

	for _, s := range students {
		wg.Add(1)
		go study(s.ID, s.StudyHours, library, &stats, start, &wg)
	}

	wg.Wait()
	totalTime := time.Since(start)
	
	stats.PrintReport(totalTime)

}
