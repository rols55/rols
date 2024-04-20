function matchCron(cron, date) {
    const fields = cron.split(' ');
  
    const minute = date.getMinutes();
    const hour = date.getHours();
    const dayOfMonth = date.getDate();
    const month = date.getMonth() + 1; // January is 0 in JS
    const dayOfWeek = date.getDay() || 7; // Sunday is 0 in JS, 7 in cron
  
    return (
      matchField(fields[0], minute) &&
      matchField(fields[1], hour) &&
      matchField(fields[2], dayOfMonth) &&
      matchField(fields[3], month) &&
      matchField(fields[4], dayOfWeek)
    );
  }
  
  function matchField(field, value) {
    if (field === '*') {
      return true;
    }
  
    const values = field.split(',');
  
    for (let i = 0; i < values.length; i++) {
      const currentValue = values[i];
  
      if (currentValue.includes('-')) {
        const [start, end] = currentValue.split('-').map(Number);
        if (value >= start && value <= end) {
          return true;
        }
      } else if (currentValue.includes('/')) {
        const [start, interval] = currentValue.split('/').map(Number);
        if ((value - start) % interval === 0) {
          return true;
        }
      } else if (Number(currentValue) === value) {
        return true;
      }
    }
  
    return false;
  }
  