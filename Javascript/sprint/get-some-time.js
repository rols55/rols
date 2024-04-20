function firstDayWeek(week, year) {
    let startWeek = new Date(year);
    if (week !== 1) {
    startWeek.setDate(startWeek.getDate() + (week - 1) * 7);
    const toMonday = startWeek.getDay() === 0? 6 : startWeek.getDay() - 1;
    startWeek.setDate(startWeek.getDate() - toMonday);
    }
   
    
    const month = (startWeek.getMonth() + 1).toString().padStart(2, '0');
    const day = startWeek.getDate().toString().padStart(2, '0');
    return `${day}-${month}-${year}`;
    
  }

  
console.log(firstDayWeek(1, '1000')) //'01-01-1000'
console.log(firstDayWeek(52, '1000')) //'22-12-1000'