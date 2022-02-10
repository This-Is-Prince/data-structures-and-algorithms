"use strict";
class ArrayADT {
    A;
    _size;
    _length;
    constructor(size) {
        this.A = new Array(size || 2);
        this._size = size || 2;
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
    makeSizeDouble() {
        this._size = this.size * 2;
        const temp = new Array(this.size);
        for (let i = 0; i < this.length; i++) {
            temp[i] = this.A[i];
        }
        this.A = temp;
    }
    add(value) {
        if (this.length === this.size) {
            this.makeSizeDouble();
            this.A[this.length] = value;
        }
        else {
            this.A[this.length] = value;
        }
        this._length++;
    }
    insert(index, value) {
        if (index > this.length) {
            console.log("index is greater than length");
        }
        else if (index < 0) {
            console.log("index can't be negative");
        }
        else {
            if (this.length === this.size) {
                this.makeSizeDouble();
            }
            if (this.length === index) {
                this.A[this.length] = value;
            }
            else {
                let tmpIndex = this.length;
                while (index !== tmpIndex) {
                    this.A[tmpIndex] = this.A[tmpIndex - 1];
                    tmpIndex--;
                }
                this.A[index] = value;
            }
            this._length++;
        }
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
const arr = new ArrayADT(3);
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
arr.add(21);
arr.add(22);
arr.display();
console.log(arr.length);
console.log(arr.size);
arr.insert(8, 0);
arr.display();
console.log(arr.length);
console.log(arr.size);
