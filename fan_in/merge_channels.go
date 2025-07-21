package fan_in

import "sync"

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	for _, channel := range channels {
		ch := channel
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range ch {
				result <- n
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
