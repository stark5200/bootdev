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

