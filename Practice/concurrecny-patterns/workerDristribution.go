package main

import (
	"fmt"
)

const (
	chunkSize = 100000
	totalSize = 100000000
)

func main() {

	chunkDistributionChannel := make(chan []float64)
	outputChan := make(chan float64)

	// spawn pool of 4 workers
	for i := 0; i < 4; i++ {
		go worker(chunkDistributionChannel, outputChan)
	}

	// generate the chunks
	numberOfChunks := generator(chunkDistributionChannel)

	workerOutputSum := 0.0

	for i := 0; i < numberOfChunks; i++ {
		workerOutputSum += <-outputChan
	}

	fmt.Println("Calculated Average :", workerOutputSum/float64(numberOfChunks))

}

func worker(inputChan <-chan []float64, outputChan chan<- float64) {
	for chunk := range inputChan {

		// calculate avg
		var sum float64
		for _, val := range chunk {
			sum += val
		}
		avg := sum / float64(len(chunk))
		outputChan <- avg

	}
}

func generator(chunkDistributionChannel chan<- []float64) int {

	work := make([]float64, totalSize)
	for i := 0; i < totalSize; i++ {
		work[i] = float64(i + 1)
	}

	numberOfChunks := len(work) / chunkSize

	go func() {
		for i := 0; i < numberOfChunks; i++ {
			end := (i + 1) * chunkSize
			if end > len(work) {
				fmt.Println("Catched")
				end = len(work)
			}

			chunkDistributionChannel <- work[i*chunkSize : end]
		}

		close(chunkDistributionChannel)
	}()

	return numberOfChunks
}
