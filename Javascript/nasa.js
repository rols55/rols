function nasa(N) {
    let result = "";
    for (let i = 1; i <= N; i++) {
      switch (true) {
        case i % 3 === 0 && i % 5 === 0:
          result += "NASA";
          break;
        case i % 3 === 0:
          result += "NA";
          break;
        case i % 5 === 0:
          result += "SA";
          break;
        default:
          result += i.toString();
      }
      if (i !== N) {
        result += " ";
      }
    }
    return result;
  }
  