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

func (e email1) cost() int {
	if e.isSubscribed {
		return len(e.body)*2
	}
	return len(e.body)*5
}

func (e email1) format() string {
	s := "Subscribed"
	if (!e.isSubscribed) {
		s = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, s)
}

type expense1 interface {
	cost() int
}

type formatter interface {
	format() string
}

type email1 struct {
	isSubscribed bool
	body         string
}

/*
Type assertions 

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

c, ok := s.(circle)
if !ok {
	// log an error if s isn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius
*/

// interfaces4 

func getExpenseReport(e expense2) (string, float64) {
	e1, ok := e.(email2)
	if ok {
		// return email e1 attributes
		toAdress := e1.toAddress
		cost := e1.cost()
		return toAdress, cost
	}
	
	e2, ok := e.(sms)
	if ok {
		// return sms e2 attributes
		toPhoneNumber := e2.toPhoneNumber
		cost := e2.cost()
		return toPhoneNumber, cost
	}
	return "", 0.0
}

// don't touch below this line

type expense2 interface {
	cost() float64
}

type email2 struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email2) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

// interfaces 5 type switches

/*
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
*/

// using switches instead
func getExpenseReport2(e expense2) (string, float64) {
	switch v := e.(type) {
	case email2:
		toAdress := v.toAddress
		cost := v.cost()
		return toAdress, cost
	case sms:
		toPhoneNumber := v.toPhoneNumber
		cost := v.cost()
		return toPhoneNumber, cost
	default:
		return "", 0.0
	}
}
