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
}
const list = new SLinkedList();
console.log(list.isEmpty());
