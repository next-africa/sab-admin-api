/**
 * Created by pdiouf on 2017-04-05.
 */
import React from 'react'
import SingleUniversity from '../containers/SingleUniversity'

const ViewUniversity  = (props) => (

    <SingleUniversity setCurrentPage={props.setCurrentPage} universitiesList={props.universitiesList} currentPageId={props.currentPageId}/>
);
ViewUniversity.propTypes = {
    setCurrentPage:  React.PropTypes.func,
};
export default  ViewUniversity;