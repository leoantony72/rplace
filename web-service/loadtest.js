const fetch = require("node-fetch");

var j = 0;
for (var i = 0; i < 520; i++) {
  var random1 = generateRandom((min = 0), (max = 999));
  var random2 = generateRandom((min = 0), (max = 999));
  var random3 = generateRandom();
  fetch(`http://localhost:8080/tile?x=${j}&y=${i}&color=${random3}`, {
    method: "POST",
  });
  console.log(j)
  j++
  // fetch(`http://localhost:8080/tile?x=0&y=0&color=${random3}`, {
  //   method: "POST",
  // });
}

function generateRandom(min = 0, max = 15) {
  // find diff
  let difference = max - min;

  // generate random number
  let rand = Math.random();

  // multiply with difference
  rand = Math.floor(rand * difference);

  // add with min value
  rand = rand + min;

  return rand;
}

// console.log(generateRandom());
