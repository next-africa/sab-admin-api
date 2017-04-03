import React from 'react'

import UniversityCard from './UniversityCard'
var UniverityPreview = React.createClass({
    render: function() {
        return (
                <div className="cards">
                    {this.props.universitiesList.map(universityData => <UniversityCard key={universityData.id} {...universityData} />)}
                </div>
        );
    }
});
export default UniverityPreview;

