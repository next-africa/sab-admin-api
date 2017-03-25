/**
 * Created by pdiouf on 2017-03-23.
 */
/**
 * Created by pdiouf on 2017-03-23.
 */

import React from 'react';
import NavBarLink from './NavBarLink'
const  NavBarItem = (props) => (
            <li >
                <NavBarLink
                    style={props.style}
                    url={props.url}
                    text={props.text}
                    controlFunc={props.controlFunc}
                />

            </li>
        );


NavBarItem.PropTypes={
    controlFunc: React.PropTypes.func,
}

export default NavBarItem;