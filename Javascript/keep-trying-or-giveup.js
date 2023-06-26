function retry(count, callback) {
  return async function (...args) {
      try {
          const res = await callback(...args);
          return res;
      } catch (error) {
          if (count > 0) {
              return retry(count - 1, callback)(...args);
          } else {
              throw error
          }
      }
  };
}

function timeout(delay, callback) {
  return (...args) => {
      return new Promise((resolve, reject) => {
          const timer = setTimeout(() => {
              reject(new Error('timeout'));
          }, delay);
          callback(...args)
              .then((res) => {
                  clearTimeout(timer);
                  resolve(res);
              })
              .catch(reject);
      });
  };
}