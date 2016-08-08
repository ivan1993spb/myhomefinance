
var React = require('react'),
    Note = require("./Note").Note,
    dates = require("./dates"),
    client = require("./client");

var loadDaysLimit = 70;

var NoteList = React.createClass({
    displayName: "NoteList",

    propTypes: {
        dateFrom: React.PropTypes.object,
        loadDays: React.PropTypes.number
    },

    getDefaultProps: function() {
        var currentDate = dates.addDays(new Date(), 1);
        return {
            dateFrom: new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), 0, 0, 0, 0),
            loadDays: 1
        };
    },

    getInitialState: function() {
        return {
            dateTo:   this.props.dateFrom,
            loadDays: this.props.loadDays < loadDaysLimit ? this.props.loadDays : loadDaysLimit,
            notes:    [],
            loading:  true
        };
    },

    increaseLoadDays: function(loadDays) {
        if (loadDays >= loadDaysLimit) {
            return loadDays;
        }
        return loadDays * 2;
    },

    doLoadMore: function(from, to, callback) {
        console.log("called NoteList.doLoadMore");
        console.log("dates 1", dates.yyyymmdd(from), dates.yyyymmdd(to));
        console.log("dates 2", from, to);

        if (typeof callback === 'function') {
            client.getNotesByDateRange(from, to, function(notes) {
                callback(notes);
            });
        }
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        this.doLoadMore(from, to, function(notes) {
            console.log("load more: setting component state" );
            var newState = {
                loading: false
            };

            if (notes.length) {
                newState.notes = this.state.notes.concat(notes);
                // return initial value
                if (this.state.loadDays != this.props.loadDays) {
                    newState.loadDays = this.props.loadDays;
                }
            } else {
                newState.loadDays = this.increaseLoadDays(this.state.loadDays);
            }

            this.setState(newState);
        }.bind(this));
    },

    componentDidMount: function() {
        console.log("ok");
        this.handleLoadMore();
    },

    handleEdit: function(index, id) {
        console.log("index:", index);
        console.log("id:", id);
        // TODO implement handleEdit
    },

    handleRemove: function(index, id) {
        console.log("index:", index);
        console.log("id:", id);
        // TODO implement handleRemove
    },

    render: function() {
        console.log("render note list", this.state.notes);
        var notes = this.state.notes.map(function(note, index) {
            return (
                <Note
                    id={note.id}
                    key={index}
                    time={note.time}
                    name={note.name}
                    text={note.text}
                    handleEdit={this.handleEdit.bind(this, index)}
                    handleRemove={this.handleRemove.bind(this, index)}
                />
            );
        }.bind(this));

        return (
            <div>
                <h2>Notes list</h2>
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>
                <hr />
                <div>
                    {notes.length > 0 ? notes : (this.state.loading ? "loading" : "empty")}
                </div>
                <hr />
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>

                {this.state.loading ? "loading..." : "loaded: "+this.state.notes.length}

                <button
                    onClick={this.handleLoadMore}
                    disabled={this.state.loading}>load more: {this.state.loadDays} days</button>

            </div>
        );
    }
});

exports.NoteList = NoteList;
