
var months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

exports.yyyymmdd = function(date) {
    var month = months[date.getMonth()],
        day = '' + date.getDate(),
        year = date.getFullYear();

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
