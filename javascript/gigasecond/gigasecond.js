var gigsecs = Math.pow(10,9);

function Gigasecond(date) {
  this.gs = new Date(date.toString());
  this.gs.setUTCSeconds(this.gs.getUTCSeconds() + gigsecs);
}

Gigasecond.prototype.date = function() {
  return this.gs;
}

module.exports = Gigasecond;