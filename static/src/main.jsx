
var ReactDOM = require('reactdom');
var NoteList = require('./notes').noteList;
var dates = require('./dates');

var NOTES = [
    {id: 1, name: "first", text: "text first text"},
    {id: 2, name: "second", text: "text second text"},
    {id: 3, name: "third", text: "text third text"},
    {id: 4, name: "fourth", text: "text fourth text"}
];

ReactDOM.render(
    <NoteList notes={NOTES} />,
    document.getElementById('content')
);
