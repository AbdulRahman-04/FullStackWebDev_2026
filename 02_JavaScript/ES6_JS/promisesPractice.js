function validateOrder(){
 const promise = new Promise ((resolve, reject)=>{
    setTimeout(()=>{
      const success = false 

      if(success){
          console.log("validating order");
           resolve()
      } else {
        reject("validation failed")
      }
       
    },10000)
 })

  return promise
}

function createOrder(){

    const promise = new Promise((resolve, reject)=>{
        setTimeout(()=>{
            const success = true
            if (success){
                console.log("creating order");
                resolve()
            } else {
                reject("order creation rejected")
            }
        }, 3000)
    })

    return promise

}
function processPayment(){

    const promise = new Promise((resolve, reject)=>{
        setTimeout(()=>{
            const success = true 
            if(success){
                console.log("processing payment");
            resolve()
            } else {
                reject("payment failes")
            }
        }, 7000)

    })
    return promise

}
function sendConfirmation(){

    const promise = new Promise((resolve, reject)=>{
        setTimeout(()=>{
            const success = true
            if (success){
                console.log("order confirmend");
            resolve()
            } else {
                reject("email send failed")
            }
        }, 4000)
    })

    return promise

}

function main(){

    validateOrder().then(()=> createOrder()).then(()=> processPayment()).then(()=> sendConfirmation()).then(()=> console.log("order completed")).catch((error)=>{
        console.log(error)
    })

}

main()