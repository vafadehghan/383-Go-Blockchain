package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := WorkQueue{make(chan Worker, maxJobs), make(chan interface{}, maxJobs)}
	for i := uint(0); i < nWorkers; i++ {
		go q.worker()
	}
	return &q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	for incomingTasks := range queue.Jobs {
		queue.Results <- incomingTasks.Run()
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
	for _ = range queue.Jobs {
	}
}
