function deepCopy(obj) {
    if (Array.isArray(obj)) {
      return obj.map(deepCopy);
    } else if (obj instanceof RegExp) {
      return new RegExp(obj);
    } else if (typeof obj === 'object' && obj !== null) {
      const result = Array.isArray(obj) ? [] : {};
      for (const key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
          result[key] = deepCopy(obj[key]);
        }
      }
      return result;
    } else {
      return obj;
    }
  }
  