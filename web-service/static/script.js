const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
console.log(canvas.width);
console.log(canvas.height);
// ctx.imageSmoothingEnabled = false;

var DEFAULT_COLOR_PALETTE = [
  "#FFFFFF", // white
  "#E4E4E4", // light grey
  "#888888", // grey
  "#222222", // black
  "#FFA7D1", // pink
  "#E50000", // red
  "#E59500", // orange
  "#A06A42", // brown
  "#E5D900", // yellow
  "#94E044", // lime
  "#02BE01", // green
  "#00D3DD", // cyan
  "#0083C7", // blue
  "#0000EA", // dark blue
  "#CF6EE4", // magenta
  "#820080", // purple
];

const get = async () => {
  const res = await fetch("http://localhost:8080/board");

  const data = await res.json();
  var uintArray = Base64Binary.decode(data.data);
  // console.log(uintArray);
  console.log(uintArray.byteLength);

  var width = 1000;
  // var i = 0;
  // uintArray.map((d) => {
  //   i++;
  //   // console.log(`Coords X:${x} | Y:${y} | Value:${d}`);
  // });
  for (var i = 0; i < uintArray.length; i++) {
    const x = Math.floor(i % width);
    const y = Math.floor(i / width);
    var d = uintArray[i];

    fillRect(d, x, y);
  }
};
get();

//websocket connection

/*---------------------------------*/

/*
@Fill the rectangle
@d - color hex
@x - x_cords
@y - y_cords
*/
const fillRect = (d, x, y) => {
  ctx.fillStyle = DEFAULT_COLOR_PALETTE[d];
  ctx.fillRect(x, y, 10, 10);
};
url = "ws://localhost:8081/test";
c = new WebSocket(url);

// send = function(data){
//   $("#output").append((new Date())+ " ==> "+data+"\n")
//   c.send({"message":"nice"})
// }

c.onopen = function () {
  // setInterval(
  //   function(){ send("ping") }
  // , 1000 )
  console.log("websocket open");
};

c.onmessage = function (msg) {
  // $("#output").append(new Date() + " <== " + msg.data + "\n");
 const json_msg = JSON.stringify(msg);
 const data = JSON.parse(msg.data);
 fillRect(data.color,data.x,data.y)
};

var Base64Binary = {
  _keyStr: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",

  /* will return a  Uint8Array type */
  decodeArrayBuffer: function (input) {
    var bytes = (input.length / 4) * 3;
    var ab = new ArrayBuffer(bytes);
    this.decode(input, ab);

    return ab;
  },

  removePaddingChars: function (input) {
    var lkey = this._keyStr.indexOf(input.charAt(input.length - 1));
    if (lkey == 64) {
      return input.substring(0, input.length - 1);
    }
    return input;
  },

  decode: function (input, arrayBuffer) {
    //get last chars to see if are valid
    input = this.removePaddingChars(input);
    input = this.removePaddingChars(input);

    var bytes = parseInt((input.length / 4) * 3, 10);

    var uarray;
    var chr1, chr2, chr3;
    var enc1, enc2, enc3, enc4;
    var i = 0;
    var j = 0;

    if (arrayBuffer) uarray = new Uint8Array(arrayBuffer);
    else uarray = new Uint8Array(bytes);

    input = input.replace(/[^A-Za-z0-9\+\/\=]/g, "");

    for (i = 0; i < bytes; i += 3) {
      //get the 3 octects in 4 ascii chars
      enc1 = this._keyStr.indexOf(input.charAt(j++));
      enc2 = this._keyStr.indexOf(input.charAt(j++));
      enc3 = this._keyStr.indexOf(input.charAt(j++));
      enc4 = this._keyStr.indexOf(input.charAt(j++));

      chr1 = (enc1 << 2) | (enc2 >> 4);
      chr2 = ((enc2 & 15) << 4) | (enc3 >> 2);
      chr3 = ((enc3 & 3) << 6) | enc4;

      uarray[i] = chr1;
      if (enc3 != 64) uarray[i + 1] = chr2;
      if (enc4 != 64) uarray[i + 2] = chr3;
    }

    return uarray;
  },
};
