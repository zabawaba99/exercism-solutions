var Hamming = function() {};

Hamming.prototype.compute = function(original, mutation) {
  if (original.length !== mutation.length) {
    throw 'DNA strands must be of equal length.';
  }

  var count = 0;
  for (var i = 0; i < original.length; i++) {
    if (original[i] !== mutation[i]) {
      count++;
    }
  }
  return count;
}

module.exports = Hamming;