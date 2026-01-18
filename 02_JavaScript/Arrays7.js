/*
  Array : array is a data structure which tsores multiple values in itself like objects buy not in key value pair
          and stores values in ordered form.

*/

let arr = [true, 45, "string"]
// console.log(arr[0]);

arr[1] = 49
// console.log(arr);

delete arr[1]
// console.log(arr);

// Array Methods: 

// .map() is an array method that takes each item in an array, changes it in some way, and returns a new array with those changed values.

// It does not modify the original array

// It always returns a new array

// You use it when you want to apply the same logic to every element

const nums = [2, 4, 6, 8];

const square = nums.map((n)=>{
    return n**2
})

// console.log(square);


const names = ["rahman", "ali", "zoya"];

const case1 = names.map((n)=>{
    return n.toUpperCase()
})

// console.log(case1);


const users = [
  { id: 1, name: "Rahman" },
  { id: 2, name: "Ayaan" },
  { id: 3, name: "Sara" }
];

const namess = users.map((n)=>{
    return n.name
})
// console.log(namess);



// .forEach() goes through each element of an array and runs some code on it.

// It does not return anything

// It is used for doing actions, not creating new arrays

// Original array is not changed automatically (unless you change it yourself)


const x = [1,34, 45, 66, 90]

const get1 = x.forEach((x)=>{
    console.log(x*2);
    
})

console.log(get1);




// filter : .filter() is an array method that checks each element and keeps only the ones that pass a condition.
// It returns a new array with the elements that match the rule.

// Original array is not changed

// Callback must return true or false

// true → keep the element

// false → remove the element

const nums1 = [1,2,3,4,5,6,7,8,9,10]

const even = nums1.filter((n)=>{
    return n%2==0
})

console.log(even);

const num2 = [1,2,3,4,5,6,7,8,9,10]

const odd = num2.filter((x)=>{
    return x%2!=0
})

console.log(odd);

// .find() is an array method that searches for the first element that matches a condition and returns that element.

// It stops as soon as it finds a match

// It returns one value, not an array

// If nothing matches → returns undefined

// Original array is not changed


const num3 = [1,2,3,4,5,6,7,8,9,10]

const finding = num3.find((x)=>{
    return x>5
})

console.log(finding);

const getVal = [true, true , true, false, false]

const finding2 = getVal.find((x)=>{
    return x === false
})

console.log(finding2);


// .some() is an array method that checks if at least one element in the array satisfies a condition.

// It returns true or false

// Stops early once it finds a match

// Does not return an element or array

// Original array is not changed


const num4 = [1,2,3,4,5]

const getVal1 = num4.some((x)=>{
    return x%5==0
})

console.log(getVal1);


// .every() is an array method that checks whether all elements in the array satisfy a condition.

// It returns true or false

// If even one element fails → returns false

// Stops early on first failure

// Original array is not changed

const nums5 = [4,6,10]

const getVal2 = nums5.every((x)=>{
    return x%2==0
})

console.log(getVal2);


// .includes() is an array (and string) method that checks whether a specific value exists in the array.

// It returns true or false

// Uses value comparison, not a condition function

// Original array is not changed

const getMyNames = ["rahman", "omer", "sneha", "ismail"]

const getVal5 = getMyNames.includes("omer")

console.log(getVal5);


// .push() adds one or more elements to the end of an array and changes the original array.

// Returns the new length of the array

const nos = [48,49,50,51,52]

nos.push(53, 54, 55)

// console.log(newArr);
console.log(nos);


// .pop() removes the last element from an array and changes the original array.

// Returns the removed element

// Mutates the array

nos.pop()
console.log(nos);
nos.pop()
console.log(nos);


// .slice() cuts out a part of an array and gives it to you in a new array.

// Original array stays the same

// You choose from where to where to cut

// The cut part comes in a new array

const numbers = [1,2,3,4,5,6,7,8,9,10]

const getMyNums = numbers.slice(4,10)

console.log(getMyNums);


// .sort() arranges the elements of an array in order (like ascending or alphabetical) and changes the original array.

// It modifies the array

// By default, it sorts as strings

// You can give a function to control the order


const getall = ["rxhman", "abd", "ismail", "suhail"]

const sorted = getall.sort()

console.log(sorted);



// .splice() array ke beech se items nikaalta hai ya add karta hai
// aur original array ko change karta hai.

const arr2 = [1, 2, 3, 4];

arr2.splice(1, 2, 99, 100);

console.log(arr2);



// .reduce() is an array method that takes all elements of an array and combines them into one final value.

// The final value can be a number, string, object, or array

// It runs element by element

// Original array is not changed

const nums7 = [10, 20, 30];

// sum = 0 , n = 10, then 20 , then 30 

const total = nums7.reduce((sum, n) => {
  return sum + n;
}, 0);
