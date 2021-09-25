class Stack<T> {
  private _store: T[] = [];

  push(val: T) {
    this._store.push(val);
  }

  pop(): T | undefined {
    return this._store.pop();
  }
}

const stack = new Stack();

stack.push(1);
stack.push(2);
stack.push(3);

for (let i = 1; i <= 4; i++) {
  let v: any = stack.pop();
  if (v != undefined) {
    console.log(v);
  } else {
    console.log("undefined");
  }
}
