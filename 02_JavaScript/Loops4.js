// Loops: there three typess of loops in js : for loop, while loop, do while loop

// for loop :

// for(i=0; i<=3; i++){
//     console.log("Hello");

// }

// while loop:
// let i = 0;

// while(i<=3){
//     console.log("JS");
//     i++
// }

// do while loop:
// let i = 0

// do {
//     console.log("Loop");
//     i++
// } while(i<=3)

// break : breaks loop at satisfied condition.
// continue: skips only that part in loop and continues

// for(i= 0; i<=10 ; i++){
//     if(i>5){
//         break
//     }
//     console.log(i);

// }

// for(j = 0; j<=10; j++){
//     if(j==9){
//         continue
//     }
//     console.log(j);

// }

// // while
// let i = 0

// while(i<=10){
//     if(i == 6){
//         break
//     }
//     console.log(i);
//     i++
// }

// let j = 0
// while(j<=10){
//     if(j==6){
//           j++
//         continue
//     }
//     console.log(j);
//     j++
// }

// do while
let i = 0;

do {
  if (i == 7) {
    i++
    break;
  }
  console.log(i);
  i++;
} while (i <= 10);

let j = 0;

do {
  if (j == 8) {
    j++
    continue;
  }
  console.log(j);
  j++;
} while (j <= 10);

// Switch statement:

let day = "thursday";

switch (day) {
  case "sunday":
    console.log("sunday");
    break;
  case "monday":
    console.log("monday");
    break;
  case "tuesday":
    console.log("tuesday");
    break;
  default:
    console.log("not a week day");
    break;
}
