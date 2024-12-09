package main

import "fmt"

type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

type publisher struct {
	subscribers []Subscriber
}

func GetNewPublisher() publisher {
	return publisher{subscribers: make([]Subscriber, 0)}
}

func (p *publisher) Register(subscriber Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p publisher) NotifyAll(msg string) {
	for _, subs := range p.subscribers {
		fmt.Println("Publisher notifying Subscriber with id: ", subs.(subscriber).subscriberId)
	}
}

type subscriber struct {
	subscriberId string
}

func GetNewSubscriber(Id string) subscriber {
	return subscriber{subscriberId: Id}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println("Subscriber Received Message: ", msg, " for subscriber id: ", s.subscriberId)
}

func main() {
	publisher := GetNewPublisher()

	subscriber := GetNewSubscriber("1")
	subscriber2 := GetNewSubscriber("2")

	publisher.Register(subscriber)
	publisher.Register(subscriber2)
	publisher.NotifyAll("Hello notifying subscriber")
}
