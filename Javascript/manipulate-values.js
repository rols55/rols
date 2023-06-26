function filterValues(cart, predicate) {
    const result = {};
  
    for (let key in cart) {
      if (predicate(cart[key])) {
        result[key] = cart[key];
      }
    }
  
    return result;
  }
  
  function mapValues(cart, transform) {
    const result = {};
  
    for (let key in cart) {
      result[key] = transform(cart[key]);
    }
  
    return result;
  }
  
  function reduceValues(cart, reducer, initialValue = 0) {
    function reduceObject(obj) {
      return Object.values(obj).reduce((acc, val) => {
        if (typeof val === 'object' && !Array.isArray(val)) {
          return reduceObject(val) + acc;
        }
        return reducer(acc, val);
      }, initialValue);
    }
  
    return reduceObject(cart);
  }