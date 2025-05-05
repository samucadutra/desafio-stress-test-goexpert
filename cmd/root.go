package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	successStatusCode = 200
	defaultURL        = "https://google.com"
	defaultRequests   = 100
	defaultWorkers    = 10
)

type TestResults struct {
	TotalTime     time.Duration
	StatusCounts  map[int]int
	TotalRequests int
}

var (
	url         string
	totalReq    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "desafio-stress-test-goexpert",
	Short: "A CLI tool for stress test requests",
	Long:  `This application performs stress testing on a web service or other endpoint and generates a report.`,
	Run:   runStressTest,
}

func runStressTest(cmd *cobra.Command, args []string) {
	results := executeStressTest()
	generateReport(results)
}

func executeStressTest() TestResults {
	startTime := time.Now()
	requests := make(chan int, totalReq)
	responses := make(chan int, totalReq)
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go sendRequestsWorker(requests, responses, &wg)
	}

	for i := 0; i < totalReq; i++ {
		requests <- i
	}
	close(requests)
	wg.Wait()
	close(responses)

	statusCounts := make(map[int]int)
	for status := range responses {
		statusCounts[status]++
	}

	return TestResults{
		TotalTime:     time.Since(startTime),
		StatusCounts:  statusCounts,
		TotalRequests: totalReq,
	}
}

func generateReport(results TestResults) {
	fmt.Println("Stress Test Report:")
	fmt.Printf("Total time: %v\n", results.TotalTime)
	fmt.Printf("Total requests: %d\n", results.TotalRequests)
	fmt.Printf("HTTP 200 responses: %d\n", results.StatusCounts[successStatusCode])

	successCount := results.StatusCounts[successStatusCode]
	if successCount < results.TotalRequests {
		fmt.Printf("Percentage of successful requests: %.2f%%\n",
			float64(successCount)/float64(results.TotalRequests)*100)
		fmt.Println("Other status codes:")
		for code, count := range results.StatusCounts {
			if code != successStatusCode {
				fmt.Printf("HTTP %d: %d\n", code, count)
			}
		}
	} else {
		fmt.Println("All requests were successful.")
	}
}

func sendRequestsWorker(requests <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}

	for range requests {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			results <- resp.StatusCode
			continue
		}

		results <- resp.StatusCode
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", defaultURL, "URL to be stress tested")
	rootCmd.Flags().IntVarP(&totalReq, "requests", "r", defaultRequests, "Total number of requests to be sent")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", defaultWorkers, "Number of concurrent requests to be sent")
}
