function ceilHelper(num) {
    
    if (num-trunc(num) === 0 || num < 0) {
        return (trunc(num))
    } else {

        return (trunc(num) +1)
    }

}

function floorHelper(num) {
    
    if (num-trunc(num) === 0 || num > 0) {
        return (trunc(num))
    } else {

        return (trunc(num)) -1
    }

}


  function roundHelper(num) {
    let decimalPart = num > 0 ? num - trunc(num) : trunc(num) - num;

    if (num >= 0) {
        return decimalPart < 0.5 ? floor(num) : ceil(num);
    } else {
        return decimalPart < 0.5 ? ceil(num) : floor(num);
    }
}

function truncHelper(num) {
    let isNegative = false;
    let n = num;
    
    if (n < 0) {
        n = -n;
        isNegative = true;
    }
    
    let exponent = 0;
    while (Math.pow(10, exponent) <= n) {
        exponent++;
    }
    exponent--;

    let mostSignificantPart = 0;
    while (exponent >= 0) {
        let power = Math.pow(10, exponent);
        while (n >= power) {
            n -= power;
            mostSignificantPart += 1;
        }
        if (exponent > 0) {
            mostSignificantPart *= 10;
        }
        exponent--;
    }
    
    if (isNegative) {
        mostSignificantPart = -mostSignificantPart;
    }
    
    return mostSignificantPart;
}



const round = (num) => roundHelper(num);
const floor = (num) => floorHelper(num);
const trunc = (num) => truncHelper(num);
const ceil  = (num) => ceilHelper(num);