
var ReactDOM = require('react-dom'),
    React = require('react'),

    Router = require('react-router').Router,
    Route = require('react-router').Route,
    IndexRoute = require('react-router').IndexRoute,
    browserHistory = require('react-router').browserHistory,

    Main = require('./Main').Main,
    NoteList = require('./NoteList').NoteList,
    History = require('./History').History,
    Graphs = require('./Graphs').Graphs,
    Index = require('./Index').Index;

ReactDOM.render((
    <Router history={browserHistory}>
        <Route path="/" component={Main}>
            <IndexRoute component={Index} />

            <Route path="/history" component={History} />
            <Route path="/notes" component={NoteList} />
            <Route path="/graphs" component={Graphs} />
        </Route>
    </Router>
), document.getElementById('main'));
