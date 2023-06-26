function get(src, path) {
    return path.split('.').reduce((acc, key) => acc ? acc[key] : undefined, src);
  }
  