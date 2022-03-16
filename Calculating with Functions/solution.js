function calculate(str) {
    const number1 = str[2]
    const number2 = str[0]
    const operator = str[1]
    let result = 0;

    switch (operator) {
        case '+':
            result = parseInt(number1) + parseInt(number2)
            break;
        case '-':
            result = parseInt(number1) - parseInt(number2)
            break;
        case '/':
            result = Math.floor(parseInt(number1) / parseInt(number2))
            break;
        case '*':
            result = parseInt(number1) * parseInt(number2)
            break;
    }
    return result;
}

function zero(partialAccount) {
    const digit = '0'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function one(partialAccount) {
    const digit = '1'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function two(partialAccount) {
    const digit = '2'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function three(partialAccount) {
    const digit = '3'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function four(partialAccount) {
    const digit = '4'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function five(partialAccount) {
    const digit = '5'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function six(partialAccount) {
    const digit = '6'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function seven(partialAccount) {
    const digit = '7'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function eight(partialAccount) {
    const digit = '8'
    return partialAccount ? calculate(partialAccount + digit) : digit
}
function nine(partialAccount) {
    const digit = '9'
    return partialAccount ? calculate(partialAccount + digit) : digit
}

function plus(partialAccount) {
    const operator = '+'
    return partialAccount + operator
}
function minus(partialAccount) {
    const operator = '-'
    return partialAccount + operator
}
function times(partialAccount) {
    const operator = '*'
    return partialAccount + operator
}
function dividedBy(partialAccount) {
    const operator = '/'
    return partialAccount + operator
}

module.exports.zero = zero;
module.exports.one = one;
module.exports.two = two;
module.exports.three = three;
module.exports.four = four;
module.exports.five = five;
module.exports.six = six;
module.exports.seven = seven;
module.exports.eight = eight;
module.exports.nine = nine;
module.exports.plus = plus;
module.exports.minus = minus;
module.exports.times = times;
module.exports.dividedBy = dividedBy;