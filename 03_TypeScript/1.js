/* Typescript :
  
 - to install typescript, u need to install node js on ur laptop/pc (node -v)
 - npm i -g typescript (global install, only once)
 - tsc -v

 there are two ways of running ts file

 1) compile the ts file into js and run it after coding is done
    tsc 1.ts -> this complies ts file into js (1.js) -> ndoe 1.js
    
 2) (good way) tsc 1.ts --watch (every time u save code in ts file js file gets automatically updated)
    -> tsc 1.ts --watch -> tsc 1.ts  -> node 1.js

*/
// basic types in typescript : 
// primitive types : number , string, boolean, undefined, null , bigint , symbol.
/* Primitive types in typescript:  there are 3 primitive types in ts, they are: number, string and boolean.

   just like in js these primitive variables are declared.

   in typescript, u can do type annotation, which means telling compiler at time of
   decalring variable its data type

 */
var myName = "rxhman";
var rollNo = 49;
var isAlive = true;
// console.log(myName, typeof(myName));
// console.log(rollNo, typeof(rollNo));
// console.log(isAlive, typeof(isAlive));
// arrays : array in ts is very powerful, we declare type of values which is going to be stored in array.
// normal js array : 
var normal = [1, true, "str"];
console.log(normal);
// ts array : 
var tsArr = ["hi", "hello", "bye"];
// console.log(tsArr);
/*

  Tuples : tuple is a special type of array,
  which stores fixed size and specific data values in itself.

  let arr = [string, boolean, number] = ["hey", true, "go"]

*/
// normal ts array: 
// let normalTs : string[] = ["hi", true]
// console.log(normalTs);
// tuples adv, u can store fix size of elem wit deiff data type in an array
var tupleArr = ["hi", "hi", true, 49];
// console.log(tupleArr);
// ENUM: An enum is a way to store a fixed set of values under a name. 
//       It makes the code clean and easy to read.
// enums se apan custom data type bnaskte with fixed values like in this eg jisme statuscode ek custom datatype h jisme fixed values h error, success etc 
// and jo b variable ku apan enum assign krre wo variable sirf enum k andar k values access krskta..
var StatusCode;
(function (StatusCode) {
    StatusCode[StatusCode["ERROR"] = 404] = "ERROR";
    StatusCode[StatusCode["SUCCESS"] = 200] = "SUCCESS";
    StatusCode[StatusCode["SERVERERROR"] = 404] = "SERVERERROR";
})(StatusCode || (StatusCode = {}));
var codes;
codes = StatusCode.SUCCESS;
// console.log(codes);
// ANY , UNKNOWN , VOID, UNDEFINED, NULL ETC:
// ANY : If a variable is declared without specifying its type and without assigning a value, TypeScript treats it as any. 
//       This means the variable can hold any type of value, which removes TypeScript’s type safety and is generally not recommended.
var variable = "hi";
// console.log(variable, typeof(variable));
// Unknown: it is a special TypeScript type that can store any value, but unlike any,
//         it does not allow operations (like arithmetic or method calls) without proper type checking..
var a = "hello";
a = a.toUpperCase();
// console.log(a);
var b = "hey";
// if (typeof b === "string"){
//     console.log(b);
// }
var num = 12;
if (typeof num === "number") {
    num = num + 10;
    //   console.log(num);
}
var bool = true;
if (typeof bool === "boolean") {
    // console.log(bool);
}
/* Void: it is a special type which tells that what datatype value
        is getting returned inside a function. if no value is getting retrned
        
        void ka matlab hai “koi value return nahi ho rahi”.

        Matlab function sirf execute hota, lekin kuch return nahi karega.
*/
function sayName() {
    // console.log("rxhman");
}
// sayName()
function myClg(college) {
    return college;
}
var value = myClg("dcet");
// console.log(value);
function age(age) {
    console.log(age);
}
// age(21)
// undefined: it literally means a variable is declared but it has been not assigned with any value.
var ba;
// console.log(ba, typeof(ba));
// never: use karte ho jab function kabhi normal tarike se end nahi hoga. 
//        Ya toh error throw karega ya infinite loop me fasa rahega.
function abcd() {
    while (true) { }
}
// abcd()
// console.log("hi");
// type interference and type annotations : 
// type inference: it means when we declare a varaible without defing its type. ts automatically checks the data type of variabke its called 
//                 the type inference
var ab = true;
// console.log(ab, typeof(ab));
//  type annotations: it means when we declare a variable while also defining its data type.
var x = 5049;
var details = {
    name: "rxhman",
    age: 21,
    isAlive: true
};
function myApp(app) {
    console.log(app);
}
var number;
number = 20;
console.log(number);
var myVar;
myVar = true;
var xy = 5049;
var akash;
akash = {
    name: "akash",
    emplyeeid: 5047
};
var myVar1 = {
    id: "123",
    email: 45,
    role: true,
    isAlive: true
};
console.log(myVar1);
