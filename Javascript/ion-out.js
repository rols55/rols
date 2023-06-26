function ionOut(str) {
    const matches = str.match(/\w*tion\b/g) || [];
    return matches.map(match => match.replace(/ion$/, ''));
  }