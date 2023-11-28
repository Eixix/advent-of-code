const path = "challenge.txt";
const file = Bun.file(path);

const text = await file.text();
const numberList = text.split('\n').map(number => Number(number));

let toReturn = 0;

numberList.forEach(number1 => {
    numberList.forEach(number2 => {
        if (number1 + number2 === 2020) toReturn = number1 * number2
    })
});

console.log(toReturn);