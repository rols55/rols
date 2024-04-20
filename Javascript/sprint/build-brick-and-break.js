// build function
export const build = (amount) => {
    let count = 1;
    const interval = setInterval(() => {
      const brick = document.createElement('div');
      brick.id = `brick-${count}`;
      brick.classList.add('brick');
      if (count % 3 === 2) {
        brick.dataset.foundation = 'true';
      }
      document.body.appendChild(brick);
      count++;
      if (count > amount) {
        clearInterval(interval);
      }
    }, 100);
  };
  
  // repair function
  export const repair = (...ids) => {
    ids.forEach(id => {
      const brick = document.getElementById(id);
      if (brick) {
        if (brick.dataset.foundation === 'true') {
          brick.dataset.repaired = 'in progress';
        } else {
          brick.dataset.repaired = 'true';
        }
      }
    });
  };
  
  // destroy function
  export const destroy = () => {
    const bricks = document.querySelectorAll('.brick');
    const lastBrick = bricks[bricks.length - 1];
    if (lastBrick) {
      lastBrick.remove();
    }
  };
  