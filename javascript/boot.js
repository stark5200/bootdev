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
