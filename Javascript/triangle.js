function triangle(char, height) {
    let output = '';
    for (let i = 1; i <= height; i++) {
      output += char.repeat(i) + (i < height ? '\n' : '');
    }
    return output;
  }
  
console.log(triangle("*",4))