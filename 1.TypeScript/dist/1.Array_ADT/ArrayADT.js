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
        const temp = this.copyFromThis();
        this.A = temp;
    }
    makeSizeHalf() {
        this._size = Math.floor(this.size / 2);
        const temp = this.copyFromThis();
        this.A = temp;
    }
    copyFromThis() {
        const temp = new Array(this.size);
        for (let i = 0; i < this.length; i++) {
            temp[i] = this.A[i];
        }
        return temp;
    }
    append(value) {
        if (this.length === this.size) {
            this.makeSizeDouble();
        }
        this.A[this.length] = value;
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
            if (this.length !== index) {
                for (let i = this.length; i > index; i--) {
                    this.A[i] = this.A[i - 1];
                }
            }
            this.A[index] = value;
            this._length++;
        }
    }
    delete(index) {
        if (index >= this.length) {
            console.log("index is greater than or equal to length");
        }
        else if (index < 0) {
            console.log("index can't be negative");
        }
        else {
            for (let i = index; i < this.length - 1; i++) {
                this.A[i] = this.A[i + 1];
            }
            this._length--;
            if (this.size / 2 >= this.length) {
                this.makeSizeHalf();
            }
        }
    }
    linearSearch(key) {
        for (let i = 0; i < this.length; i++) {
            if (this.A[i] === key) {
                return i;
            }
        }
        return -1;
    }
    improvedLinearSearch(key) {
        for (let i = 0; i < this.length; i++) {
            if (this.A[i] === key) {
                const temp = this.A[i];
                this.A[i] = this.A[0];
                this.A[0] = temp;
                return 0;
            }
        }
        return -1;
    }
    binarySearch_ThroughLoop(key) {
        let low = 0;
        let high = this.length - 1;
        let mid = 0;
        while (low <= high) {
            mid = Math.floor((low + high) / 2);
            if (key === this.A[mid]) {
                return mid;
            }
            else if (key < this.A[mid]) {
                high = mid - 1;
            }
            else {
                low = mid + 1;
            }
        }
        return -1;
    }
    binarySearch_ThroughRecursion(key) {
        return this.binarySearch_ThroughRecursion_Util(key, 0, this.length - 1);
    }
    binarySearch_ThroughRecursion_Util(key, low, high) {
        if (low <= high) {
            let mid = Math.floor((low + high) / 2);
            if (key === this.A[mid]) {
                return mid;
            }
            else if (key < this.A[mid]) {
                return this.binarySearch_ThroughRecursion_Util(key, low, mid - 1);
            }
            else {
                return this.binarySearch_ThroughRecursion_Util(key, mid + 1, high);
            }
        }
        else {
            return -1;
        }
    }
    get(index) {
        if (index >= 0 && index < this.length) {
            return this.A[index];
        }
        else {
            return null;
        }
    }
    set(index, value) {
        if (index >= 0 && index < this.length) {
            this.A[index] = value;
            return true;
        }
        else {
            return false;
        }
    }
    max() {
        let max = Number.MIN_VALUE;
        for (let i = 0; i < this.length; i++) {
            if (this.A[i] > max) {
                max = this.A[i];
            }
        }
        return max;
    }
    min() {
        let min = Number.MAX_VALUE;
        for (let i = 0; i < this.length; i++) {
            if (this.A[i] < min) {
                min = this.A[i];
            }
        }
        return min;
    }
    sum() {
        let sum = 0;
        for (let i = 0; i < this.length; i++) {
            sum += this.A[i];
        }
        return sum;
    }
    avg() {
        return this.sum() / this.length;
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
arr.delete(0);
arr.delete(4);
// arr.delete(6);
arr.display();
console.log(arr.length);
console.log(arr.size);
console.log(arr.linearSearch(22));
console.log(arr.binarySearch_ThroughLoop(21));
console.log(arr.binarySearch_ThroughRecursion(3));
