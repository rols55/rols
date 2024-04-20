function format(date, string) {  
    const parts = string.split(/(\b\b)/);
    const formated = parts.map(function(element) {
        return formater(date, element);
    });
    function formater(date, string) {

      switch (string) {
        case 'y':
          return String(Math.abs(date.getFullYear()));
        case 'yyyy':
          return String(Math.abs(date.getFullYear())).padStart(4, '0');
        case 'G':
            return date.getFullYear() > 0 ? 'AD' : 'BC';
        case 'GGGG':
          return date.getFullYear() > 0 ? 'Anno Domini' : 'Before Christ';
        case 'M':
          return String(date.getMonth() + 1);
        case 'MM':
          return String(date.getMonth() + 1).padStart(2, '0');
        case 'MMM':
          return date.toLocaleString('default', { month: 'short' });
        case 'MMMM':
          return date.toLocaleString('default', { month: 'long' });
        case 'd':
          return String(date.getDate());
        case 'dd':
          return String(date.getDate()).padStart(2, '0');
        case 'E':
          return date.toLocaleString('default',{weekday: 'short'});
        case 'EEEE':
          return date.toLocaleString('default',{weekday: 'long'});
        case 'h':
          return String(date.getHours() % 12 || 12);
        case 'hh':
          return String(date.getHours() % 12 || 12).padStart(2, '0');
        case 'm':
          return String(date.getMinutes());
        case 'mm':
          return String(date.getMinutes()).padStart(2, '0');
        case 's':
          return String(date.getSeconds());
        case 'ss':
          return String(date.getSeconds()).padStart(2, '0');
        case 'H':
          return String(date.getHours());
        case 'HH':
          return String(date.getHours()).padStart(2, '0');
        case 'a':
          return date.getHours() < 12 ? 'AM' : 'PM';
        default:
          return string;
        }
    }
    //console.log(formated);
    return formated.join('');  
    
}
/*
const landing = new Date('July 20, 1969, 20:17:40')
const returning = new Date('July 21, 1969, 17:54:12')
const eclipse = new Date(-585, 4, 28)
const ending = new Date('2 September 1945, 9:02:14')

console.log(format(eclipse,'y')) // 585
console.log(format(eclipse, 'yyyy G')) //'0585 BC'
console.log(format(landing, 'yyyy G')) // '1969 AD'
console.log(format(eclipse, 'MMM')) // 'May'
console.log(format(ending, 'MMMM')) //'September'
console.log(format(landing, 'E')) //'Sun'
console.log(format(landing, 'H:m:s')) // '20:17:40'
*/