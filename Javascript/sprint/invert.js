function invert(obj) {
    const invertedObj = {};
    
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        invertedObj[obj[key]] = key;
      }
    }
    
    return invertedObj;
  }