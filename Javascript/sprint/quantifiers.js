function every(array, conditionFn) {
    for (let i = 0; i < array.length; i++) {
      if (!conditionFn(array[i])) {
        return false;
      }
    }
    return true;
  }
  
  function some(array, conditionFn) {
    for (let i = 0; i < array.length; i++) {
      if (conditionFn(array[i])) {
        return true;
      }
    }
    return false;
  }
  
  function none(array, conditionFn) {
    for (let i = 0; i < array.length; i++) {
      if (conditionFn(array[i])) {
        return false;
      }
    }
    return true;
  }
  