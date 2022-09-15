function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/*
 * @param {ListNode} head
 * @return {ListNode}
 */
var sortList = function (head) {
  if (head == undefined) {
    return head;
  }
  const arr = [];

  while (head != null) {
    arr.push(head);
    head = head.next;
  }

  arr.sort(function (a, b) {
    return a.val - b.val;
  });

  for (let i = 0; i < arr.length; i++) {
    if (arr.length == i + 1) {
      arr[i].next = null;
    } else {
      arr[i].next = arr[i + 1];
    }
  }

  return arr[0];
};

let n3 = new ListNode(3, undefined);
let n1 = new ListNode(1, n3);
let n2 = new ListNode(2, n1);
let n4 = new ListNode(4, n2);

console.log(sortList(n4));
