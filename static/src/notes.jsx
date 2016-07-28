
var NoteList = React.createClass({
    handleRemove: function(i, note) {
        // var notes = this.props.notes;
        // notes.splice(i, 1);
        // this.replaceProps({notes: notes});
        console.log(i, note);
    },

    render: function() {
        var notes = this.props.notes.map(function(note, i) {
            return (
                <tr>
                    <td>{note.id}</td>
                    <td>{note.name}</td>
                    <td>{note.text}</td>
                    <td><button onClick={this.handleRemove.bind(this, i, note.id)}>delete</button></td>
                </tr>
            );
        }.bind(this));

        return (
            <div>
                <h1>Notes list</h1>
                <table className="pure-table">
                    <thead>
                        <tr>
                            <th>id</th>
                            <th>name</th>
                            <th>text</th>
                            <th>action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {notes}
                    </tbody>
                </table>
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
