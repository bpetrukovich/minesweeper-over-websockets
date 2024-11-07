import { Board } from "./Board.js";
import { ws } from "./ws.js";

document.addEventListener("contextmenu", (e) => e.preventDefault());

ws.onmessage = (message) => {
  const data = JSON.parse(message.data);
  console.log(data);
  if (data.action == "gameOver") {
    alert("game over");
  } else if (data.action == "in") {
    alert("win");
  }
  if (!board) {
    board = new Board(data.board);
    new p5(visualize);
  } else {
    board.update(data.board);
  }
};

document.getElementById("new-game").onclick = () => {
  if (!ws) {
    return;
  }
  console.log("send");
  ws.send(
    JSON.stringify({
      action: "newGame",
      coords: {
        y: 0,
        x: 0,
      },
    }),
  );
};

let board = null;

function visualize(p5) {
  let flagImg = null;
  p5.preload = function () {
    flagImg = p5.loadImage("flag.png");
  };

  let cellSize = null;
  p5.setup = function () {
    const windowWidth = window.innerWidth;
    const windowHeight = window.innerHeight;
    cellSize =
      Math.min(windowWidth, windowHeight) / Math.max(board.height, board.width);
    p5.createCanvas(cellSize * board.width, cellSize * board.height);
    p5.background(0);
    board.createView(p5, cellSize, flagImg);
  };

  let last = null;
  p5.mouseDragged = function () {
    const cell = getHoveredCell();
    if (!cell) {
      return;
    }
    if (last && last != cell) {
      last.unpress();
    }
    cell.press();
    last = cell;
  };

  p5.mouseReleased = function () {
    const cell = getHoveredCell();
    if (!cell) {
      return;
    }

    cell.unpress();

    if (p5.mouseButton === p5.RIGHT) {
      cell.toggleFlag();
    } else {
      cell.click();
    }
  };

  p5.touchEnded = function () {
    const cell = getHoveredCell();
    if (!cell) {
      return;
    }
    cell.tap();
  };

  p5.mousePressed = function () {
    const cell = getHoveredCell();
    if (!cell) {
      return;
    }
    if (p5.mouseButton === p5.LEFT) {
      cell.press();
    }
  };

  function getHoveredCell() {
    const i = Math.floor(p5.mouseY / cellSize);
    const j = Math.floor(p5.mouseX / cellSize);
    if (i >= board.height || j >= board.width) {
      return null;
    }
    return board.grid[i][j];
  }
}
