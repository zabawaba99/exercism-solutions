var Bob = function() {};

isYelling = function(input) {
  if (!/[A-Z]/.test(input)) {
    // does not contain any capital letters
    return false;
  }

  if (input !== input.toUpperCase()) {
    // all characters are not uppercase
    return false;
  }

  return true;
}

Bob.prototype.hey = function(input) {
  if (isYelling(input)) {
    return 'Whoa, chill out!';
  }

  if (input[input.length-1] === '?') {
    return 'Sure.';
  }

  if (/^\s*$/.test(input)) {
    // there are only white spaces
    return 'Fine. Be that way!';
  }

  return 'Whatever.';
};

module.exports = Bob;
