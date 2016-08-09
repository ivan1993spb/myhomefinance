
var React = require('react');

var Overlay = React.createClass({
    displayName: "Overlay",

    getInitialState: function() {
        return {
            open: true,
        };
    },

    render: function() {
        return (
            <div>
                {/* background */}
                <div className="overlay-background"></div>
                {/* content */}
                <div className="overlay-block">
                    <h1>{this.props.topic}<button>X</button></h1>
                    <div>
                        {this.props.children}
                    </div>
                </div>
            </div>
        );
    }
});

exports.Overlay = Overlay;
