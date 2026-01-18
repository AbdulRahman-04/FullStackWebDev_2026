/*
  ES6 (ECMAScript 2015) was introduced in 2015.
  It brought major improvements to JavaScript such as:
  let & const, arrow functions, classes, modules,
  promises, template literals, and destructuring.
  ES6 is considered the foundation of modern JavaScript.
*/

/*
  try...catch is used for error handling in JavaScript.
  It allows us to catch runtime errors and handle them
  without crashing the entire program.
  If an error occurs inside try, control moves to catch.
*/

// try {
//     // const a = [19 
//     const a = [12]

//     console.log(a+b);
        

    
// } catch (error) {
//     console.log(error, error.message);
    
// }


// DESTRUCTURING : 


// object destructuring : 

// object k andar k keys jo names h wahi naam k variables m aap uska data store krskte e.g below : 

let student = {
    name: "arman",
    age: 25,
    isAlve : true
}

const {name, age, isAlve} = student
// console.log(name, age, isAlve);


// arry destructuring: isme aap array k values ku koi b var name m store krskte e.g : 

let arr = ["hey", true, 49, "six"]

const [str, bool, num, str1] = arr

// console.log(str, bool, num, str1);


// For of loop : this loop executes for array values and string values variables only.

let newArr = ["h", "i", "a", "c"]

for(let x of newArr){
    console.log(x);
    
}

let str2 = "hellobhai"

for(let y of str2){
    console.log(y);
    
}

// for in loop : applies only on objects 
let newObj = {
    "key1": true,
    "key2": false
}

// for (let j in newObj){
//     console.log(j, newObj[j]);
    
// }

let newObj1 = {
     "key1": true,
    "key2": false
}

for (let m in newObj1){
    console.log(m , newObj1[m]);
    
}


// set time out and set intervals : used for functions only 

// settimeout : it is used to execute a function after certain amount of time for one.

function myName(name){
    setTimeout(()=>{
        console.log(name);
        
    },5000)
}

// myName("rxhman")


// setinterval : it is used to execute function repeatedly unlike timeout.

function hey(){
    setInterval(()=>{
        console.log("hey");
        
    }, 2000)
}

// hey()

// to control setinterval use clear interval in built method 

function heyya(){
    let stopIt = setInterval(()=>{
        console.log("heyya");
        
    }, 2000)

    setTimeout(()=>{
        clearInterval(stopIt)
    }, 11000)
}

// heyya()


// DATE AND TIME IN JS : 

const now = new Date();

console.log(now.getFullYear()); // year
console.log(now.getMonth());    
console.log(now.getDate());     // day of month

console.log(now.getHours());   
console.log(now.getMinutes());
console.log(now.getSeconds());


// human readable 

console.log(now.toDateString()); // Sun Jan 18 2026
console.log(now.toTimeString()); // 21:35:10 GMT+0530
console.log(now.toLocaleString()); 
