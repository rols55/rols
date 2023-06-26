function adder(numbers, initalValue = 0) {
    return numbers.reduce((sum, value) => sum + value, initalValue);
  }
  
function sumOrMul(numbers, initalValue = 0) {
  return numbers.reduce((result, num) => {
    if (num % 2 === 0) {
      return result * num;
    } else {
      return result + num;
    }
  }, initalValue)
}
function funcExec(functions, initialValue) {
    return functions.reduce((result, func) => func(result), initialValue);
  }
  
//console.log(adder([9, 24, 7, 11, 3], 10)) // 64
//console.log(sumOrMul([29, 23, 3, 2, 25])) // 135
//console.log(sumOrMul([18, 17, 7, 13, 25], 12)) // 278