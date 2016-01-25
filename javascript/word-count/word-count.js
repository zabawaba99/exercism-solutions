function Words() {};

Words.prototype.count = function(sentence) {
  var wordCount = {};
  sentence.split(/[\s]+/).map(function(word){
    if (word.length === 0) {
      return;
    }

    if (typeof wordCount[word] !== 'number') {
      wordCount[word] = 0;
    }

    wordCount[word] = wordCount[word] + 1;
  });

  return wordCount;
}

module.exports = Words;