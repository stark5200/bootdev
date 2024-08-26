package main

import "fmt"

/*
Maps
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
Copy icon
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
Copy icon
The len() function works on a map, it returns the total number of key/value pairs.

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2
Copy icon
Assignment
We can speed up our contact-info lookups by using a map!

Key-based map lookup: O(1)
Slice brute-force search: O(n)
Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user's name and phone number. The first name in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/

type MyError struct{}

func (m MyError) Error() string {
	return "invalid sizes"
}

/// Note don't use loops inside data structures

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, MyError{}
	}

	userMap := make(map[string]user)
	for i, n := range names{
		userMap[n] = user{name: n, phoneNumber: phoneNumbers[i]}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

/*
Mutations
Insert an element
m[key] = elem
Copy icon
Get an element
elem = m[key]
Copy icon
Delete an element
delete(m, key)
Copy icon
Check if a key exists
elem, ok := m[key]
Copy icon
If key is in m, then ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Assignment
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
Note on passing maps
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.
*/

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if !ok {
		return false, fmt.Errorf("not found")
	} 
	if !elem.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

/*
Count Instances
Remember that you can check if a key is already present in a map by using the second return value from the index operation.

names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}
Copy icon
Assignment
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the getCounts function. It takes a slice of strings messagedUsers and a map of string -> int validUsers. It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, mUser := range messagedUsers {
		if n, ok := validUsers[mUser]; ok {
			validUsers[mUser] = n+1
		}
	}
}

