/**
 * Created by pdiouf on 2017-04-03.
 */

import  React  from  'react'
import {Link} from 'react-router-dom'


const UniversityCard = (props) =>(
    // Create the card.

        <Link to={`universitiy/${props.countryCode}${props.id+1}`}>
            <div id={`univ-{props.key}`} className="card-item">
                <div className={props.color}>
                    <div className="card-info">
                        <div className="card-name">{props.universityName}</div>
                        <div className="card-desc">{props.website}</div>
                        <div className="card-desc">{props.address.state}</div>
                    </div>
                </div>
                <div className="clear"></div>
                <div className="languages">{props.selectedLanguages.length} languages</div>
            </div>
        </Link>
);

export default  UniversityCard;
