function pyramid(char, height) {
    let result = '';
    const maxChars = height * 2 - 1;
    for (let i = 1; i <= height; i++) {
      const numChars = i * 2 - 1;
      const numSpaces = (maxChars - numChars) / 2 * char.length;
      result += ' '.repeat(numSpaces) + char.repeat(numChars);
      if (i !== height) {
        result += '\n';
      }
    }
    return result;
  }