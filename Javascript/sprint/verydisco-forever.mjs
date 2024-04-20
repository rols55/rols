import { writeFile } from 'fs';
const args = process.argv[2].split(' ')
let writer = []
for (const words of args) {
  let first = words.slice(0, Math.ceil(words.length / 2));
  let modifiedString = words.replace(first, '') + first;
  writer.push(modifiedString);
}
const toWrite = writer.join(' ')

writeFile('verydisco-forever.txt' , toWrite, (err) => {
    if (err) {
      console.error('Error writing to file:', err);
      return;
    }
    console.log('Data written to file successfully.');
  });