import React from 'react'

import UniversityCard from './UniversityCard'
const UniverityPreview  = (props) =>(
        <div className="cards">
            {props.universitiesList.map(universityData => <UniversityCard key={universityData.id} {...universityData} setCurrentPage={props.setCurrentPage} />)}
        </div>

);

UniverityPreview.propType = {
    setCurrentPage: React.PropTypes.func,
}
export default UniverityPreview;

