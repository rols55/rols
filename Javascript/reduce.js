function fold(array, callback, accumulator) {
    let result = accumulator;
    for (let i = 0; i < array.length; i++) {
      result = callback(result, array[i]);
    }
    return result;
  }
  
  function foldRight(array, callback, accumulator) {
    let result = accumulator;
    for (let i = array.length - 1; i >= 0; i--) {
      result = callback(result, array[i]);
    }
    return result;
  }
  
  function reduce(array, callback) {
    if (array.length < 1) {
      throw new Error('Array must have at least one element.');
    }
    let result = array[0];
    for (let i = 1; i < array.length; i++) {
      result = callback(result, array[i]);
    }
    return result;
  }
  
  function reduceRight(array, callback) {
    if (array.length < 1) {
      throw new Error('Array must have at least one element.');
    }
    let result = array[array.length - 1];
    for (let i = array.length - 2; i >= 0; i--) {
      result = callback(result, array[i]);
    }
    return result;
  }
  