
var React = require('react');

var Overlay = React.createClass({
    displayName: "Overlay",

    propTypes: {
        topic: React.PropTypes.string.isRequired,
        close: React.PropTypes.func.isRequired
    },

    render: function() {
        return (
            <div>
                {/* background */}
                <div className="overlay-background" onClick={this.props.close}></div>
                {/* content */}
                <div className="overlay-block">
                    <h1>{this.props.topic}<button onClick={this.props.close}>X</button></h1>
                    <div>
                        {this.props.children}
                    </div>
                </div>
            </div>
        );
    }
});

exports.Overlay = Overlay;
