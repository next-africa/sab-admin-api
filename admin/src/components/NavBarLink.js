/**
 * Created by pdiouf on 2017-03-23.
 */
/**
 * Created by pdiouf on 2017-03-23.
 */

import React from 'react';
const NavBarLink =(props)=>(

            <a onClick={props.controlFunc}key={props.key} href={props.url}>
                <span className={props.style}></span>
                {props.text}
            </a>
        );


NavBarLink.PropTypes={
    controlFunc: React.PropTypes.func,
}

export default NavBarLink;