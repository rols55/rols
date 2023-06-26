function groupPrice(str) {
    const regex = /(\S+\d+\.\d+)/g;
    const matches = str.match(regex);
    if (!matches) {
      return [];
    }
    return matches.map(match => [match, match.match(/\d+/g)[0], match.match(/\d+/g)[1]]);
  }
  
  console.log(groupPrice('The price is USD12.31')); // [["USD12.31", "12", "31"]]
  console.log(groupPrice('No price here')); // []
  console.log(groupPrice('The price of the cereals is $4.00. and this1.20'))
  console.log(groupPrice('this, 0.32, is not a match'))