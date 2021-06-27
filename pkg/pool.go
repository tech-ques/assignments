package pkg

import (
	"fmt"
	"time"
)

const (
	LIMIT int = 10
)

func worker(id int, jobs <-chan User, results chan<- User) {
	for j := range jobs {
		fmt.Println("worker:", id, "started sending notification for user: ", j.FirstName)
		Notify(j)
		time.Sleep(100 * time.Millisecond)
		// fmt.Println("worker:", id, "finished sending notification for user: ", j.FirstName)
		results <- j
	}
}

func StartProcess(users []User) {
	jobs := make(chan User, len(users))
	results := make(chan User, len(users))

	// limit workers initiated
	for w := 1; w <= LIMIT; w++ {
		go worker(w, jobs, results)
	}

	// Sending jobs
	for _, u := range users {
		jobs <- u
	}

	close(jobs)

	for a := 0; a < len(users); a++ {
		j := <-results
		fmt.Println("finished sending notification for user: ", j.FirstName)
	}
}
