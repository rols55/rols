const throttle = (func, delay) => {
    let timeoutId;
    let lastExecTime = 0;
  
    return (...args) => {
      const currentTime = Date.now();
      const timeSinceLastExec = currentTime - lastExecTime;
  
      if (timeSinceLastExec >= delay) {
        func.apply(this, args);
        lastExecTime = currentTime;
      } else {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => {
          func.apply(this, args);
          lastExecTime = Date.now();
        }, delay - timeSinceLastExec);
      }
    };
  };

  const opThrottle = (func, delay, options = { trailing: false, leading: false }) => {
    let lastCallTime = 0;
    let timeoutId;
    let shouldExecute = false;
  
    return (...args) => {
      const currentTime = Date.now();
      const elapsedTime = currentTime - lastCallTime;
  
      if (options.leading && elapsedTime >= delay) {
        func.apply(this, args);
        lastCallTime = currentTime;
      }
  
      if (options.trailing) {
        clearTimeout(timeoutId);
  
        if (!timeoutId && !options.leading) {
          func.apply(this, args);
          lastCallTime = currentTime;
        }
  
        shouldExecute = true;
  
        timeoutId = setTimeout(() => {
          if (shouldExecute) {
            func.apply(this, args);
            lastCallTime = currentTime;
            shouldExecute = false;
          }
          timeoutId = null;
        }, delay);
      }
    };
  };
  