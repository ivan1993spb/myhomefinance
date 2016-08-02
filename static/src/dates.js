
exports.yyyymmdd = function(date) {
    var mm = date.getMonth() + 1; // getMonth() is zero-based
    var dd = date.getDate();

    return [date.getFullYear(), !mm[1] && '0', mm, !dd[1] && '0', dd].join('');
};

exports.addDays = function(date, days) {
    var newDate = new Date();
    newDate.setDate(date.getDate() + days);
    return newDate;
};
