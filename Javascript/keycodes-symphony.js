export const compose = () => {
    const body = document.querySelector('body');
    const notes = [];
  
    const createNote = (key) => {
      const note = document.createElement('div');
      note.classList.add('note');
      note.style.backgroundColor = `#${key.codePointAt(0).toString(16).padStart(6, '0')}`
      console.log(key)
      note.textContent = key;
      body.appendChild(note);
      notes.push(note);
    };
  
    const handleKeyPress = (event) => {
      const { key } = event;
      if (key.match(/^[a-z]$/)) {
        createNote(key.toLowerCase());
      } else if (key === 'Backspace') {
        const lastNote = notes.pop();
        if (lastNote) {
          body.removeChild(lastNote);
        }
      } else if (key === 'Escape') {
        for (const note of notes) {
          body.removeChild(note);
        }
        notes.length = 0;
      }
    };
  
    document.addEventListener('keydown', handleKeyPress);
  };
  