console.log("Hello world");
console.log("mai hu don");
console.log("mai hu john");


// this is one line comment


/*

 this is multi line comment 

*/



// variables : variables in js are like containers which can store data of any value in itself. three ways t ocreate variable

// var : same name variable can be redcleared and value can be updated

var a = 12
var a  = 25
console.log(a);

// let : variable value can be changed but same name variable cannot be redeclared using let
let b = 10
b = 49
// let b = 12
console.log(b);

// const : neither value nor variable name can be redeclared 
const name = "rahman"
// name = "syed"
console.log(name);


// DATATYPES IN JS 

// PRIMTIVE                               // NON PRIMITIVE

// string                                 Object, Arrays and functions.
// number
// boolean
// undefined
// null
// bigint
// symbols



// Type conversion : there are two types of conversion in js. implicit conversion and explicit conversion 

// implicit conversion : use + operator to converts to string data type 

let ab = '50' + 49
console.log(ab, typeof(ab));

let myTrue = true + "true"
console.log(myTrue, typeof(myTrue));


let cd = "49" - 12
console.log(cd, typeof(cd));

let myBool = true - 1
console.log(myBool, typeof(myBool));


// explicit conversions : built in methods for conversions 


// conversion to string : 
let num = 25
let myStr = String(num)
console.log(myStr, typeof(myStr));

let val = false 
let boolVal = String(val)
console.log(boolVal, typeof(boolVal));


// conversion to number
let myStr1 = "5049"
let conv1 = Number(myStr1)
console.log(conv1, typeof(conv1));

let myBool1 = true 
let conv2 = Number(myBool1)
console.log(conv2, typeof(conv2));


//conversion to boolean
let value = "true"
let conv3 = Boolean(value)
console.log(conv3, typeof(conv3));

let number = 1
let conv4 = Boolean(number)
console.log(conv4, typeof(conv4));


// Scoping in js : 

// block scope : a variable declared inside a {} is called blcok scope variable.

//                                    VAR               LET                CONST

// inside of block                    ✅                 ✅                 ✅
// outside of bloc                    ✅                 ❌                  ❌


// function scope : a varable declared inside function is called function scope.

//                                    VAR               LET                CONST

// inside of block                    ✅                 ✅                 ✅
// outside of bloc                    ❌                ❌                  ❌


// global scope : a var declaredd neither inside any block nor function 


//                                    VAR               LET                CONST

// inside of block                    ✅                 ✅                 ✅
// outside of bloc                    ✅                 ✅                 ✅


