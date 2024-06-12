package main

import (
    "fmt"
    "time"
)

// Task represents a unit of work to be processed by a worker.
type Task struct {
    ID   int
    Work func() error
}

// worker is a goroutine that processes tasks from the tasks channel and sends results to the results channel.
func worker(id int, tasks <-chan Task, results chan<- error) {
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task.ID)
        result := task.Work() //execute the task
        results <- result //push the result to the results channel
    }
}

// startWorkerPool initializes the specified number of workers to process tasks from the tasks channel.
func startWorkerPool(numWorkers int, tasks <-chan Task, results chan<- error) {
    for i := 1; i <= numWorkers; i++ {
        go worker(i, tasks, results)
    }
}

func main() {
	start := time.Now()
    numWorkers := 30
    numTasks := 100
    
    tasks := make(chan Task, numTasks)// buffered channel
    results := make(chan error, numTasks)//buffered channel

    startWorkerPool(numWorkers, tasks, results)
    
    // Submit tasks
    for i := 1; i <= numTasks; i++ {
        task := Task{
            ID: i,
            Work: func() error {
                // Simulate work
                time.Sleep(time.Second)
                return nil
            },
        }
        tasks <- task
    }
    
    close(tasks)
    
    // Collect results
    for i := 1; i <= numTasks; i++ {
        err := <-results //reading the data from the results channel
        if err != nil {
            fmt.Printf("Task %d failed with error: %v\n", i, err)
        } else {
            fmt.Printf("Task %d succeeded\n", i)
        }
    }
	fmt.Printf("Took %fs \n", time.Since(start).Seconds())
}
