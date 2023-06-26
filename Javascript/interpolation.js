const interpolation = ({ step, start, end, callback, duration }) => {
    const distance = end - start; //1
    const stepSize = distance / step; //0.2
    const durationSize = duration / step // 2
    const interpolate = (currentStep) => {
      if (currentStep >= step) return;

      const x = start + currentStep * stepSize
      const point = durationSize + (durationSize * currentStep);
      const interpolatedPoint = [x, point];
  
      setTimeout(() => {
        callback(interpolatedPoint);
        interpolate(currentStep + 1);
      }, duration / step);
    };
  
    interpolate(0);
  };
  