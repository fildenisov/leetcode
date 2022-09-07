/**
 * @param {number} x
 * @returns {bool}
 */
var isSelfDividing = function (x) {
  var s = x.toString();
  var n = 0;
  for (var i = 0; i < s.length; i++) {
    n = s[i];
    if (x == 0 || x % n != 0) {
      return false;
    }
  }
  return true;
};

/**
 *
 * @param {number} left
 * @param {number} right
 * @return {number[]}
 */
var selfDividingNumbers = function (left, right) {
  var res = [];
  for (i = left; i <= right; i++) {
    if (isSelfDividing(i)) {
      res.push(i);
    }
  }
  return res;
};

console.log(selfDividingNumbers(1, 22));
