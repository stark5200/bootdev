/// Enums

/*
Lack of Enums
My least favorite part of Go? Glad you asked. It's Go's lack of enums, sum types, tagged unions, etc. Compared to other statically typed languages like:

Rust
TypeScript
OCaml
Go's type system just isn't as powerful. It's more similar to C's type system than it is to Rust's. It's more concerned with simplicity than it is with expressiveness.

Error Handling
In Rust, like Go, errors are just values. In Go, we write something like this:

user, err := getUser()
if err != nil {
    return fmt.Errorf("failed to get user: %w", err)
}
// do something with user
Copy icon
In Rust, we can do something like this:

let user_result = get_user();
let user = match user_result {
    Ok(user) => user,
    Err(error) => return Err(format!("failed to get user: {}", error)),
};
Copy icon
In Rust, the get_user function returns a Result type: a type that is either an Ok or an Err. The compiler forces the developer to handle the error case before they can continue with the happy path (using the user data).

In Go, the developer can choose to happily ignore the error value if they choose and use the user data, even if it's invalid (probably nil or an empty struct).

The support for enums in Rust makes it easier to write bug-free code.

Assignment
A lazy Go programmer wrote the handleEmailBounce function... just because the compiler doesn't force us to check errors doesn't mean we shouldn't!

Take a look at the updateStatus and track methods in the user.go file. Handle their errors properly, and use the fmt.Errorf function and the %w formatting verb to add useful context to the errors.

If updateStatus fails, return an error saying "error updating user status: ERR"
If track fails, return an error saying "error tracking user bounce: ERR"
Where ERR is the error returned by the method.
*/

package main

import (
	"fmt"
)

func (a *analytics) handleEmailBounce(em email) error {
	err1 := em.recipient.updateStatus(em.status)
	if err1 != nil {
	    return fmt.Errorf("error updating user status: %w", err1)
	}
	err2 := a.track(em.status)
	if err2 != nil {
	    return fmt.Errorf("error tracking user bounce: %w", err2)
	}
	return nil
}

/// Iota 

/*
Iota
Go has a language feature, that when used with a type alias (and if you squint really hard), kinda looks like an enum (but it's not). It's called iota.

type sendingChannel int

const (
    Email sendingChannel = iota
    SMS
    Phone
)
Copy icon
The iota keyword is a special keyword in Go that creates a sequence of numbers. It starts at 0 and increments by 1 for each constant in the const block. So in the example above, Email is 0, SMS is 1, and Phone is 2.

Go developers sometimes use iota to create a sequence of constants to represent a set of related values, much like you would with an enum in other languages. But remember, it's not an enum. It's just a sequence of numbers.

Assignment
Define an emailStatus type that uses iota syntax to represent the following states:

emailBounced: 0
emailInvalid: 1
emailDelivered: 2
emailOpened: 3
*/

type emailStatus int

const (
    emailBounced emailStatus = iota
    emailInvalid
    emailDelivered
	emailOpened
)