// function checkUser(Callback){

//     setTimeout(()=>{
//         console.log("checking user");
//         Callback()
//     }, 3000)

// }


// function fetchProfile(Callback){
 
//     setTimeout(()=>{
//         console.log("fetching pfp");
//         Callback()
//     }, 1000)

// }

// function sendEmail(Callback){

//     setTimeout(()=>{
        
//     console.log("sending email");
//     Callback()
//     }, 500)

// }

// function main(){

//     checkUser(()=>{
//         fetchProfile(()=>{
//             sendEmail(()=>{
//                 console.log("done");
                
//             })
//         })
//     })
//     console.log("other requests processing");
    


// }

// main()



function credentials(Callback){
    setTimeout(()=>{
        console.log("validating credentials");
        Callback()
    })
    
}

function generateToken(Callback){
   setTimeout(()=> {
     console.log("generating token");
     Callback()
   })
    
}

function session(Callback){
   setTimeout(()=>{
     console.log("storing sessin");
     let error = new Error("couldn't store session")
     Callback(error)
   })
    
}

function main(){
    console.log("other req processing");
    credentials(()=>{
        generateToken(()=>{
            session((error)=>{
                if (error){
                   console.log(error);
                   return
                }
                console.log("login successful");
                
            })
        })
    })
    
}

main()