
var NoteRow = React.createClass({
    render: function() {
        return (
            <div>
                <b>{this.props.note.id} - {this.props.note.name}</b>
                <p>{this.props.note.text}</p>
            </div>
        );
    }
});

var NoteList = React.createClass({
    render: function() {
        var notes = [];
        this.props.notes.forEach(function(note) {
            notes.push(<NoteRow note={note}/>);
        });

        return (
            <div>
                <h1>Notes list</h1>
                <div>{notes}</div>
            </div>
        );
    }
});

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
