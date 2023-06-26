const reverse = input => {
    if (typeof input === "string") {
      return input.split("").reduceRight((acc, val) => acc + val, "");
    } else {
      return input.reduceRight((acc, val) => [...acc, val], []);
    }
  };