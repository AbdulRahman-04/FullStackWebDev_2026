// function one() {
//   const promise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//       console.log("one");
//         resolve()
//     }, 2000);

//   });

//   return promise;
// }

// function two() {
//   const promise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//       console.log("two");
//         resolve()
//     }, 2000);

//   });

//   return promise
// }

// function main() {
//   one().then(()=> two())
//   console.log("processing other requests");
// }

// main();

// ex : task file upload flow

function fileSize() {
  const promise = new Promise((resolve, reject) => {
    setTimeout(() => {
      const success = true;

      if (success) {
        console.log("file size oke");
        resolve();
      } else {
        reject("file too big");
      }
    }, 1000);
  });

  return promise;
}

function uploaded() {
  const promise = new Promise((resolve, reject) => {
    setTimeout(() => {
      const success = true;
      if (success) {
        console.log("file uploaded");
        resolve();
      } else {
        reject("file upload failes");
      }
    }, 1000);
  });
  return promise;
}

function metaDataSave() {
  const promise = new Promise((resolve, reject) => {
    setTimeout(() => {
      const success = true;

      if (success) {
        console.log("metadata save");
        resolve();
      } else {
        reject("metadata couldn't be saved");
      }
    }, 3000);
  });

  return promise;
}

function main() {
  fileSize()
    .then(() => uploaded())
    .then(() => metaDataSave())
    .then(() => console.log("upload completed"))
    .catch((err) => {
      console.log(err);
    });
}

main();
