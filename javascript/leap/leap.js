var Leap = function(year) {
	this.year = year
};

Leap.prototype.isLeap = function() {
	if (this.year % 4 != 0) {
		return false;
	}

	if (this.year % 100 == 0 && this.year % 400 != 0) {
		return false;
	}

	return true
}

module.exports = Leap