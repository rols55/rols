const isPositive = num => num > 0
const abs = num => (num === 0 ? 0 : isPositive(num) ? num : -num)