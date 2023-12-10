package filehandle

import (
	"sync"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func ProcessFile(fileData string, routines int) utils.Summary {
	var summary utils.Summary
	var wg sync.WaitGroup
	channel := make(chan utils.Summary)
	chunk := len(fileData) / routines
	startIndex := 0
	endIndex := chunk
	for iterations := 0; iterations < routines; iterations++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			Counts(fileData[start:end], channel)
		}(startIndex, endIndex)
		// go Counts(fileData[startIndex:endIndex], channal)
		startIndex = endIndex
		endIndex += chunk

	}
	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(channel)
	}()

	// Receive results from the channel and aggregate them
	for counts := range channel {
		summary.LineCount += counts.LineCount
		summary.WordsCount += counts.WordsCount
		summary.VowelsCount += counts.VowelsCount
		summary.PuncuationsCount += counts.PuncuationsCount
	}

	return summary
}
