const WEEK_LENGTH = 14;
const WEEKDAYS = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday', 'secondMonday', 'secondTuesday', 'secondWednesday', 'secondThursday', 'secondFriday', 'secondSaturday', 'secondSunday'];

function addWeek(date) {
  const epoch = new Date('0001-01-01');
  const daysSinceEpoch = Math.floor((date - epoch) / (1000 * 60 * 60 * 24));
  const weekDayIndex = daysSinceEpoch % WEEK_LENGTH;
  return WEEKDAYS[weekDayIndex];
}

function timeTravel({ date, hour, minute, second }) {
  const newDate = new Date(date);
  newDate.setHours(hour);
  newDate.setMinutes(minute);
  newDate.setSeconds(second);
  return newDate;
}
