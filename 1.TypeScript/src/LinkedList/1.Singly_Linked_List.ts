class SNode {
  private _data!: number;
  private _next!: SNode | null;
  public set data(value: number) {
    this._data = value;
  }
  public get data(): number {
    return this._data;
  }
  public set next(value: SNode | null) {
    this._next = value;
  }
  public get next(): SNode | null {
    return this._next;
  }
  constructor(data: number, next: SNode | null) {
    this.data = data;
    this.next = next;
  }
}

class SLinkedList {
  private head!: SNode | null;
  constructor() {
    this.head = null;
  }
  public isEmpty() {
    return this.head === null;
  }
}

const list = new SLinkedList();
console.log(list.isEmpty());
