function getCleanRank(reviewWords) {
  let bad = 0
  
  if (reviewWords.includes('dang')) {bad += 1;} 
  if (reviewWords.includes('shoot')) {bad += 1;} 
  if (reviewWords.includes('heck')) {bad += 1;} 

  if ( bad == 0) {
    return 'clean'
  } else if ( bad == 1) {
    return 'dirty'
  } else if ( bad >= 2) {
    return 'filthy'
  } 
}


// Don't edit below this line

function test(reviewWords) {
  const cleanRank = getCleanRank(reviewWords)
  console.log(`'${reviewWords}' has rank: ${cleanRank}`)
}

test([ 'avril', 'lavigne', 'has', 'best', 'dang', 'tour' ])
test([ 'what', 'a', 'bad', 'film' ])
test([ 'oh', 'my', 'heck', 'I', 'hated', 'it' ])
test([ 'ripoff' ])
test([ 'That', 'was', 'a', 'pleasure' ])
test([ 'shoot!', 'I', 'cant', 'say', 'I', 'liked', 'the', 'dang', 'thing' ])
test([ 'shoot', 'dang', 'heck' ])

const movies = [
  'oh brother where art thou',
  'oceans eleven',
  'fight club',
  'the island',
  'shutter island',
  'the magnificent seven'
]

function logArray(arr) {
  for (const a of arr) {
    console.log(` - ${a}`)
  }
  console.log('---')
}

// don't touch above this line

logArray(movies.slice(2))
logArray(movies.slice(0, -2))

for (let i = 0; i < 100; i++) {
  console.log(i)
}

const printCleanReviews = (reviews, badWord) => {
  // ?
}

// don't touch below this line

printCleanReviews([ 'The movie sucks', 'I love it', 'What garbage', 'so sucky' ], 'suck')
console.log('---')
printCleanReviews([ 'The movie sucks', 'I love it', 'What darn crap', 'darn, so sucky' ], 'darn')
console.log('---')

const movieExists = (movies, title) => {
  for (let i = 0; i < movies.length; i++) {
    console.log("Looking at:", movies[i]);
    if (movies[i] == title) {
      console.log("Found:", movies[i]);
      break;
    }
  }
}
// don't touch below this line

movieExists([ 'Incredibles', 'Tangled', 'Frozen' ], 'Frozen')
console.log('---')
movieExists([ 'The Quick and the Dead', 'The Magnificent 7', 'Frozen' ], 'The Magnificent 7')
console.log('---')
movieExists([ 'Dead', 'Alive', 'Flight', 'Rocky' ], 'Flight')
console.log('---')
movieExists([ 'Dead', 'Alive', 'Flight', 'Rocky' ], 'None')
console.log('---')

for (let i = 5; i < 16; i++) {
  console.log(i)
}

// clean reviews
const printCleanReviews = (reviews, badWord) => {
  for (let i = 0; i < reviews.length; i++) {
    if (reviews[i].includes(badWord)) {
      continue;
    }
    console.log("Clean review:", reviews[i]);
  }
}

for (let i = 10; i > 0; i--) {
  if (i == 1) {
    console.log(`${i} star`);
  }
  else {
    console.log(`${i} stars`)
  }
}
      
const printCleanReviews2 = (reviews, badWord) => {
  for (let review of reviews) {
    if (!review.includes(badWord)) {
      console.log("clean review:", review)
    }
  }
}

// don't touch below this line

printCleanReviews2([ 'The movie sucks', 'I love it', 'What garbage', 'so sucky' ], 'suck')
console.log('---')
printCleanReviews2([ 'The movie sucks', 'I love it', 'What darn crap', 'darn, so sucky' ], 'darn')
console.log('---')
function getMovieRecord(title, stars, username) {
  return {
    title: title, 
    stars: stars, 
    username: username, 
    id: title+"-"+username,
  }
}

// Don't touch below this line

logObject(getMovieRecord('oh brother where art thou', 3, 'wagslane'))
logObject(getMovieRecord('frozen', 5.5, 'elonmusk'))
logObject(getMovieRecord('toy story', 4, 'prince'))

