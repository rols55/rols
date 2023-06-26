function dayOfTheYear(date) {
    const startOfYear = new Date(date.getFullYear(), 0, 1);
    const diff = date.getTime() - startOfYear.getTime();
    const oneDay = 1000 * 60 * 60 * 24;
    //if diff is less than one day, return 1
    if (diff < oneDay) {
        return 1;
    }else{
    //if diff is more than one day, return diff/oneDay
    return Math.floor(diff / oneDay) +1
  }
}
console.log(dayOfTheYear(new Date('0001-01-01'))); //1
console.log(dayOfTheYear(new Date('1664-08-09'))); //222
