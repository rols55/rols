import { readdirSync } from 'fs';

const printGuestList = (guests) => {
  guests.forEach((guest, index) => {
    const { firstName, lastName } = guest;
    const lineNumber = index + 1;
    console.log(`${lineNumber}. ${lastName} ${firstName}`);
  });
};

const getGuests = (directoryPath) => {
  const entries = readdirSync(directoryPath);
  const guests = entries
    .map((entry) => {
      const [firstName, lastName] = entry.replace('.json', '').split('_')
      return { lastName,firstName  };
    })
    .sort((a, b) => a.lastName.localeCompare(b.lastName));
  return guests
};

const main = () => {
  const directoryPath = process.argv[2];
  const guests = getGuests(directoryPath);
  printGuestList(guests);
};

main();
