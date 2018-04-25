package sence

type PriorityQueue []*Telegram

func (queue PriorityQueue) Begin() *Telegram {
	if queue.Count() > 0 {
		return queue[0]
	}
	return nil
}

func (queue PriorityQueue) Count() int {
	return len(queue)
}

func (queue *PriorityQueue) Insert() {

}

func DispatchMessage(delay int64, receiver, sender int, msg Message) bool {

	return true
}

type Telegram struct {
	DispathTime      int64
	Receiver, Sender int
	Msg              Message
}

type Message = int
