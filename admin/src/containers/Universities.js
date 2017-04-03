/**
 * Created by pdiouf on 2017-03-25.
 */
import React from 'react'
import UniversityPreview from '../components/UniversityPreview'
var app =  {};

// Some initial universities to start with.

app.dirty = false;

// The card manager/holder.
var CardManager = React.createClass({
    getInitialState: function() {
        return {universities: []};
    },

    // When our component mounts we should update the universities with the initial ones.
    componentDidMount: function() {
        var _self = this;
        this.setState({universities: this.props.universitiesList});

        // We'll just cheat a bit and set an interval to watch changes.
        setInterval(function() {
            if (app.dirty) {
                app.dirty = false;
                _self.setState({universities: this.props.universitiesList});
            }
        }, 500);
    },

    // Render our cycle of universities.
    render: function() {
        return (
            <div className="card-cycle">
                <UniversityPreview universitiesList={this.props.universitiesList} />
            </div>
        );
    }
});


// The card application.
var Universities = React.createClass({
    render: function() {
        return (

                <div>
                    <h3>Universities</h3>
                    <CardManager  universitiesList={this.props.universitiesList}/>
                </div>
        );
    }
});

export default Universities;
