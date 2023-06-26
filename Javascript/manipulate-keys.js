function filterKeys(cart, predicate) {
    return Object.keys(cart).reduce((result, key) => {
      if (predicate(key)) {
        result[key] = cart[key];
      }
      return result;
    }, {});
  }
  
  function mapKeys(cart, transform) {
    return Object.keys(cart).reduce((result, key) => {
      const transformedKey = transform(key);
      result[transformedKey] = cart[key];
      return result;
    }, {});
  }
  
  function reduceKeys(cart, reducer, initialValue) {
    const keys = Object.keys(cart);
  
    if (keys.length === 0) {
      return initialValue || '';
    }
  
    let result = initialValue !== undefined ? initialValue : keys[0];
  
    for (let i = initialValue !== undefined ? 0 : 1; i < keys.length; i++) {
      result = reducer(result, keys[i]);
    }
  
    return result;
  }
  
  
  const nutrients = { carbohydrates: 12, protein: 20, fat: 5 }


console.log(reduceKeys(nutrients, (acc, cr) =>acc.concat(', ', cr)))// output: carbohydrates, protein, fat
const cart = {
    vinegar: 80,
    sugar: 100,
    oil: 50,
    onion: 200,
    garlic: 22,
    paprika: 4,
  }
  console.log(reduceKeys(cart, (acc, cr) => (acc += (cr.length <= 4) & 1), 0)) // output: 1