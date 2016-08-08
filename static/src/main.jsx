
var ReactDOM = require('react-dom');
var React = require('react');

var Router = require('react-router').Router;
var Route = require('react-router').Route;
var IndexRoute = require('react-router').IndexRoute;
var browserHistory = require('react-router').browserHistory;

var Main = require('./Main').Main;
var NoteList = require('./NoteList').NoteList;
var History = require('./History').History;
var Index = require('./Index').Index;

ReactDOM.render((
    <Router history={browserHistory}>
        <Route path="/" component={Main}>
            <IndexRoute component={Index}/>

            <Route path="/history" component={History} />
            <Route path="/notes" component={NoteList} />
        </Route>
    </Router>
), document.getElementById('main'));
