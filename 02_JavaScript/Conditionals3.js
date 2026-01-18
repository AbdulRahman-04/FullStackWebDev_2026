// Conditionals in js :

/*
 there are 4 types of conditionals in js : 

  - If statement : single condition
  - If else : two conditions 
  - If else if : more than 2 conditions 
  - nested if else : aalternate of if else if

*/

// if statement : 
let age = 20
if (age >= 18) {
  console.log("you can get license");
  
}

// if else : 
let num = 26

if (num%2 == 0) {
    console.log("even num", num);
    
} else {
    console.log("odd num", num);
    
}

// if else if : 
let marks = 85

if (marks>= 90){
    console.log("A");
    
} else if(marks >= 80 && marks < 90){
    console.log("B");
    
} else if(marks >= 70 && marks < 80){
    console.log("C");
    
} else if (marks >= 60 && marks <70){
    console.log("D");
    
} else {
    console.log("Fail");
    
}

// nested if else 

let isLoggedin = false
let role = "rahman"

if(isLoggedin){
    if(role === "admin"){
        console.log("welcm admin");
        
    } else {
        console.log("wlcm user");
        
    }
} else {
    console.log("pls login first");
    
}