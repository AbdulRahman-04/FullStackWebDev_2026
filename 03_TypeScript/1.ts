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

let myName : string = "rxhman"
let rollNo : number = 49
let isAlive : boolean = true

// console.log(myName, typeof(myName));
// console.log(rollNo, typeof(rollNo));
// console.log(isAlive, typeof(isAlive));


// arrays : array in ts is very powerful, we declare type of values which is going to be stored in array.

// normal js array : 

let normal = [1, true, "str"]
console.log(normal);

// ts array : 
let tsArr : string[] = ["hi", "hello", "bye"]
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
let tupleArr : [string, string, boolean, number] = ["hi", "hi", true, 49]
// console.log(tupleArr);


// ENUM: An enum is a way to store a fixed set of values under a name. 
//       It makes the code clean and easy to read.

// enums se apan custom data type bnaskte with fixed values like in this eg jisme statuscode ek custom datatype h jisme fixed values h error, success etc 
// and jo b variable ku apan enum assign krre wo variable sirf enum k andar k values access krskta..

enum StatusCode {
    ERROR = 404,
    SUCCESS = 200,
    SERVERERROR = 404
}

let codes : StatusCode;

codes = StatusCode.SUCCESS
// console.log(codes);




// ANY , UNKNOWN , VOID, UNDEFINED, NULL ETC:

// ANY : If a variable is declared without specifying its type and without assigning a value, TypeScript treats it as any. 
//       This means the variable can hold any type of value, which removes TypeScript’s type safety and is generally not recommended.

let variable:any = "hi"
// console.log(variable, typeof(variable));



// Unknown: it is a special TypeScript type that can store any value, but unlike any,
//         it does not allow operations (like arithmetic or method calls) without proper type checking..

let a:any = "hello"
a = a.toUpperCase()
// console.log(a);

let b : unknown = "hey"

// if (typeof b === "string"){
//     console.log(b);
    
// }


let num : unknown = 12 

if (typeof num === "number"){
  num = num + 10
//   console.log(num);
  
}

let bool :  unknown = true

if(typeof bool === "boolean"){
    // console.log(bool);
    
}



/* Void: it is a special type which tells that what datatype value 
        is getting returned inside a function. if no value is getting retrned 
        
        void ka matlab hai “koi value return nahi ho rahi”.

        Matlab function sirf execute hota, lekin kuch return nahi karega.        
*/

function sayName() : void {
    // console.log("rxhman");
    
}

// sayName()

function myClg(college: string): string{
    return college
}

let value : string = myClg("dcet")
// console.log(value);

function age(age: number) : void{
    console.log(age);
    
}

// age(21)



// undefined: it literally means a variable is declared but it has been not assigned with any value.

let ba:undefined;
// console.log(ba, typeof(ba));



// never: use karte ho jab function kabhi normal tarike se end nahi hoga. 
//        Ya toh error throw karega ya infinite loop me fasa rahega.

function abcd():never{
    while(true){}
}

// abcd()

// console.log("hi");


// type interference and type annotations : 


// type inference: it means when we declare a varaible without defing its type. ts automatically checks the data type of variabke its called 
//                 the type inference

let ab = true
// console.log(ab, typeof(ab));

//  type annotations: it means when we declare a variable while also defining its data type.
let x : number = 5049
// console.log(x, typeof(x));


// interface: an interface is like a rulebox used for objects , as it tells what properties an object must have and also their types.

interface myObjDetails {
    name: string,
    age: number,
    isAlive: boolean
}

let details: myObjDetails = {
    name : "rxhman",
    age: 21,
    isAlive: true
}

// console.log(details);

interface StatusCodes {
    errors : number,
    succes : number,
    appAlive: boolean
}

function myApp(app: StatusCodes):void{
    
    console.log(app);
    
}

// myApp({
//         errors: 400,
//         succes : 200,
//         appAlive: true
//     })


// extending interfaces : existing interface ki values ku leke new interface m daalskte.

// interface User {
//     name : string
// }

// interface Admin extends User {
//     email: string,
//     password: string
// }

// function myUser(obj : Admin){
//     console.log(obj)
// }

// myUser({name: "rxhman", email: "rrr", password:"hsjajak"})


// type alaises: creating own custom types 

type chacha = number

let number : chacha;
number = 20
console.log(number);

// type alaises real usecase : 

type value = number | string | boolean

let myVar : value;
myVar = true
// console.log(myVar);

// | : this is called union type, can be used as or operator of js 

type myCustom = number | string

let xy : myCustom = 5049
// console.log(xy);


// Intersection type (&) ka matlab:
// Do ya zyada types ko combine kar dena, aur final type me sabke properties mandatory ho jaate hain.

// Socho:

// Union (|) = ya ye ya wo

// Intersection (&) = ye bhi + wo bhi

type person = {
    name: string
}

type employyeId = {
    emplyeeid : number
}

type myEmployee = person & employyeId

let akash: myEmployee

akash = {
    name: "akash",
    emplyeeid: 5047
}

// console.log(akash);


type User = {
    id: string,
    email  : number
}

type Admin = {
    role: boolean,
    isAlive: boolean
}

type myFinalType = User & Admin

let myVar1 : myFinalType =  {
    id: "123",
    email: 45,
    role: true,
    isAlive: true
}

// console.log(myVar1);

