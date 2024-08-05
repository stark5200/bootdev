package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if (mToSend.sender.number == 0 || mToSend.sender.name == "" || mToSend.recipient.number == 0 || mToSend.recipient.name == "") {
		return false }
	return true
}

// embedded struct
type sender struct {
	user
	rateLimit int
}


type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	formatted := "Authorization: Basic " + a.username + ":" + a.password
	return formatted
}

/*To be honest, you should not stress about memory layout. However, if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help. You can also use the reflect package to debug the memory layout of a struct:*/

//typ := reflect.TypeOf(stats{})
//fmt.Printf("Struct is %d bytes\n", typ.Size())

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}



/// example

type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50

package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below
