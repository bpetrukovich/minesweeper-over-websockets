export class CellView {
  color = 0;
  constructor(p5, y, x, size, value, flagImg) {
    this.p5 = p5;
    this.x = x;
    this.y = y;
    this.size = size;
    this.value = value;
    this.flagImg = flagImg;

    const COLORS = [
      p5.color(0),
      p5.color(0, 0, 255),
      p5.color(0, 123, 0),
      p5.color(202, 0, 0),
      p5.color(3, 31, 103),
      p5.color(134, 56, 6),
      p5.color(3, 110, 189),
      p5.color(2, 2, 2),
      p5.color(128, 128, 128),
    ];

    this.color = COLORS[value];
  }

  drawHover() {
    this.p5.fill(150);
    this.p5.strokeWeight(4);
    this.p5.stroke(100);
    this.p5.square(this.x, this.y, this.size);
  }

  drawInactive() {
    this.p5.fill(220);
    this.p5.strokeWeight(4);
    this.p5.stroke(100);
    this.p5.square(this.x, this.y, this.size);
  }

  drawActive() {
    this.p5.fill(150);
    this.p5.strokeWeight(4);
    this.p5.stroke(100);
    this.p5.square(this.x, this.y, this.size);
    if (!this.value) {
      return;
    }
    this.p5.textSize(30);
    this.p5.textStyle(this.p5.BOLD);
    this.p5.textFont("Courier New");
    this.p5.fill(this.color);
    this.p5.textAlign(this.p5.CENTER, this.p5.CENTER);
    this.p5.noStroke();
    this.p5.text(this.value, this.x + this.size / 2, this.y + this.size / 2);
  }

  drawFlag() {
    this.p5.fill(220);
    this.p5.strokeWeight(4);
    this.p5.stroke(100);
    this.p5.square(this.x, this.y, this.size);
    this.p5.imageMode(this.p5.CENTER);
    this.p5.image(
      this.flagImg,
      this.x + this.size / 2,
      this.y + this.size / 2,
      this.size / 1.5,
      this.size / 1.5,
    );
  }
}
