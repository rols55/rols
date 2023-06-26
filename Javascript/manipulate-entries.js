/*const nutritionDB = {
    tomato:  { calories: 18,  protein: 0.9,   carbs: 3.9,   sugar: 2.6, fiber: 1.2, fat: 0.2   },
    vinegar: { calories: 20,  protein: 0.04,  carbs: 0.6,   sugar: 0.4, fiber: 0,   fat: 0     },
    oil:     { calories: 48,  protein: 0,     carbs: 0,     sugar: 123, fiber: 0,   fat: 151   },
    onion:   { calories: 0,   protein: 1,     carbs: 9,     sugar: 0,   fiber: 0,   fat: 0     },
    garlic:  { calories: 149, protein: 6.4,   carbs: 33,    sugar: 1,   fiber: 2.1, fat: 0.5   },
    paprika: { calories: 282, protein: 14.14, carbs: 53.99, sugar: 1,   fiber: 0,   fat: 12.89 },
    sugar:   { calories: 387, protein: 0,     carbs: 100,   sugar: 100, fiber: 0,   fat: 0     },
    orange:  { calories: 49,  protein: 0.9,   carbs: 13,    sugar: 9,   fiber: 0.2, fat: 0.1   },
  }
  const groceriesCart1 = { oil: 500, onion: 230, garlic: 220, paprika: 480 }*/
// filterEntries: filters using both key and value, passed as an array ([k, v])
function filterEntries(cart, predicate) {
    return Object.fromEntries(
      Object.entries(cart).filter(([key, value]) => predicate([key, value]))
    );
  }
  
  // mapEntries: changes the key, the value or both, passed as an array ([k, v])
  function mapEntries(cart, transformer) {
    return Object.fromEntries(
      Object.entries(cart).map(([key, value]) => transformer([key, value]))
    );
  }
  
  // reduceEntries: reduces the entries passing keys and values as an array ([k, v])
  function reduceEntries(cart, reducer, initialValue) {
    return Object.entries(cart).reduce(reducer, initialValue);
  }
  
  // totalCalories: returns the total calories of a cart
  function totalCalories(cart) {
    return Math.round(reduceEntries(
      cart,
      (acc, [key, value]) => acc + value * (nutritionDB[key].calories/100),
      0
    )*10) /10;
  }
  
  // lowCarbs: filters the items of the cart with less than 50 grams of carbs
  function lowCarbs(cart) {
    return filterEntries(
      cart,
      ([key, value]) => value * (nutritionDB[key].carbs/100) < 50
    );
}
function removeTrailingZeros(number) {
    const decimalString = number.toString().split('.')[1] || ''; // Get the decimal part of the number as a string
    const othernumber = number.toString().split('.')[0]; // Get the integer part of the number as a string
    let trimmedDecimal = decimalString.slice(0, 4); // Keep at most four decimal places
  
    for (let i = trimmedDecimal.length; i > 0; i--) {
        if (trimmedDecimal[i] != '0'){
            break;
        }else {
            trimmedDecimal = trimmedDecimal.subslice(0, i);
        }
    
    }
    const trimmedNumber = othernumber + '.' + trimmedDecimal; // Add the integer part and the decimal part of the number to create the
  
    const result = parseFloat(trimmedNumber);
    return isNaN(result) ? 0 : result;
  }
  
  
  // cartTotal: calculates the nutritional facts for each item in the grocery cart
  function cartTotal(cart) {
    return mapEntries(cart, ([key, value]) => {
      const item = nutritionDB[key];
      return [key, {
        calories: removeTrailingZeros(value * item.calories /100),
        protein: removeTrailingZeros(value * item.protein /100),
        carbs: removeTrailingZeros(value * item.carbs /100),
        sugar: removeTrailingZeros(value * item.sugar /100),
        fiber: removeTrailingZeros(value * item.fiber /100),
        fat: removeTrailingZeros(value * item.fat /100),
      }];
    });
  }
  //console.log(lowCarbs(groceriesCart1), { oil: 500, onion: 230 })
