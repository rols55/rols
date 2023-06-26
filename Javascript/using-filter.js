function filterShortStateName(arr) {
    return arr.filter(str => str.length < 7);
  }
function filterStartVowel(arr) {
  return arr.filter(str => /^[aeiou]/i.test(str));
}
function filter5Vowels(arr) {
    const vowels = ['a', 'e', 'i', 'o', 'u'];
    return arr.filter(str => {
      const vowelCount = str.split('').filter(char => vowels.includes(char.toLowerCase())).length;
      return vowelCount >= 5;
    });
  }
function filter1DistinctVowel(arr) {
  return arr.filter(str => {
    const distinctVowels = [...new Set(str.toLowerCase().match(/[aeiou]/g))];
    return distinctVowels.length === 1;
  });
}
function multiFilter(arr) {
    return arr.filter(obj =>
      obj.capital.length >= 8 &&
      !/^[aeiou]/i.test(obj.name) &&
      /[aeiou]/i.test(obj.tag) &&
      obj.region !== 'South'
    );
  }
  