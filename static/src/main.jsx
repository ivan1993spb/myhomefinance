
var ReactDOM = require('react-dom');
var React = require('react');

// var Router = require('react-router').Router;
// var Route = require('react-router').Route;
// var IndexRoute = require('react-router').IndexRoute;
// var Link = require('react-router').Link;
// var browserHistory = require('react-router').browserHistory;
// var IndexLink = require('react-router').IndexLink;

var NoteList = require('./note').NoteList;
// var dates = require('./dates');
// var ACTIVE = { color: 'red' };

// var Counter = React.createClass({
//     getInitialState: function() {
//         return {count: 1};
//     },
//
//     handleClick: function() {
//         this.setState({count: this.state.count + 1});
//     },
//
//     render: function() {
//         var items = [];
//         for (var i = 0; i < this.state.count; i++) {
//             items.push(<div>{i}</div>);
//             console.log(i);
//         }
//         return (
//             <div>
//                 <h1>Counter</h1>
//                 {items}
//                 <button onClick={this.handleClick}>OK</button>
//             </div>
//         );
//     }
// });

// var App = React.createClass({
//     render: function(test) {
//         return (
//             <div>
//                 <h1>Test</h1>
//                 <ul>
//                     <li><Link      to="/"      activeStyle={ACTIVE}>/</Link></li>
//                     <li><IndexLink to="/"      activeStyle={ACTIVE}>/IndexLink</IndexLink></li>
//
//                     <li><IndexLink      to="/notes" activeStyle={ACTIVE}>/notes</IndexLink></li>
//                     <li><Link      to="/notes" activeStyle={ACTIVE}>/notes</Link></li>
//                 </ul>
//                 <span>{test}</span>
//             </div>
//         );
//     }
// });

// var Index = React.createClass({
//     render: function() {
//         return (<div>Index</div>);
//     }
// });

// ReactDOM.render((
//     <Router history={browserHistory}>
//         <Route path="/" component={App}>
//             <IndexRoute component={Index}/>
//             <Route path="/" component={Index}/>
//             <Route path="/notes" component={NoteList}/>
//         </Route>
//     </Router>
// ), document.getElementById('content'));

ReactDOM.render(
    (<NoteList from={new Date("2016-08-04")} days={20} />), document.getElementById('content')
);
