const isFriday = Date => Date.getDay() === 5 ? true : false;
const isWeekend = Date => Date.getDay() === 6 || 0 ? true : false;
// Returns true if year of the date is a leap year
function isLeapYear(date) {
    const year = date.getFullYear();
    return (year % 4 === 0 && year % 100 !== 0) || year % 400 === 0;
  }
  
  // Returns true if the date is the last day of the month
  function isLastDayOfMonth(date) {
    const month = date.getMonth();
    const year = date.getFullYear();
    const lastDay = new Date(year, month + 1, 0).getDate();
    return date.getDate() === lastDay;
  }
  