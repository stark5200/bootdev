/*
type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
*/

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	return msg.getMessage(), len(msg.getMessage())*3
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (c contractor) getName() string {
	return c.name
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// interfaces 3

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body)*2
	}
	return len(e.body)*5
}

func (e email) format() string {
	s := "Subscribed"
	if (!e.isSubscribed) {
		s = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, s)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
