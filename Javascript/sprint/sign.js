const sign = number => number === 0? 0 : number > 0? 1 : -1
const sameSign = (a, b) => sign(a) === sign(b)