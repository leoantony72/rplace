const canvas = document.getElementById("myCanvas");
const ctx = canvas.getContext("2d");

const CELL_SIZE = 10; // Each cell is 10x10 pixels
const GRID_SIZE = 100; // Total grid size of 1000x1000 cells
const CANVAS_SIZE = CELL_SIZE * GRID_SIZE; // Total canvas size in pixels
canvas.width = CANVAS_SIZE;
canvas.height = CANVAS_SIZE;
console.log(canvas.width, canvas.height);
let activeColorIndex = 5;
const DEFAULT_COLOR_PALETTE = [
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

// Initialize grid data (initially all white)
let grid = new Array(GRID_SIZE).fill().map(() => new Array(GRID_SIZE).fill(0));

// Draw initial grid
function drawGrid() {
  for (let x = 0; x < GRID_SIZE; x++) {
    for (let y = 0; y < GRID_SIZE; y++) {
      ctx.fillStyle = DEFAULT_COLOR_PALETTE[grid[x][y]];
      ctx.fillRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }
  }
}
// Function to fetch initial board data from server
const getInitialBoard = async () => {
  const res = await fetch("http://localhost:8080/board");
  const data = await res.json();
  const uintArray = Base64Binary.decode(data.data);

  console.log(uintArray.byteLength);

  const width = canvas.width;

  for (let i = 0; i < uintArray.length; i++) {
    const x = Math.floor(i % width);
    const y = Math.floor(i / width);
    var d = uintArray[i];
    // console.log(x, y);

    const gridX = Math.floor(x / CELL_SIZE);
    const gridY = Math.floor(y / CELL_SIZE);
    ctx.fillStyle = DEFAULT_COLOR_PALETTE[d];
    ctx.fillRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE, CELL_SIZE);
  }
};

// Function to handle mouse click and place pixel
function handleClick(event) {
  const rect = canvas.getBoundingClientRect();
  const mouseX = event.clientX - rect.left;
  const mouseY = event.clientY - rect.top;

  const gridX = Math.floor(mouseX / CELL_SIZE);
  const gridY = Math.floor(mouseY / CELL_SIZE);

  // Update local grid and send to server
  // Update local grid and send to server
  grid[gridX][gridY] = activeColorIndex;

  // Draw updated pixel on canvas
  ctx.fillStyle = DEFAULT_COLOR_PALETTE[activeColorIndex];
  ctx.fillRect(gridX * CELL_SIZE, gridY * CELL_SIZE, CELL_SIZE, CELL_SIZE);

  // Send pixel data to server (simulate)
  sendPixelToServer(gridX, gridY, activeColorIndex);
}

// Simulate sending pixel data to server (replace with actual server communication)
function sendPixelToServer(x, y, colorIndex) {
  console.log(`Sending pixel (${x},${y}) with color ${colorIndex} to server`);
  fetch(`http://localhost:8080/tile?x=${x}&y=${y}&color=${colorIndex}`, {
    method: "POST",
  });
  // Here you would typically send the pixel data to your server via WebSocket or HTTP POST
}

// Event listener for mouse click to place pixel
canvas.addEventListener("click", handleClick);

const colorBoxes = document.querySelectorAll(".colorBox");

colorBoxes.forEach((colorBox) => {
  colorBox.addEventListener("click", function () {
    const colorIndex = parseInt(colorBox.getAttribute("data-color"));
    activeColorIndex = colorIndex;
    console.log(`Active color index updated to ${activeColorIndex}`);
  });
});

// Initial draw of the grid and fetch initial board data
drawGrid();
getInitialBoard();

// WebSocket setup (replace with your actual WebSocket logic)
const socket = new WebSocket("ws://localhost:8081/test");

socket.onopen = function () {
  console.log("WebSocket connected.");
};

socket.onmessage = function (msg) {
  const data = JSON.parse(msg.data);
  const { x, y, color } = data;

  // Update local grid with WebSocket data
  grid[x][y] = color;

  // Draw updated pixel on canvas
  ctx.fillStyle = DEFAULT_COLOR_PALETTE[color];
  ctx.fillRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE, CELL_SIZE);
};

var Base64Binary = {
  _keyStr: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",

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
