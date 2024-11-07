import { CellView } from "./CellView.js";
import { ws } from "./ws.js";

export class Cell {
  hovered = false;
  isFlag = false;
  isActive = false;
  neighbors = 0;
  pressed = false;
  view = null;
  constructor(i, j, data) {
    this.isFlag = data.IsFlag;
    this.neighbors = data.Neighbors;
    this.isActive = data.IsActive;
    this.i = i;
    this.j = j;
  }

  createView(p5, y, x, size, flagImg) {
    this.view = new CellView(p5, y, x, size, this.neighbors, flagImg);
  }

  update(data) {
    this.isFlag = data.IsFlag;
    this.neighbors = data.Neighbors;
    this.isActive = data.IsActive;
  }

  draw() {
    if (this.isFlag) {
      this.view.drawFlag();
      return;
    }
    if (this.isActive) {
      this.view.drawActive();
    } else {
      if (this.pressed) {
        this.view.drawHover();
        return;
      }

      this.view.drawInactive();
    }
  }

  click() {
    if (this.isFlag) {
      return;
    }

    const data = {
      action: "click",
      coords: {
        y: this.i,
        x: this.j,
      },
    };

    ws.send(JSON.stringify(data));
  }

  tap() {
    if (this.isActive && !this.isFlag) {
      this.click();
    } else {
      this.toggleFlag();
    }
  }

  toggleFlag() {
    if (this.isActive) {
      return;
    }
    const data = {
      action: "toggleFlag",
      coords: {
        y: this.i,
        x: this.j,
      },
    };
    ws.send(JSON.stringify(data));
    console.log(data);

    this.isFlag = !this.isFlag;
    this.draw();
  }

  press() {
    this.pressed = true;
    this.draw();
  }

  unpress() {
    this.pressed = false;
    this.draw();
  }
}
