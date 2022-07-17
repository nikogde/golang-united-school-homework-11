package batch

import (
	"time"
	"sync"
)

type user struct {
	ID int64
}

type users struct {
	mu    sync.Mutex
	users []user
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	result := &users{}
	ch := make(chan int, pool)
	var wg sync.WaitGroup

	for i := 0; i < int(n); i++ {
		wg.Add(1)
		ch <- int(i)
		go func(i int) {
			u := getOne(int64(i))
			<-ch
			result.mu.Lock()
			result.users = append(result.users, u)
			result.mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	return result.users
}
