function findExpression(num) {
    const queue = [{ value: 1, expression: '1' }];
    
    while (queue.length > 0) {
      const { value, expression } = queue.shift();
      
      if (value === num) {
        return expression;
      }
      
      if (value * 2 <= num) {
        queue.push({
          value: value * 2,
          expression: `${expression} ${mul2}`
        });
      }
      
      if (value + 4 <= num) {
        queue.push({
          value: value + 4,
          expression: `${expression} ${add4}`
        });
      }
    }
    
    return undefined;
  }
  