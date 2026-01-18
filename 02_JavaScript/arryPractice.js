/* =====================================================
   ARRAY PRACTICE – CONFIDENCE SET (NO REDUCE)
   Instructions:
   - Har question ke neeche TODO complete karo
   - Console.log se output check karo
===================================================== */


/* 1️⃣ MAP
   Double each number in the array
   Expected: [6, 12, 18, 24]
*/
const mapNums = [3, 6, 9, 12];
// TODO
// const resultMap = ...
// console.log("map:", resultMap);


/* 2️⃣ FOREACH
   Print: "Value is: X" for each element
*/
const forEachItems = ["apple", "banana", "mango"];
// TODO
// forEachItems.forEach(...)
// console.log("forEach done");


/* 3️⃣ FILTER
   Get only numbers greater than 50
   Expected: [60, 72, 90]
*/
const filterScores = [45, 60, 72, 30, 90];
// TODO
// const resultFilter = ...
// console.log("filter:", resultFilter);


/* 4️⃣ FIND
   Find the first user whose role is "admin"
*/
const findUsers = [
  { id: 1, role: "user" },
  { id: 2, role: "admin" },
  { id: 3, role: "admin" }
];
// TODO
// const resultFind = ...
// console.log("find:", resultFind);


/* 5️⃣ SOME
   Check if at least one even number exists
   Expected: true
*/
const someNums = [3, 7, 9, 11, 14];
// TODO
// const resultSome = ...
// console.log("some:", resultSome);


/* 6️⃣ EVERY
   Check if all numbers are positive
   Expected: false
*/
const everyNums = [2, 4, 6, -8];
// TODO
// const resultEvery = ...
// console.log("every:", resultEvery);


/* 7️⃣ INCLUDES
   Check if "react" exists in the array
   Expected: true
*/
const includeSkills = ["html", "css", "js", "react"];
// TODO
// const resultIncludes = ...
// console.log("includes:", resultIncludes);


/* 8️⃣ PUSH
   Add "nodejs" and "mongodb" at the end
*/
const pushTech = ["html", "css", "js"];
// TODO
// pushTech.push(...)
// console.log("push:", pushTech);


/* 9️⃣ POP
   Remove last element and print it
*/
const popNums = [10, 20, 30, 40];
// TODO
// const removed = ...
// console.log("pop removed:", removed);
// console.log("after pop:", popNums);


/* 🔟 SLICE
   Get middle 3 elements
   Expected: [3, 4, 5]
*/
const sliceArr = [1, 2, 3, 4, 5, 6, 7];
// TODO
// const resultSlice = ...
// console.log("slice:", resultSlice);


/* 1️⃣1️⃣ SORT
   Sort numbers in ascending order
   Expected: [1, 5, 25, 40, 100]
*/
const sortNums = [40, 100, 1, 5, 25];
// TODO
// sortNums.sort(...)
// console.log("sort:", sortNums);


/* 1️⃣2️⃣ SPLICE
   Insert "red" and "blue" in the middle
   Expected: ["black", "red", "blue", "white"]
*/
const spliceColors = ["black", "white"];
// TODO
// spliceColors.splice(...)
// console.log("splice:", spliceColors);


/* =====================================================
   END
===================================================== */


const resultMap = mapNums.map((x)=>{
    return x*2
})

// console.log(resultMap);


// const forEachVal = forEachItems.forEach((x)=>{
//     console.log(`value is ${x}`);
    
// })


// const greatNum = filterScores.filter((x)=>{
//     return x>50
// })
// console.log(greatNum);


const getAdmin = findUsers.find((x)=>{
    return x.role === "admin"
})

// console.log(getAdmin);


const oneEven = someNums.some((x)=>{
    return x%2==0
})

console.log(oneEven);


const post = everyNums.every((x)=>{
    return x>0
})
console.log(post);



const checkReact = includeSkills.includes("react")
console.log(checkReact);


pushTech.push("node.js", "mongo_db")
console.log(pushTech);

pushTech.pop()
console.log(pushTech);


const mySlice = sliceArr.slice(2,5)
console.log(mySlice);


const sorted = sortNums.sort((a,b)=> a-b)
console.log(sorted);


spliceColors.splice(1, 0, "red", "blue")
console.log(spliceColors);






