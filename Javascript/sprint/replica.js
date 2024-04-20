const isObject = (value) => typeof value === 'object' && value !== null;

const replica = (target, ...sources) => {
  for (const source of sources) {
    for (const key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        const sourceValue = source[key];
        const targetValue = target[key];

        if (Array.isArray(targetValue) && !Array.isArray(sourceValue)) {
          target[key] = sourceValue; // Replace the entire array in the target with the value from the source
        } else if (Array.isArray(sourceValue) && !Array.isArray(targetValue)) {
          target[key] = sourceValue; // Replace the array in the target with the array from the source
        } else if (isObject(sourceValue)) {
          if (!isObject(targetValue)) {
            target[key] = sourceValue; // Replace the value in the target with the value from the source if the target value is not an object
          } else {
            replica(targetValue, sourceValue); // Recursive call for nested objects
          }
        } else {
          target[key] = sourceValue; // Assign the value directly
        }
      }
    }
  }

  return target;
};



console.log(replica({ a: [1, 2, 4] }, { a: { b: [4] } }).a) // { b: [4] }
/*
console.log(replica({ a: 4 }, { a: { b: 1 } }).a.b) //1))
*/
console.log(replica({ a: { b: [2] } }, { a: [4] }).a) // [4]

console.log(replica({ con: console.log }, { reg: /hello/ }))/* {
  con: console.log,
  reg: /hello/,
})*/