/**
 * Created by pdiouf on 2017-04-04.
 */
import React from 'react'
import UniversityForm from '../containers/UniversityForm'
const NewUniversity = (props) => (
    <div className="NewUniversity">
        <h4 className="page-header">New University </h4>
        <UniversityForm  universitiesList={props.universitiesList} {...props}  setCurrentPage={props.setCurrentPage}/>
    </div>
);

NewUniversity.propTypes= {
    setCurrentPage : React.PropTypes.func,

}
export default NewUniversity;
