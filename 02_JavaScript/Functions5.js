/*
  
  Functions in js : functions are a block of code which gets executed whenever we call them.
                     There are various types of functions in hjs such as : 
                     
                     -> normal function
                     -> function with parameter
                     -> function with return 
                     -> default parameter function
                     -> function expression
                     -> Arrow function
                     -> Anonymous function
                     -> IIFE

*/

// normal function: 
function SayHello(){
    console.log("Hello");
    
}

// SayHello()


// function with parameter and arguement:
function MyDetails(name, age){
    console.log(name, age);
    
} 

// MyDetails("rahman", 5049)

// function with return

function Value(marks){
    // console.log(marks);
    return marks
    
}
 
let value = Value(88)
// console.log(value);


// default parameter function : 
function myAge(age = 21){
    console.log(age);
    
}

// myAge()


// function expression : 
let dcet = function(){
    console.log(1);
    
}

// dcet()


// arrow function: 
let myName = ()=>{
    console.log("rahman");
    
}

myName()

let sum =(a, b)=>{
  console.log(a+b);
  
}

sum(2,9)

let strLength = (str)=>{
    console.log(str.length);    
}

strLength("hello")

let myAge1 = (age) => {
    return age
}

let value1 = myAge1(21)
console.log(value1);


// anonymous function : 

let greet = function() {
    console.log("greet");
    
}

greet()

// iife : 
// (function(boy){
//    console.log(boy);
   
//  }("boy"))


// hoisting : hoisitng means calling the fnction even before its written but we cant do this on arrow funcs, func expression.

sayname("rahman")

function sayname(name){
    console.log(name);
    
}


// recursion : a function calling itself inside its own function is called recursion. it is an infinite loop
function rollNo(){
    console.log(5049);
    rollNo()
}

rollNo()