
function crosswordSolver(crossword, words) {
    console.log()
    if (!Array.isArray(words) || typeof crossword !== 'string') { console.log('Error'); return "Error";}
    for (let word of words) {if ((typeof word)!=='string') { console.log('Error'); return "Error";} }
    if (crossword.length === 0 || words.length === 0) { console.log('Error'); return "Error";}
    let crosswordNet = readCrosswordNet(crossword);
    if (!crosswordNet) { console.log('Error'); return "Error";}
    //sum of puzzle starting cells must be less or equal to sum of rows & collumns
    if (crosswordNet.sumOfPlacesForWord !== words.length) { console.log('Error'); return "Error";}
    if (crosswordNet.startingCellsCoordinates.length !== words.length) { console.log('Error'); return "Error";}
    crosswordNet.checkAcross = checkAcross;
    crosswordNet.checkDown = checkDown;
    crosswordNet.searchPlaceAndFillIn = searchPlaceAndFillIn;
    crosswordNet.delWord = delWord;

    // words must be unique
    for (let i = 0; i < words.length; i++) {
        if (words[i].length < 2) { console.log('Error'); return "Error";}// a word must have at least 2 letters
        for (let j = i + 1; j < words.length; j++) {
            if (words[i] === words[j]) {
                console.log('Error');
                return "Error";
            }
        }

    }

    const solutions = crosswordSearchSolutions(crosswordNet, words);

    if (solutions.length !== 1) { console.log('Error'); return "Error";}
    let solution = solutions[0][0].join('');
    for (let i = 1; i < solutions[0].length; i++) {
        solution = solution.concat('\n', solutions[0][i].join(''));
    }
    console.log(solution)
    
    return solution;
}

function crosswordSearchSolutions(crossword, words) {
    let solutions = [];
    if (words.length === 0) {
        solutions.push(copyCrossword(crossword));
        return solutions;
    }

    const currentWord = words.pop();
    let startCellNumber = 0;
    while (startCellNumber < crossword.startingCellsCoordinates.length) {
        const place = crossword.searchPlaceAndFillIn(currentWord, startCellNumber);

        if (place) {
            let solution = crosswordSearchSolutions(crossword, words);
            if (solution.length !== 0) {
                solutions.push(...solution);
            }
            crossword.delWord(place);
            startCellNumber = place.startCellNumber + 1;
        } else {
            words.push(currentWord);
            return solutions;
        }
    }
    words.push(currentWord);
    return solutions;
}

// returns an object of the crossword's net
function readCrosswordNet(str) {
    if (!str) return undefined;
    let symbols = str.split('\n').map(row => row.split(''));
    let crosswordNet = [];
    let startingCellsCoordinates = [];
    let sumOfPlacesForWord = 0;
    const acrossLen = symbols[0].length;
    const downLen = symbols.length;

    for (let i = 0; i < downLen; i++) {
        if (symbols[i].length !== acrossLen) {
            return undefined // rows have different lengths
        }
        // create an across row for our crosswordNet
        let across = [];
        for (let j = 0; j < acrossLen; j++) {
            if (symbols[i][j] === '.') {
                across.push({
                    letter: '.',
                });
            } else {
                let number = +symbols[i][j];
                if (isNaN(number) || number < 0 || number > 2) return undefined; // invalid symbol

                across.push({
                    number: number,
                    letter: '',
                });
                // save the starting position
                if (number !== 0) {
                    sumOfPlacesForWord += number;
                    // check if it's across
                    if (j < acrossLen - 1 && symbols[i][j + 1] !== '.' && (j === 0 || (symbols[i][j - 1] === '.'))) {
                        startingCellsCoordinates.push({ i: i, j: j, direction: 'across' });
                        number--;
                    }
                    // check if it's down
                    if (i < downLen - 1 && symbols[i + 1][j] !== '.' && (i === 0 || (symbols[i - 1][j] === '.'))) {
                        startingCellsCoordinates.push({ i: i, j: j, direction: 'down' });
                        number--;
                    }
                    if (number < 0) { // therre were across and down places, but number was only 1
                        return undefined;
                    }
                }
            }
        }
        crosswordNet.push(across);
    }
    crosswordNet.acrossLen = acrossLen;
    crosswordNet.downLen = downLen;
    crosswordNet.startingCellsCoordinates = startingCellsCoordinates;
    crosswordNet.sumOfPlacesForWord = sumOfPlacesForWord;
    return crosswordNet;
}

// returns true if it is possible to put the word to the crossword across starting at the given coordinates.
// and false in the other case.
function checkAcross(word, startCellNumber) {
    const iStart = this.startingCellsCoordinates[startCellNumber].i;
    let j = this.startingCellsCoordinates[startCellNumber].j;
    for (let letter of word) {
        if (j >= this.acrossLen || this[iStart][j].letter === '.') { return false; } // the place is less than the word
        if (this[iStart][j].letter === '' || this[iStart][j].letter === letter) {
            j++;
        } else {
            return false;
        };
    }
    if (j < this.acrossLen && this[iStart][j].letter !== '.') { return false; } // the place is bigger than the word
    return true;
}

function checkDown(word, startCellNumber) {
    let i = this.startingCellsCoordinates[startCellNumber].i;
    const jStart = this.startingCellsCoordinates[startCellNumber].j;
    for (let letter of word) {
        if (i >= this.downLen || this[i][jStart].letter === '.') { return false; } // the place is less than the word
        if (this[i][jStart].letter === '' || this[i][jStart].letter === letter) {
            i++;
        } else {
            return false;
        };
    }
    if (i < this.downLen && this[i][jStart].letter !== '.') { return false; } // the place is bigger than the word
    return true;
}

// returns a place for the word or 
// undefined if it was impossible to put the word to any places
function searchPlaceAndFillIn(word, startCellNumber) {
    let place = [];
    for (let s = startCellNumber; s < this.startingCellsCoordinates.length; s++) {
        const iStart = this.startingCellsCoordinates[s].i;
        const jStart = this.startingCellsCoordinates[s].j;
        if (this.startingCellsCoordinates[s].direction === 'across' && this.checkAcross(word, s)) {
            let k = jStart;
            for (let letter of word) {
                place.push({ i: iStart, j: k, letter: this[iStart][k].letter });
                this[iStart][k].letter = letter;
                k++;
            }
            place.startCellNumber = s;
            this[iStart][jStart].number--;
            return place;
        }
        if (this.startingCellsCoordinates[s].direction === 'down' && this.checkDown(word, s)) {
            let k = iStart;
            for (let letter of word) {
                place.push({ i: k, j: jStart, letter: this[k][jStart].letter });
                this[k][jStart].letter = letter;
                k++;
            }
            place.startCellNumber = s;
            this[iStart][jStart].number--;
            return place;
        }
    }
    return undefined;
}

// deletes a word from the given place in the crossword
function delWord(place) {
    // console.log('delWord');
    // console.log(this);
    // console.log(place);
    // console.log(place.i,place.j);
    this[place[0].i][place[0].j].number++;
    for (let cell of place) {
        this[cell.i][cell.j].letter = cell.letter;
    }

}

function copyCrossword(crossword) {
    let crosswordCopy = [];
    for (let i = 0; i < crossword.downLen; i++) {
        let row = [];
        for (let j = 0; j < crossword.acrossLen; j++) {
            row.push(crossword[i][j].letter);
        }
        crosswordCopy.push(row);
    }
    return crosswordCopy;
}
/**/
module.exports = crosswordSolver;