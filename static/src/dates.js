
exports.yyyymmdd = function(date) {
    var month = '' + (date.getMonth() + 1),
        day = '' + date.getDate(),
        year = date.getFullYear();

    if (month.length < 2) {
        month = '0' + month;
    }

    if (day.length < 2) {
        day = '0' + day;
    }

    return [year, month, day].join('-');
};

exports.addDays = function(date, days) {
    var newDate = new Date();
    newDate.setTime(date.getTime() + days * 86400000);
    return newDate;
};
