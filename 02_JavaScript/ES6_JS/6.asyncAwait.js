// async function hamesha Promise return karta hai

// await = “jab tak Promise resolve na ho, yahin ruk” (sirf us function ke andar)

// in promise
let todos = fetch("https://jsonplaceholder.typicode.com/todos")
// .then((resp)=> resp.json()).then((data)=>console.log(data))


// async function always returns a Promise
async function getData() {

    // fetch() returns a Promise
    // await waits for it to resolve
    // data is NOT a Promise, it is the Response object
    const data = await fetch("https://jsonplaceholder.typicode.com/todos");

    // data.json() RETURNS a Promise
    // await is needed again to get actual JSON data
    const jsonData = await data.json();

    console.log(jsonData);
}

getData();
