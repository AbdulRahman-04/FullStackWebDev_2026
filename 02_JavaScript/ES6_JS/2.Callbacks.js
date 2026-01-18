// function checkInventory(Callback){
//     setTimeout(()=>{
//         console.log("creating inventory");
//           Callback()
//     }, 1000)
    
// }

// function createOrder(Callback){
//     setTimeout(()=>{
//         console.log("order creating");
        
//         Callback()
//     }, 1000)
    
// }

// function chargePayment(Callback){
//     setTimeout(()=>{
//         console.log("charging payment");
//         Callback()
//     }, 2000)
    
// }

// function sendInVoice(Callback){
//    setTimeout(()=>{
//      console.log("invoice sent!");
//      Callback()
//    }, 100)
// }



// function main(){

//     // Callback hell 
// checkInventory(()=>{
//   createOrder(()=>{
//      chargePayment(()=>{
//         sendInVoice(()=>{
//             console.log("all done!");
            
//         })
//      })
//   })
// })



// console.log("other requests processing");

// }

// main()



// handling err in call backs : 
function checkInventory(Callback){
    setTimeout(()=>{
        console.log("creating inventory");
          Callback()
    }, 1000)
    
}

function createOrder(Callback){
    setTimeout(()=>{
        console.log("order creating");
        const error = new Error("error creating an order")
        Callback(error)
    }, 1000)
    
}

function chargePayment(Callback){
    setTimeout(()=>{
        console.log("charging payment");
        Callback()
    }, 2000)
    
}

function sendInVoice(Callback){
   setTimeout(()=>{
     console.log("invoice sent!");
     Callback()
   }, 100)
}



function main(){

    // Callback hell 
checkInventory(()=>{
  createOrder((error)=>{
    if(error){
        console.log(error);
        
    }
     chargePayment(()=>{
        sendInVoice(()=>{
            console.log("all done!");
            
        })
     })
  })
})



console.log("other requests processing");

}

main()