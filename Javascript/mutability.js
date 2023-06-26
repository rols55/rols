// Define clone1 and clone2 constants by creating a new object with the spread operator
const clone1 = { ...person }
const clone2 = { ...person }

// Define samePerson as a reference to the original person object
const samePerson = person

// Modify person by increasing age by one and changing country to 'FR'
person.age++
person.country = 'FR'