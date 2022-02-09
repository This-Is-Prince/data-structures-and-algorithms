"use strict";
class SNode {
    _data;
    _next;
    set data(value) {
        this._data = value;
    }
    get data() {
        return this._data;
    }
    set next(value) {
        this._next = value;
    }
    get next() {
        return this._next;
    }
    constructor(data, next) {
        this.data = data;
        this.next = next;
    }
}
class SLinkedList {
    head;
    constructor() {
        this.head = null;
    }
    isEmpty() {
        return this.head === null;
    }
    create(arr) {
        arr.forEach((value) => {
            this.add(value);
        });
    }
    add(value) {
        if (this.head === null) {
            this.head = new SNode(value, null);
        }
        else {
            let temp = this.head;
            while (temp.next !== null) {
                temp = temp.next;
            }
            temp.next = new SNode(value, null);
        }
    }
    displayFromStart_ThroughLoop() {
        let temp = this.head;
        while (temp !== null) {
            if (temp.next === null) {
                process.stdout.write(`${temp.data}`);
            }
            else {
                process.stdout.write(`${temp.data}, `);
            }
            temp = temp.next;
        }
        console.log();
    }
    displayFromStart_ThroughRecursion() {
        this.displayFromStart_ThroughRecursion_RecursiveFunction(this.head);
        console.log();
    }
    displayFromStart_ThroughRecursion_RecursiveFunction(curr) {
        if (curr !== null) {
            if (curr.next === null) {
                process.stdout.write(`${curr.data}`);
            }
            else {
                process.stdout.write(`${curr.data}, `);
            }
            this.displayFromStart_ThroughRecursion_RecursiveFunction(curr.next);
        }
    }
    displayFromLast_ThroughRecursion() {
        this.displayFromLast_ThroughRecursion_RecursiveFunction(this.head);
        console.log();
    }
    displayFromLast_ThroughLoop() {
        console.log();
    }
    displayFromLast_ThroughRecursion_RecursiveFunction(curr) {
        if (curr !== null) {
            this.displayFromLast_ThroughRecursion_RecursiveFunction(curr.next);
            if (curr === this.head) {
                process.stdout.write(`${curr.data}`);
            }
            else {
                process.stdout.write(`${curr.data}, `);
            }
        }
    }
}
const list = new SLinkedList();
list.add(1);
list.add(2);
list.add(3);
list.add(4);
list.displayFromStart_ThroughLoop();
console.log(list.isEmpty());
list.displayFromStart_ThroughRecursion();
console.log(list.isEmpty());
list.create([5, 6, 7, 8, 9]);
list.displayFromStart_ThroughLoop();
list.displayFromLast_ThroughRecursion();