function logObject(obj) {
  for (const key in obj) {
    console.log(` - ${key}: ${obj[key]}`)
  }
  console.log('---')
}

const name = 'Apple'
const radius = 2
const color = 'red'
const apple = {
  name,
  radius,
  color,
}

function addID(movieRecord) {
  movieRecord.id = `${movieRecord.title}-${movieRecord.username}`
  return movieRecord
}

const company = {
  employees: {
    ceo: {
      name: 'Elon',
      salary: 0
    },
    engineers: [
      {
        name: 'Marie',
        salary: 225000
      },
      { 
        name: 'George',
        salary: 205000
      }
    ]
  }
}

const user = {
  getFirstReview() {
    // ?
  },
  reviews: [ 'I hate Ice Age', 'I didn\'t enjoy it at all', 'What a fabulous film' ],
  name: 'Bob Doogle'
}

// don't touch below this line

console.log(`${user.name}'s first review is: ${user.getFirstReview()}`)

// methods quiz

const tree = {
  height: 256,
  color: 'green',
  cut() {
    // ...
  }
}
const tree2 = {
  height: 256,
  color: 'green',
  cut() {
    this.height /= 2
  }
}
const author = {
  firstName: 'Lane',
  lastName: 'Wagner',
  getName: () => {
    return `${this.firstName} ${this.lastName}`
  }
}
console.log(author.getName())
// Prints: undefined undefined
// because the parent scope (the scope outside of the author object)
// never defined .firstName and .lastName properties

try {
  const car = {}
  console.log(car.make.name)
} catch (err) {
  console.log(err.message)
  // Cannot read properties of undefined (reading 'name')
} finally {
  console.log('I will always run')
}

try {
  const car = {}
  console.log(car.make.name)
} catch (err) {
  console.log(err.message)
  // Cannot read properties of undefined (reading 'name')
} finally {
  console.log('I will always run')
}

throw new Error('something bad happened')

function getMovieRecord(movieId) {
  if (movieId === 1) {
    return { name: 'Apollo 13', duration: 128 }
  }
  if (movieId === 2) {
    return { name: '2001: A Space Odyssey', duration: 300 }
  }
  if (movieId === 3) {
    return { name: 'Interstellar', duration: 4000 }
  }
  throw new Error('movie id not found')
}

function logObject(obj) {
  for (const key in obj) {
    console.log(` - ${key}: ${obj[key]}`)
  }
  console.log('---')
}

/*
This is the final required lesson for this course
There are 10 required lessons that you have not yet completed. Use the dropdown menus at the top left of the screen to find the incomplete lessons and finish them!
Running Node.js
You've been able to run Python code locally by installing the Python interpreter and then running:

python main.py
Copy icon
Similarly, if you install Node.js on your local machine, you'll be able to run

node main.js
Copy icon
Node.js is basically the JS equivalent to the python interpreter you've used before. Before Node, the only way to run JavaScript code was in the browser. You can still do that from scratch as well using your browser's dev tools!

 NVM
I highly recommend using "Node Version Manager" to manage your Node.js installation. NVM makes it easy to:

Install multiple versions of Node
Update your Node version
Keep your Node version configurations separate on a per-project basis
*/

const review = {
  text: 'This movie was awful',
  stars: 2,
  author: {
    firstName: 'Johnny',
    lastName: 'Comelately',
    createdAt: '2022-08-17T15:41:25+00:00'
  }
}

// don't touch above this line

console.log(review.author.firstName)

const user = {
  getFirstReview() {
    return this.reviews[0]
  },
  reviews: [ 'I hate Ice Age', 'I didn\'t enjoy it at all', 'What a fabulous film' ],
  name: 'Bob Doogle'
}

// don't touch below this line

console.log(`${user.name}'s first review is: ${user.getFirstReview()}`)

function getState(review) {
  return review.author?.location?.state
}

// don't touch below this line

function test(review) {
  const state = getState(review)
  if (state) {
    console.log(`Adding ${state} to the database...`)
  } else {
    console.log('No state found...')
  }
}

test({
  text: 'This movie was awful',
  stars: 2,
  author: {
    firstName: 'Johnny',
    lastName: 'Comelately',
    createdAt: '2022-08-17T15:41:25+00:00',
    location: {
      state: 'Utah'
    }
  }
})

