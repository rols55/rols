// Cities Only
function citiesOnly(arr) {
    return arr.map(obj => obj.city);
  }
  
  // Upper Casing States
  function upperCasingStates(arr) {
    return arr.map(str => str.replace(/\b\w/g, c => c.toUpperCase()));
  }
  
  // Fahrenheit to Celsius
  function fahrenheitToCelsius(arr) {
    return arr.map(temp => {
      const fahrenheit = parseFloat(temp);
      const celsius = Math.floor((fahrenheit - 32) * 5 / 9);
      return `${celsius}°C`;
    });
  }
  
// Trim Temp
function trimTemp(arr) {
    return arr.map(obj => ({
        ...obj,
        temperature: obj.temperature.replace(/\s/g, '')
    }));
  }
  
  
  
  // Temp Forecasts
  function tempForecasts(arr) {
    return arr.map(obj => {
      const celsius = fahrenheitToCelsius([obj.temperature])[0];
      const city = upperCasingStates([obj.city])[0];
      const state = upperCasingStates([obj.state])[0];
      return `${celsius}elsius in ${city}, ${state}`;
    });
  }
  /*
  console.log(trimTemp([
    { city: 'Los Angeles', temperature: '  101 °F   ' },
    { city: 'San Francisco', temperature: ' 84 ° F   ' },
  ]))

  console.log(trimTemp([
    {
      city: 'Los Angeles',
      state: 'california',
      region: 'West',
      temperature: '101°F',
    },
    {
      city: 'San Francisco',
      state: 'california',
      region: 'West',
      temperature: '84°F',
    },
    { city: 'Miami', state: 'Florida', region: 'South', temperature: '112°F' },
    {
      city: 'New York City',
      state: 'new york',
      region: 'North East',
      temperature: '0°F',
    },
    { city: 'Juneau', state: 'Alaska', region: 'West', temperature: '21°F' },
    {
      city: 'Boston',
      state: 'massachussetts',
      region: 'North East',
      temperature: '45°F',
    },
    {
      city: 'Jackson',
      state: 'mississippi',
      region: 'South',
      temperature: '70°F',
    },
    { city: 'Utqiagvik', state: 'Alaska', region: 'West', temperature: '-1°F' },
    {
      city: 'Albuquerque',
      state: 'new mexico',
      region: 'West',
      temperature: '95°F',
    },
  ]))
  */