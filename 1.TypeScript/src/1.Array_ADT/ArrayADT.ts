class ArrayADT<T> {
  private A!: T[];
  private _size!: number;
  private _length!: number;

  constructor(size: number | undefined) {
    this.A = new Array(size || 2);
    this._size = size || 2;
    this._length = 0;
  }

  public get size(): number {
    return this._size;
  }
  public get length(): number {
    return this._length;
  }

  public isEmpty() {
    return this.length === 0;
  }
  private makeSizeDouble() {
    this._size = this.size * 2;
    const temp = new Array(this.size);
    for (let i = 0; i < this.length; i++) {
      temp[i] = this.A[i];
    }
    this.A = temp;
  }

  public append(value: T) {
    if (this.length === this.size) {
      this.makeSizeDouble();
    }
    this.A[this.length] = value;
    this._length++;
  }

  public insert(index: number, value: T) {
    if (index > this.length) {
      console.log("index is greater than length");
    } else if (index < 0) {
      console.log("index can't be negative");
    } else {
      if (this.length === this.size) {
        this.makeSizeDouble();
      }

      if (this.length !== index) {
        for (let i = this.length; i > index; i--) {
          this.A[i] = this.A[i - 1];
        }
      }

      this.A[index] = value;
      this._length++;
    }
  }

  public display() {
    if (!this.isEmpty()) {
      for (let i = 0; i < this.length; i++) {
        if (this.length === i + 1) {
          process.stdout.write(`${this.A[i]}`);
        } else {
          process.stdout.write(`${this.A[i]}, `);
        }
      }
    }
    console.log();
  }
}
const arr = new ArrayADT<number>(3);
arr.display();
console.log(arr.length);
console.log(arr.size);

arr.append(1);
arr.append(2);
arr.append(3);
arr.append(4);
arr.display();
console.log(arr.length);
console.log(arr.size);

arr.append(11);
arr.append(12);
arr.display();
console.log(arr.length);
console.log(arr.size);

arr.append(21);
arr.append(22);
arr.display();
console.log(arr.length);
console.log(arr.size);
arr.insert(5, 0);
arr.display();
console.log(arr.length);
console.log(arr.size);
