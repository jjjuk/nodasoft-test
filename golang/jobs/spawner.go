package jobs

import "sync"

type TaskSpawner struct {
	running bool
	wg      sync.WaitGroup
}

func NewTaskSpawner() *TaskSpawner {
	return &TaskSpawner{running: false}
}

func (sp *TaskSpawner) IsRunning() bool {
	return sp.running
}

func (sp *TaskSpawner) Start(tc chan *Task) {
	sp.running = true
	sp.wg.Add(1)
	go func() {
		defer sp.wg.Done()

		for id := 1; sp.IsRunning(); id++ {
			task := NewTask(id)
			tc <- task // передаем таск на выполнение
			// time.Sleep(20 * time.Millisecond) // немного потротлить его хочется да
		}
	}()
}

func (sp *TaskSpawner) Stop() {
	sp.running = false
	sp.wg.Wait()
}
