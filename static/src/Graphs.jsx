
var React = require('react'),
    ChartistGraph = require('react-chartist');

var Graphs = React.createClass({
    displayName: "Graphs",

    render: function() {

        var data = {
            labels: ['W1', 'W2', 'W3', 'W4', 'W5', 'W6', 'W7', 'W8', 'W9', 'W10'],
            series: [
                [10, 2,  4,  8,  15, 3,  1,  4,  6,  2,  21],
                [2,  3,  10, 11, 2,  10, 11, 13, 10, 22, 2],
                [13, 24, 15, 5,  3,  15, 12, 15, 4,  17, 5]
            ]
        };

        var options = {
            axisX: {
                labelInterpolationFnc: function(value, index) {
                    return index % 2 === 0 ? value : null;
                }
            },
            showArea: true
        };

        var type = 'Line';

        var style = {
            height: '500px'
        };

        return (
            <div>
                <ChartistGraph data={data} options={options} type={type} style={style}/>
            </div>
        );
    }
});

exports.Graphs = Graphs;
