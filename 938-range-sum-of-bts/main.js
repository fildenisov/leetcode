function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

/*
 * @param {TreeNode} root
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var rangeSumBST = function (root, low, high) {
  let res = 0;
  if (root === undefined || root === null) {
    return res;
  }

  if (root.val >= low && root.val <= high) {
    res += root.val;
  }

  res += rangeSumBST(root.left, low, high);
  res += rangeSumBST(root.right, low, high);

  return res;
};

const n3 = new TreeNode(3);
const n7 = new TreeNode(7);
const n18 = new TreeNode(18);

const n5 = new TreeNode(5, n3, n7);
const n15 = new TreeNode(15, undefined, n18);

const n10 = new TreeNode(10, n5, n15);

console.log(rangeSumBST(n10, 7, 15));
