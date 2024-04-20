export const createCircle = () => {
    document.addEventListener('click', (event) => {
      const { pageX, pageY } = event;
      const circle = document.createElement('div');
      circle.classList.add('circle');
      circle.style.background = 'white';
      circle.style.left = `${pageX - 25}px`;
      circle.style.top = `${pageY - 25}px`;
      document.body.appendChild(circle);
    });
  };
  
  export const moveCircle = () => {
    document.addEventListener('mousemove', (event) => {
      const { pageX, pageY } = event;
      const circles = document.querySelectorAll('.circle');
      const lastCircle = circles[circles.length - 1];
      if (lastCircle) {
        const box = document.querySelector('.box');
        const circleRect = lastCircle.getBoundingClientRect();
        const boxRect = box.getBoundingClientRect();
        const isCircleInsideBox =
          circleRect.left > boxRect.left - 1 &&
          circleRect.top > boxRect.top -1 &&
          circleRect.right < boxRect.right + 1 &&
          circleRect.bottom < boxRect.bottom + 1;
        if (isCircleInsideBox) {
            lastCircle.style.background = 'var(--purple)';
        }
        if (lastCircle.style.background !== 'var(--purple)') {
            lastCircle.style.left = `${pageX-25}px`;
          lastCircle.style.top = `${pageY-25}px`;
        } else {
          lastCircle.style.left = `${Math.min(Math.max(pageX-25, boxRect.left), boxRect.right - circleRect.width)}px`;
          lastCircle.style.top = `${Math.min(Math.max(pageY-25, boxRect.top), boxRect.bottom - circleRect.height)}px`;
        }
      }
    });
  };
  
  
  export const setBox = () => {
    const box = document.createElement('div');
    box.classList.add('box');
    document.body.appendChild(box);
  };
  