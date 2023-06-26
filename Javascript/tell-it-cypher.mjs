import { readFile, writeFile } from 'node:fs/promises';
import { argv } from 'node:process'
const guestList = await readFile(argv[2], 'utf8')
if (argv[3] === 'encode') {
    let buff = Buffer.from(guestList);
    let data = buff.toString('base64')
    if (argv[4]) {
        writeFile(`${argv[4]}`, data)
    } else {
        writeFile('cypher.txt', data)
    }
} else if (argv[3] === 'decode') {
    let buff = Buffer.from(guestList, 'base64');
    let data = buff.toString('ascii')
    if (argv[4]) {
        writeFile(`${argv[4]}`, data)
    } else {
        writeFile('clear.txt', data)
    }
}