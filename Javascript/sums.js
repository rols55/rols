function sums(n) {
  const result = [];

  function partition(sum, max, parts) {
    if (sum === n) {
      result.push(parts);
      return;
    }
    for (let i = 1; i <= max && sum + i <= n; i++) {
      partition(sum + i, i, [...parts, i]);
    }
  }

  partition(0, n, []);
  const uniquePartitions = [];
  for (const partition of result) {
    const sortedPartition = partition.slice().sort();
    if (!uniquePartitions.some(p => p.join(',') === sortedPartition.join(','))) {
      uniquePartitions.push(sortedPartition);
    }
  }
  return uniquePartitions.slice(0, uniquePartitions.length - 1);
}
console.log(sums(4)); // [ [1, 1, 1, 1], [1, 1, 2], [1, 3], [2, 2] ]
console.log(sums(2));
console.log(sums(0));