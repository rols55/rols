// Multiply without using *
function multiply(a, b) {
    let result = 0;
    const absA = Math.abs(a);
    const absB = Math.abs(b);
    for (let i = 0; i < absB; i++) {
      result += absA;
    }
    if ((a < 0) !== (b < 0)) {
      result = -result;
    }
    return result;
  }
  
  // Divide without using /
  function divide(a, b) {
    let quotient = 0;
    let divisor = b < 0 ? -b : b;
    let dividend = a < 0 ? -a : a;
  
    while (dividend >= divisor) {
      dividend -= divisor;
      quotient++;
    }
  
    return (a < 0) !== (b < 0) ? -quotient : quotient;
  }
  
  // Modulo without using %
  function modulo(a, b) {
    if (b === 0) {
      throw new Error("Cannot divide by zero");
    }
    const absA = Math.abs(a);
    const absB = Math.abs(b);
    let result = absA;
    while (result >= absB) {
      result -= absB;
    }
    if (a < 0) {
      result = -result;
    }
    return result;
  }
  