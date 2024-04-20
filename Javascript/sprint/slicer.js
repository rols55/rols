function slice(strOrArr, start, end) {
    const len = strOrArr.length;
    start = start < 0 ? Math.max(start + len, 0) : Math.min(start, len);
    end = end === undefined ? len : end < 0 ? Math.max(end + len, 0) : Math.min(end, len);
    const result = [];
    for (let i = start; i < end; i++) {
      result.push(strOrArr[i]);
    }
    return typeof strOrArr === 'string' ? result.join('') : result;
  }