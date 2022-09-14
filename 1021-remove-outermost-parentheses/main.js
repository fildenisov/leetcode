/**
 * @param {string} s
 * @return {string}
 */
var removeOuterParentheses = function (s) {
  let start = 0;
  let counter = 1;
  let result = "";

  for (let i = 1; i < s.length; i++) {
    if (s[i] == "(") {
      counter++;
    } else {
      counter--;
    }

    if (counter === 0) {
      result += s.slice(start + 1, i);
      start = i + 1;
      counter = 1;
      i++;
    }
  }
  return result;
};

var removeOuterParenthesesV2 = function (s) {
  const a = s.split("");
  let start = 0;
  let counter = 1;

  for (let i = 1; i < a.length; i++) {
    if (a[i] == "(") {
      counter++;
    } else {
      counter--;
    }

    if (counter === 0) {
      a.splice(start, 1);
    i--
      a.splice(i, 1);
	i--
      start = i + 1;
      counter = 1;
      i++;
    }
  }
  return a.join("")
};

console.log(removeOuterParenthesesV2("(()())(())(()(()))")); // ()()()()(())
