function split(str, delimiter) {
    const result = [];
    let start = 0;
  
    if (delimiter === '') {
      for (let i = 0; i < str.length; i++) {
        result.push(str[i]);
      }
    } else {
      while (true) {
        const index = str.indexOf(delimiter, start);
  
        if (index === -1) {
          result.push(str.slice(start));
          break;
        }
  
        result.push(str.slice(start, index));
        start = index + delimiter.length;
      }
    }
  
    return result;
  }
  

  function join(arr, separator) {
    let result = '';
    for (let i = 0; i < arr.length; i++) {
      result += arr[i];
      if (i !== arr.length - 1) {
        result += separator;
      }
    }
    return result;
  }