function DnaTransscriber() {};

var mapping = {
  'G': 'C',
  'C': 'G',
  'T': 'A',
  'A': 'U'
}

DnaTransscriber.prototype.toRna = function(dna) {
  rna = dna.split('').map(function(letter){
    return mapping[letter];
  })
  return rna.join('');
}

module.exports = DnaTransscriber;