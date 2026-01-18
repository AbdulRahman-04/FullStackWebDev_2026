/*

 objects: objects are non primitive data types in js which can store more than one value in itself

*/

let obj1 = {
    name: "rxhman",
    rollNo: 49
}

console.log(obj1);

// crud on objects : 

// above we created obj, now we will update it 
obj1.name = "Syed Abdul Rahman"
obj1.rollNo = 5049
console.log(obj1);

delete obj1.rollNo
console.log(obj1);


let student = {
    name: "rxhman",
    age: 21,
    college : " dcet",
    profession: "full stack web dev"
}

console.log(student);
console.log(student.name, student.profession);

student.college = "DCET"

delete student.college
console.log(student);


// object methods : 

// obj keys : 
console.log(Object.keys(student));

// obj values :
console.log(Object.values(student));

// obj entries: 
console.log(Object.entries(student));

// obj.fromenteries : converts arry into obj
let arr = [["a", 1], ["b", 2]]
console.log(Object.fromEntries(arr));

// obj.is
let a = 10
let b = 20
console.log(Object.is(a,b));

// obj.seal: 
 let user = {
    name: "bhaiyya"
 }

 Object.seal(user)
 user.age = 21
 console.log(user);
 

 // obj.toString(): 
 let obj = { name: "Rahman", age: 21 };

console.log(obj.toString());


// obj.assign: assigns keys and values of one obj to other obj

let obj3 = {
    name: "fxhad"
}

let obj2 = {
    age: "22"
}

console.log(Object.assign(obj3, obj2));


// obj.freeze: 
Object.freeze(obj3)
obj3.add = "plus"
console.log(obj3);
