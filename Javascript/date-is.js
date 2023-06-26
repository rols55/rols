const isValid = (date) => date === Date.now() ? true : !isNaN(date) && typeof date === 'number' ? true : date instanceof Date && !isNaN(date.getTime())

function isAfter(date1, date2) {
  if (isValid(date1) && isValid(date2)){
    return date1 > date2;
  }
}

  
  function isBefore(date1, date2) {
    if (isValid(date1) && isValid(date2)){
      return date1 < date2
    }
  }
  
  function isFuture(date) {
    const now = new Date();
    return isValid(date) && isAfter(date, now);
  }
  
  function isPast(date) {
    const now = new Date();
    return isValid(date) && isBefore(date, now);
  }