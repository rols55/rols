import { readdirSync } from 'fs';
import { resolve } from 'path';

const countEntriesInDirectory = (directoryPath) => {
  const entries = readdirSync(directoryPath);
  return entries.length;
};

const getCurrentDirectory = () => process.cwd();

const main = () => {
  const directoryName = process.argv[2] || getCurrentDirectory();
  const directoryPath = resolve(directoryName);
  const entryCount = countEntriesInDirectory(directoryPath);
  console.log(entryCount);
};

main();
