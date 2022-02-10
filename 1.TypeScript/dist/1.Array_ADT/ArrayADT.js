"use strict";
class ArrayADT {
    A;
    _size;
    _length;
    constructor(size) {
        this.A = new Array(size);
        this._size = size;
        this._length = 0;
    }
    get size() {
        return this._size;
    }
    get length() {
        return this._length;
    }
    isEmpty() {
        return this.length === 0;
    }
    add(value) {
        if (this.length === this.size) {
            this._size = this.size * 2;
            const temp = new Array(this.size);
            for (let i = 0; i < this.length; i++) {
                temp[i] = this.A[i];
            }
            this.A = temp;
            this.A[this.length] = value;
        }
        else {
            this.A[this.length] = value;
        }
        this._length++;
    }
    display() {
        if (!this.isEmpty()) {
            for (let i = 0; i < this.length; i++) {
                if (this.length === i + 1) {
                    process.stdout.write(`${this.A[i]}`);
                }
                else {
                    process.stdout.write(`${this.A[i]}, `);
                }
            }
        }
        console.log();
    }
}
const arr = new ArrayADT(2);
arr.display();
console.log(arr.length);
console.log(arr.size);
arr.add(1);
arr.add(2);
arr.add(3);
arr.add(4);
arr.display();
console.log(arr.length);
console.log(arr.size);
arr.add(11);
arr.add(12);
arr.display();
console.log(arr.length);
console.log(arr.size);
