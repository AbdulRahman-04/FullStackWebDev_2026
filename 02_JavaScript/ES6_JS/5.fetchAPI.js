// fetch API: used to fetch data from a given URL
// response comes in raw format, so we convert it to JSON
// fetchPromise is a Promise which says:
// "data will come, but not yet"

// todos ek Promise hai
// jo bol raha hai:
// “fetch method use karke data aayega, abhi nahi aaya”

let fetchPromise = fetch("https://jsonplaceholder.typicode.com/todos")
// .then((responseUrlKa)=> responseUrlKa.json()).then((printKro)=> console.log(printKro)) 


// task1 : 

const usersFetchPromise = fetch("https://jsonplaceholder.typicode.com/users")
// .then((data)=> data.json()).then((Resp)=> console.log(Resp))

// task 2 :
// const completeTodoPromise = fetch("https://jsonplaceholder.typicode.com/todos")
// .then((data)=> data.json()).then((resp)=> console.log(resp.filter((x)=>{
//     return x.completed === true
// })))