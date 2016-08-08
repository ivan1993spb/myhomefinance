
var React = require('react');
var Link = require('react-router').Link;
var IndexLink = require('react-router').IndexLink;

var Main = React.createClass({
    displayName: "Main",

    render: function() {
        return (
            <div>
                <ul role="nav">
                    <li><IndexLink to="/" activeClassName="menu-btn-active">index</IndexLink></li>
                    <li><Link to="/notes" activeClassName="menu-btn-active">notes</Link></li>
                    <li><Link to="/history" activeClassName="menu-btn-active">history</Link></li>
                </ul>
                <p>_______</p>
                {this.props.children}
                <p>_______</p>
            </div>
        );
    }
});

exports.Main = Main;
