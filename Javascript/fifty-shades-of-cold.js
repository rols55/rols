// generateClasses function
import { colors } from './fifty-shades-of-cold.data.js';

export const generateClasses = () => {
  const head = document.querySelector('head');
  const style = document.createElement('style');

  let classCSS = '';
  for (const color of colors) {
    const className = color.replace(/\s+/g, '-').toLowerCase();
    classCSS += `.${className} { background: ${color}; }\n`;
  }

  style.textContent = classCSS;
  head.appendChild(style);
};

export const generateColdShades = () => {
    const body = document.querySelector('body');
  
    for (let i = 0; i <= colors.length; i++) {
      const color = colors[i];
      if (color) {
        const colorName = color.toLowerCase();
        if (
          colorName.includes('aqua') ||
          colorName.includes('blue') ||
          colorName.includes('turquoise') ||
          colorName.includes('green') ||
          colorName.includes('cyan') ||
          colorName.includes('navy') ||
          colorName.includes('purple')
        ) {
          const className = color.replace(/\s+/g, '-').toLowerCase();
          const div = document.createElement('div');
          div.className = className;
          div.textContent = color;
          body.appendChild(div);
        }
      }
    }
    
  };
  
  
  

// choseShade function
export const choseShade = (shade) => {
  const divs = document.querySelectorAll('div');

  for (const div of divs) {
    div.className = '';
    div.classList.add(shade);
}
}