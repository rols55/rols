function fusion(obj1, obj2) {
    // Helper function to check if a value is an object
    const isObject = (value) => typeof value === 'object' && value !== null && !Array.isArray(value);
  
    // Iterate over the keys of the second object
    for (let key in obj2) {
      // Check if the key exists in the first object
      if (obj1.hasOwnProperty(key)) {
        const value1 = obj1[key];
        const value2 = obj2[key];
  
        // Check the types of the values
        if (Array.isArray(value1) && Array.isArray(value2)) {
          // Concatenate arrays
          obj1[key] = value1.concat(value2);
        } else if (typeof value1 === 'string' && typeof value2 === 'string') {
          // Concatenate strings with a space
          obj1[key] = value1 + ' ' + value2;
        } else if (typeof value1 === 'number' && typeof value2 === 'number') {
          // Add numbers
          obj1[key] = value1 + value2;
        } else if (isObject(value1) && isObject(value2)) {
          // Recursively merge objects
          obj1[key] = fusion(value1, value2);
        } else {
          // Replace with the value from the second object for type mismatch
          obj1[key] = value2;
        }
      } else {
        // Add the key-value pair from the second object to the first object
        obj1[key] = obj2[key];
      }
    }
  
    return obj1;
  }
  