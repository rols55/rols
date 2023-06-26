const pronoun = (str) => {
    const pronouns = ['i', 'you', 'he', 'she', 'it', 'they', 'we'];
    const pronounObject = {};
    const marks = str.replace(/[^\w\s]|,$/g, '');
    const lines = marks.split('\n');
    const sentences = lines.map((line) => line.trim()).filter((line) => line !== '');
  
    for (const sentence of sentences) {
      const words = sentence.split(' ');
  
      for (let i = 0; i < words.length; i++) {
        let word = words[i].toLowerCase();
  
        // Remove punctuation marks and trailing commas from the word
        
  
        if (pronouns.includes(word)) {
          if (!pronounObject[word]) {
            pronounObject[word] = { word: [], count: 0 };
          }
  
          pronounObject[word].count++;
  
          const nextWord = words[i + 1];
          if (nextWord && !pronouns.includes(nextWord.toLowerCase())) {
            pronounObject[word].word.push(nextWord);
          }
        }
      }
    }
  
    return pronounObject;
  };