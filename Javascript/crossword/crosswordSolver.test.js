const crosswordSolver = require('./crosswordSolver');
const assert = require('assert');
const error = "Error";
let testCounter = 0;
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'alan', 'ciao', 'anta']
const solution = `casa
i..l
anta
o..n`

assert.strictEqual(crosswordSolver(puzzle, words), solution);
console.log(`Test ${++testCounter} passed.`)

const puzzle1 = `...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
100000..0...0..
.0.....1001000.
.0.1....0.0....
.10000000.0....
.0.0......0....
.0.0.....100...
...0......0....
..........0....`
const words1 = [
  'sun',
  'sunglasses',
  'suncream',
  'swimming',
  'bikini',
  'beach',
  'icecream',
  'tan',
  'deckchair',
  'sand',
  'seaside',
  'sandals',
]
const solution1 = `...s...........
..sunglasses...
...n....u......
.s......n...s..
.w....deckchair
bikini..r...n..
.m.....seaside.
.m.b....a.a....
.icecream.n....
.n.a......d....
.g.c.....tan...
...h......l....
..........s....`

assert.strictEqual(crosswordSolver(puzzle1, words1), solution1);
console.log(`Test ${++testCounter} passed.`)

const puzzle2 = `..1.1..1...
10000..1000
..0.0..0...
..1000000..
..0.0..0...
1000..10000
..0.1..0...
....0..0...
..100000...
....0..0...
....0......`
const words2 = [
  'popcorn',
  'fruit',
  'flour',
  'chicken',
  'eggs',
  'vegetables',
  'pasta',
  'pork',
  'steak',
  'cheese',
]
const solution2 = `..p.f..v...
flour..eggs
..p.u..g...
..chicken..
..o.t..t...
pork..pasta
..n.s..b...
....t..l...
..cheese...
....a..s...
....k......`

assert.strictEqual(crosswordSolver(puzzle2, words2), solution2);
console.log(`Test ${++testCounter} passed.`)

const puzzle3 = `...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
100000..0...0..
.0.....1001000.
.0.1....0.0....
.10000000.0....
.0.0......0....
.0.0.....100...
...0......0....
..........0....`
const words3 = [
  'sun',
  'sunglasses',
  'suncream',
  'swimming',
  'bikini',
  'beach',
  'icecream',
  'tan',
  'deckchair',
  'sand',
  'seaside',
  'sandals',
].reverse()
const solution3 = `...s...........
..sunglasses...
...n....u......
.s......n...s..
.w....deckchair
bikini..r...n..
.m.....seaside.
.m.b....a.a....
.icecream.n....
.n.a......d....
.g.c.....tan...
...h......l....
..........s....`

assert.strictEqual(crosswordSolver(puzzle3, words3), solution3);
console.log(`Test ${++testCounter} passed.`)

const puzzle4 = '2001\n0..0\n2000\n0..0'
const words4 = ['casa', 'alan', 'ciao', 'anta']
//every assert.throws turns into assert.strictEqual with an error like this one below:

assert.strictEqual(crosswordSolver(puzzle4, words4), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle5 = '0001\n0..0\n3000\n0..0'
const words5 = ['casa', 'alan', 'ciao', 'anta']

assert.strictEqual(crosswordSolver(puzzle5, words5), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle6 = '2001\n0..0\n1000\n0..0'
const words6 = ['casa', 'casa', 'ciao', 'anta']

assert.strictEqual(crosswordSolver(puzzle6, words6), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle7 = ''
const words7 = ['casa', 'alan', 'ciao', 'anta']

assert.strictEqual(crosswordSolver(puzzle7, words7), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle8 = 123
const words8 = ['casa', 'alan', 'ciao', 'anta']

assert.strictEqual(crosswordSolver(puzzle8, words8), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle9 = ''
const words9 = 123

assert.strictEqual(crosswordSolver(puzzle9, words9), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle10 = '2000\n0...\n0...\n0...'
const words10 = ['abba', 'assa']

assert.strictEqual(crosswordSolver(puzzle10, words10), error)
console.log(`Test ${++testCounter} passed.`)

const puzzle11 = '2001\n0..0\n1000\n0..0'
const words11 = ['aaab', 'aaac', 'aaad', 'aaae']

assert.strictEqual(crosswordSolver(puzzle11, words11), error)
console.log(`Test ${++testCounter} passed.`)
