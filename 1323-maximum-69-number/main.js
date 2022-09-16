/**
 * @param {number} num
 * @return {number}
 */
var maximum69Number = function (num) {
  let s = "" + num;
  for (let i = 0; i < s.length; i++) {
    console.log(s[i]);
    if (s[i] == "6") {
      s = s.slice(0, i) + "9" + s.slice(i + 1);
      break;
    }
  }
  return parseInt(s);
};

console.log(maximum69Number(9669));
console.log(maximum69Number(666));
