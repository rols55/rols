const indexOf = (arr, val, fromIndex = 0) => {
    for (let i = fromIndex; i < arr.length; i++) {
      if (arr[i] === val) {
        return i;
      }
    }
    return -1;
  };
  
  const lastIndexOf = (arr, val, fromIndex = arr.length - 1) => {
    for (let i = fromIndex; i >= 0; i--) {
      if (arr[i] === val) {
        return i;
      }
    }
    return -1;
  };
  
  const includes = (arr, val) => {
    for (let i = 0; i < arr.length; i++) {
      if (arr[i] === val) {
        return true;
      }
    }
    return false;
  };
  