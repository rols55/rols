function longWords(array) {
    return array.every((element) => typeof element === 'string' && element.length >= 5);
  }
  function oneLongWord(array) {
    return array.some((element) => typeof element === 'string' && element.length >= 10);
  }
  function noLongWords(array) {
    return !array.some((element) => typeof element === 'string' && element.length >= 7);
  }
  