import { Cell } from "./Cell.js";

export class Board {
  constructor(boardApi) {
    this.height = boardApi.Height;
    this.width = boardApi.Width;
    this.grid = [];
    for (let i = 0; i < this.height; i++) {
      const row = [];
      for (let j = 0; j < this.width; j++) {
        row.push(new Cell(i, j, boardApi.Grid[i][j]));
      }
      this.grid.push(row);
    }
  }

  createView(p5, size, flagImg) {
    this.forEach((cell, i, j) => {
      cell.createView(p5, i * size, j * size, size, flagImg);
    });
    this.draw();
  }

  forEach(callback) {
    this.grid.forEach((row, i) => {
      row.forEach((cell, j) => {
        callback(cell, i, j);
      });
    });
  }

  draw() {
    this.forEach((cell) => {
      cell.draw();
    });
  }

  update(boardApi) {
    this.forEach((cell) => {
      const cellApi = boardApi.Grid[cell.i][cell.j];
      cell.update(cellApi);
    });

    this.draw();
  }
}
