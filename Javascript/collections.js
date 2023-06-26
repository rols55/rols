const arrToSet = arr => new Set(arr);

const arrToStr = arr => arr.join('');

const setToArr = set => [...set];

const setToStr = set => [...set].join('');

const strToArr = str => str.split('');

const strToSet = str => new Set(str);

const mapToObj = map => {
  const obj = {};
  for (let [key, value] of map) {
    obj[key] = value;
  }
  return obj;
};

const objToArr = obj => Object.values(obj);

const objToMap = obj => new Map(Object.entries(obj));

const arrToObj = arr => {
  const obj = {};
  arr.forEach((value, index) => {
    obj[index] = value;
  });
  return obj;
};

const strToObj = str => {
  const obj = {};
  str.split('').forEach((char, index) => {
    obj[index] = char;
  });
  return obj;
};

const superTypeOf = (val) => {
    const type = typeof val;
    if (type === 'object') {
        if (val === null) return 'null';
        switch (val.constructor) {
          case Array:
            return 'Array';
          case Object:
            return 'Object';
          case Set:
            return 'Set';
          case Map:
            return 'Map';
          default:
            return 'unknown object';
        }
    }else{
        return type === 'undefined'? type : type.charAt(0).toUpperCase() + type.slice(1);
    }
};