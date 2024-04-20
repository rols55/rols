function flat(arr, depth = 1) {
    return depth > 0
      ? arr.reduce((acc, cur) => acc.concat(Array.isArray(cur) ? flat(cur, depth - 1) : cur), [])
      : arr.slice();
  }
  