function letterSpaceNumber(str) {
    const regex = /([a-z])\s(\d(?!\d)(?=[^\w]|$))/gi;
    const result = [];
    let match;
    while ((match = regex.exec(str))) {
      result.push(match[1] + ' ' + match[2]);
    }
    return result;
  }
  
console.log(letterSpaceNumber('This is a 17 sentence with 9 a space in 5it'));