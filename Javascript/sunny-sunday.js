function sunnySunday(date) {
    const firstDay = new Date('0001-01-01');
    const weekdays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
    const theDay = date - firstDay;
    const calc = Math.floor(theDay / (1000 * 60 * 60 * 24))
    return weekdays[calc % 6];
  }
  
  console.log(sunnySunday(new Date('0001-01-07'))) // Monday
  console.log(sunnySunday(new Date('0001-01-01'))) // Monday
  console.log(sunnySunday(new Date('0001-12-01'))) // Friday
  console.log(sunnySunday(new Date('2048-12-08'))) // Thursday