test({
  text: 'This movie was okay...',
  stars: 5
})

test({
  text: 'This movie was awful',
  stars: 2,
  author: {
    firstName: 'Jill',
    lastName: 'Comelately',
    createdAt: '2022-08-17T15:41:25+00:00',
    location: {
      state: 'Nevada'
    }
  }
})

test({
  text: 'This movie was awful',
  stars: 2,
  author: {
    firstName: 'George',
    lastName: 'Jimenez',
    createdAt: '2022-08-17T15:41:25+00:00'
  }
})

const getCountsByTitle = (movies) => {
  let movieCounts = {}
  for (const movie of movies) {
    if (!movieCounts[movie]) {
      movieCounts[movie] = 0
    }
    movieCounts[movie]++
  }
  return movieCounts
}

// don't touch below this line

function test(movies) {
  const counts = getCountsByTitle(movies)
  for (const [ movie, count ] of Object.entries(counts)) {
    console.log(`'${movie}' has ${count} reviews`)
  }
  console.log('---')
}

function getMovieRecord(movieId) {
  if (movieId === 1) {
    return { name: 'Apollo 13', duration: 128 }
  }
  if (movieId === 2) {
    return { name: '2001: A Space Odyssey', duration: 300 }
  }
  if (movieId === 3) {
    return { name: 'Interstellar', duration: 4000 }
  }
  throw new Error('movie id not found')
}

/*
package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	return ParsedURL{
		protocol: "",
		username: "",
		password: "",
		hostname: "",
		port:     "",
		pathname: "",
		search:   "",
		hash:     "",
	}
}

*/

function main() {
  try {
    logObject(getMessageRecord(1));
    logObject(getMessageRecord(2));
    logObject(getMessageRecord(3));
    logObject(getMessageRecord(4));
  } catch (err) {
    console.log(err.message);
  }
}

// don't touch below this line

function getMessageRecord(messageId) {
  if (messageId === 1) {
    return { content: "Welcome to Textio!", timestamp: "2025-01-01T12:00:00Z" };
  }
  if (messageId === 2) {
    return {
      content: "Your order has shipped",
      timestamp: "2025-01-02T12:00:00Z",
    };
  }
  if (messageId === 3) {
    return {
      content: "Reminder: Payment due soon",
      timestamp: "2025-01-03T12:00:00Z",
    };
  }
  throw new Error("text id not found");
}

function logObject(obj) {
  for (const key in obj) {
    console.log(` - ${key}: ${obj[key]}`);
  }
  console.log("---");
}

main();

let messageLen = 10;
let maxMessageLen = 20;
console.log(
  "Trying to send a message of length:",
  messageLen,
  "and a max length of:",
  maxMessageLen,
);

// don't touch above this line

if (messageLen <= maxMessageLen) {
  console.log("Message sent");
} else {
  console.log("Message not sent");
}

function billingCost(plan) {
  switch (plan) {
    case "basic":
      return 10.0;
    case "pro":
      return 20.0;
    case "enterprise":
      return 50.0;
    default:
      return 0.0;
  }
}

// don't touch below this line

console.log(`The cost for a basic plan is $${billingCost("basic").toFixed(2)}`);
console.log(`The cost for a pro plan is $${billingCost("pro").toFixed(2)}`);
console.log(
  `The cost for a enterprise plan is $${billingCost("enterprise").toFixed(2)}`,
);
console.log(`The cost for a free plan is $${billingCost("free").toFixed(2)}`);
console.log(
  `The cost for a unknown plan is $${billingCost("unknown").toFixed(2)}`,
);

const userCredits = -2;

// don't touch above this line

if (userCredits >= 1) {
  console.log("Sending message...");
} else {
  console.log("Not enough credits.");
}

const name = "James Holden";
const provider = "AT&T";
const phoneNumber = "555-123-4567";
const subscriptionType = null;

// don't touch above this line

console.log(
  `Creating ${subscriptionType ?? "Guest"} Profile for ${name} with ${provider} at ${phoneNumber}.`,
);

console.log("Starting Textio server...");
