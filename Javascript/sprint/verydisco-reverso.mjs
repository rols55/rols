import { readFile } from 'fs';


// Read the file asynchronously
readFile(process.argv[2], 'utf8', (err, data) => {
  if (err) {
    console.error(err);
    return;
  }
  const args = data.split(' ')

    for (const words of args) {
    let first = words.slice(0, Math.floor(words.length / 2));
    let modifiedString = words.replace(first, '') + first;
    process.stdout.write(modifiedString + ' ');
    }

});

