// generateLetters function
export const generateLetters = () => {
    const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const totalLetters = 120
    const letterRange = alphabet.length;
  
    for (let i = 0; i < 120; i++) {
      const letter = document.createElement('div');
      const index = Math.floor(Math.random() * letterRange);
      const selectedLetter = alphabet.charAt(index);
      const fontSize = 11 + i;
      let fontWeight;
  
      if (i < totalLetters / 3) {
        fontWeight = 300;
      } else if (i < (totalLetters / 3) * 2) {
        fontWeight = 400;
      } else {
        fontWeight = 600;
      }
  
      letter.textContent = selectedLetter;
      letter.style.fontSize = `${fontSize}px`;
      letter.style.fontWeight = fontWeight;
      document.body.appendChild(letter);
    }
  };

  