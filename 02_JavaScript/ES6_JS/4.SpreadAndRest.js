/*
=====================================
SPREAD & REST OPERATOR (ES6)
=====================================
*/

/*
-------------------------------------
SPREAD OPERATOR (...)
-------------------------------------

Definition:
The spread operator expands an array or object
into individual elements or properties.
It is used to copy, merge, or pass values.
*/


// 1️⃣ ARRAY EXPANSION (COPY / MERGE)

let arr1 = [1, 2, 3, 4];

// copy array
let arr2 = [...arr1];

console.log(arr1, arr1.length);
console.log(arr2, arr2.length);

// merge arrays
let arr3 = [5, 6];
let mergedArr = [...arr1, ...arr3];

console.log(mergedArr);


// 2️⃣ OBJECT EXPANSION (COPY / ADD / UPDATE)

let a = {
  x: "x",
  y: "y"
};

// copy and add new key
let b = { ...a, c: 3 };
console.log(b);

// copy and update value
let c = { ...a, y: "updated" };
console.log(c);


// 3️⃣ FUNCTION ARGUMENTS (SPREAD)

function sum(a, b, c) {
  console.log(a + b + c);
}

const nums = [1, 2, 3];
sum(...nums);


/*
-------------------------------------
REST OPERATOR (...)
-------------------------------------

Definition:
The rest operator collects multiple values
into a single array.
It is mostly used in function parameters.
*/


// 4️⃣ FUNCTION PARAMETERS (REST)

function addAll(...numbers) {
  console.log(numbers);
}

addAll(1, 2, 3, 4, 5);


// 5️⃣ REST WITH CALCULATION

function total(...nums) {
  let sum = 0;

  for (let n of nums) {
    sum += n;
  }

  console.log(sum);
}

total(10, 20, 30);


// 6️⃣ REST WITH NORMAL PARAMETERS

function showUser(name, ...skills) {
  console.log("Name:", name);
  console.log("Skills:", skills);
}

showUser("Rahman", "JS", "React", "Node");


/*
-------------------------------------
SPREAD vs REST (SHORT)
-------------------------------------

Spread → breaks values apart
Rest   → collects values together

*/
