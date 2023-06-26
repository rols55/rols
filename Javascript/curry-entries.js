// defaultCurry: Merges two objects, with values from the second object overriding the first object
const defaultCurry = obj1 => obj2 => ({ ...obj1, ...obj2 });

// mapCurry: Applies a mapping function to each entry in an object
const mapCurry = fn => obj =>
  Object.entries(obj).reduce((acc, [key, value]) => {
    const [newKey, newValue] = fn([key, value]);
    acc[newKey] = newValue;
    return acc;
  }, {});

// reduceCurry: Applies a reducing function to the entries of an object
const reduceCurry = fn => (obj, initialValue) =>
  Object.entries(obj).reduce(fn, initialValue);

// filterCurry: Filters the entries of an object based on a predicate function
const filterCurry = predicate => obj => {
    const entries = Object.entries(obj).filter(predicate);
    return entries.reduce((acc, [key, value]) => ({ ...acc, [key]: value }), {});
  };
  
/*
// Example usage with personnel object
const personnel = {
  lukeSkywalker: { id: 5, pilotingScore: 98, shootingScore: 56, isForceUser: true },
  sabineWren: { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
  zebOrellios: { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false },
  ezraBridger: { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true },
  calebDume: { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true },
};
*/
// reduceScore: Calculates the total shooting score for force users
const reduceScore = (personnel,initialValue=0) => {
    const forceUsers = filterCurry(([_, value]) => value.isForceUser)(personnel);
    const shooting = reduceCurry((acc, [_, value]) => acc + value.shootingScore)(forceUsers,initialValue);
    const piloting = reduceCurry((acc, [_, value]) => acc + value.pilotingScore)(forceUsers,0);
    return shooting + piloting 
  };
  

// filterForce: Returns force users with shootingScores equal to or higher than 80
const filterForce = filterCurry(([, { shootingScore, isForceUser }]) => {
  return isForceUser && shootingScore >= 80;
});

// mapAverage: Returns a new object with the average score for each person
const mapAverage = mapCurry(([key, value]) => {
  const averageScore = (value.shootingScore + value.pilotingScore)/2;
  return [key, { ...value, averageScore }];
});

// Testing the functions
const mergedObj = defaultCurry(
  {
    http: 403,
    connection: 'close',
    contentType: 'multipart/form-data',
  }
)({
  http: 200,
  connection: 'open',
  requestMethod: 'GET',
});
console.log(mergedObj);

const mappedObj = mapCurry(([k, v]) => [`${k}_force`, v])(personnel);
console.log(mappedObj);

const reducedValue = reduceCurry((acc, [k, v]) => (acc += v))(
  { a: 1, b: 2, c: 3 },
  0
);
console.log(reducedValue);

const filteredObj = filterCurry(
  ([k, v]) => typeof v === 'string' || k === 'arr'
)({
  str: 'string',
  nbr: 1,
  arr: [1, 2],
});
console.log(filteredObj);

const totalScore = reduceScore(personnel);
console.log(totalScore);

const forceUsers = filterForce(personnel);
console.log(forceUsers);

const averageScores = mapAverage(personnel);
console.log(averageScores);
const paramater = 420
console.log(reduceScore(personnel,paramater)) //420