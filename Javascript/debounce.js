// debounce function without leading option
const debounce = (func, delay = 0) => {
    let timeoutId;
  
    return function (...args) {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        func.apply(this, args);
      }, delay);
    };
  };
  
  // opDebounce function with leading option
  function opDebounce(func, wait, leading = false) {
    let timer;
    return function(...args) {
      if (leading && !timer) {
        func.apply(this, args);
      }
      clearTimeout(timer);
      timer = setTimeout(() => {
        if (!leading) {
          func.apply(this, args);
        }
        timer = null;
      }, wait);
    };
  }