function filter(array, callback) {
    const result = [];
    for (let i = 0; i < array.length; i++) {
      if (callback(array[i], i, array)) {
        result.push(array[i]);
      }
    }
    return result;
  }
  
  function reject(array, callback) {
    return filter(array, (element, index, arr) => !callback(element, index, arr));
  }
  
  function partition(array, callback) {
    const passed = [];
    const failed = [];
    for (let i = 0; i < array.length; i++) {
      const element = array[i];
      if (callback(element, i, array)) {
        passed.push(element);
      } else {
        failed.push(element);
      }
    }
    return [passed, failed];
  }
  