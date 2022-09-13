/**
 * @param {number[][]} grid
 * @return {number}
 */
var projectionArea = function (grid) {
  let xCount = 0;
  let yCount = 0;
  let zCount = 0;

  let maxInJ = [];

  for (let i = 0; i < grid.length; i++) {
    let maxInGrid = 0;

    for (let j = 0; j < grid[i].length; j++) {
      if (maxInJ.length == j) {
        maxInJ.push(0);
      }

      if (grid[i][j] != 0) {
        xCount++;
      }

      if (grid[i][j] > maxInGrid) {
        maxInGrid = grid[i][j];
      }

      if (grid[i][j] > maxInJ[j]) {
        maxInJ[j] = grid[i][j];
      }
    }

    yCount += maxInGrid;
  }

  for (let i = 0; i < maxInJ.length; i++) {
    zCount += maxInJ[i];
  }

  return xCount + yCount + zCount;
};

/*
1. count all v != 0
2. sum max from all grids
3. sum max from all grid index


*/

console.log(
  projectionArea([
    [1, 2],
    [3, 4],
  ])
); // 17
console.log(projectionArea([[2]])); // 5
console.log(
  projectionArea([
    [1, 0],
    [0, 2],
  ])
); // 8
console.log(
  projectionArea([
    [1, 4],
    [1, 1],
  ])
); // 14 = 4 + 5 + 5
