// pimp function
import { styles } from './pimp-my-style.data.js'

let currentIndex = 0

export const pimp = () => {
  const button = document.querySelector('.button');
  
  if (!button.classList.contains('unpimp')) {
    const style = styles[currentIndex];
    button.classList.add(style);
    
    if (currentIndex === styles.length -1) {
      button.classList.add('unpimp');
    }
    currentIndex = currentIndex !== styles.length -1 ? currentIndex + 1 : currentIndex
  } else {
    const style = styles[currentIndex];
    button.classList.remove(style);
    
    if (currentIndex === 0) {
      button.classList.remove('unpimp');
    }
    currentIndex = currentIndex !== 0? currentIndex - 1 : currentIndex
  }
};
