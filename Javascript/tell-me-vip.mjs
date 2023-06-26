

import { readdirSync ,readFile, readFileSync} from 'fs';
import { writeFile} from 'fs/promises';
const printGuestList = async (guests) => {
    const viips = []
  guests.forEach((guest, index) => {
    const { firstName, lastName } = guest;
    const lineNumber = index + 1;
    viips.push(`${lineNumber}. ${lastName} ${firstName}`);
  })
  const write = viips.join('\n')
  try {
    await writeFile('vip.txt', write);
    console.log('Data written to file successfully.');
  } catch (err) {
    console.error('Error writing to file:', err);
  }
  
};

const getGuests = (directoryPath) => {
  const entries = readdirSync(directoryPath);
  const vips = entries.reduce((result, item) => {
  let answer = JSON.parse(readFileSync(directoryPath+'/'+ item, 'utf8')).answer;
    if (answer === 'yes') {
        result.push(item)
    }
    return result;
  }, []);
  const guests = vips
    .map((entry) => {
      const [firstName, lastName] = entry.replace('.json', '').split('_')
      return { lastName,firstName  };
    })
    .sort((a, b) => a.lastName.localeCompare(b.lastName));
  return guests
};

const main = () => {
  const directoryPath = process.argv[2];
  console.log(directoryPath)
  const guests = getGuests(directoryPath);
  printGuestList(guests);
};

main();
