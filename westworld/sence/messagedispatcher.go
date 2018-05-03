package sence

import (
	"fmt"
	"sort"
	"time"

	"github.com/laonsx/statemachines/westworld"
)

var messageQueue PriorityQueue

type PriorityQueue []*Telegram

func (queue PriorityQueue) Len() int {

	return len(queue)
}

func (queue PriorityQueue) Less(i, j int) bool {

	return queue[i].DispathTime < queue[j].DispathTime
}

func (queue PriorityQueue) Swap(i, j int) {

	queue[i], queue[j] = queue[j], queue[i]
}

func (queue PriorityQueue) HaveNewDispatchMessage(currTime int64) bool {

	if queue.Len() == 0 {

		return false
	}

	if queue[0].DispathTime > currTime {

		return false
	}

	return true
}

func (queue *PriorityQueue) Push(telegram *Telegram) {

	*queue = append(*queue, telegram)
	sort.Sort(*queue)
}

func (queue *PriorityQueue) Pop() *Telegram {

	telegram := (*queue)[0]
	*queue = (*queue)[1:]

	return telegram
}

func DispatchMessage(receiverId, senderId int, msg Message) bool {

	receiver, _ := EntityManager.Load(receiverId)
	if receiver == nil {

		fmt.Println("..1")
		return false
	}

	if entity, ok := receiver.(westworld.BaseEntity); ok {
		fmt.Println("..3")
		return entity.HandleMessage(&Telegram{
			DispathTime: 0,
			Receiver:    receiverId,
			Sender:      senderId,
			Msg:         msg,
		})
	}
	fmt.Println("..2")
	return false
}

func DispatchDelayMessage(delay int64, receiver, sender int, msg Message) {

	currTime := time.Now().Unix()
	messageQueue.Push(&Telegram{
		DispathTime: currTime + delay,
		Receiver:    receiver,
		Sender:      sender,
		Msg:         msg,
	})
}

type Telegram struct {
	DispathTime      int64
	Receiver, Sender int
	Msg              Message
}

type Message = int

const (
	MSG_TEST1 Message = iota + 1000
	MSG_TEST2
	MSG_TEST3
	MSG_TEST4
	MSG_TEST5
)
