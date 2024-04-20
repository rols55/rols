function pick(obj, keys) {
    if (!Array.isArray(keys)) {
      keys = [keys];
    }
    
    return Object.fromEntries(Object.entries(obj).filter(([key]) => keys.includes(key)));
  }
  
  function omit(obj, keys) {
    if (!Array.isArray(keys)) {
      keys = [keys];
    }
    
    return Object.fromEntries(Object.entries(obj).filter(([key]) => !keys.includes(key)));
  }
  