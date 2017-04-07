/**
 * Created by pdiouf on 2017-04-03.
 */

import  React  from  'react'


const UniversityCard = (props) =>(
    // Create the card.
    <a href="#" onClick={(event) => {props.setCurrentPage(event, {page:'viewUniversity',id:props.id});}}>

    <div id={`univ${+props.id}`} className="card-item">
                <div className={props.color}>
                    <div className="card-info">
                        <div className="card-name">{props.name}</div>
                        <div className="card-desc">{props.website}</div>
                        <div className="card-desc">{props.address.state}</div>
                    </div>
                </div>
                <div className="clear"></div>
                <div className="languages">{props.selectedLanguages.length} languages</div>
            </div>
    </a>
);
export default  UniversityCard;
