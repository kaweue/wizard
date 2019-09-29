package install

import "time"

func Install(p chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 100)
			p <- i
		}
		close(p)
	}()
}
