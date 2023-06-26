const args = process.argv[2].split(' ')

for (const words of args) {
  let first = words.slice(0, Math.ceil(words.length / 2));
  let modifiedString = words.replace(first, '') + first;
  process.stdout.write(modifiedString + ' ');
}
