/**
 * Created by pdiouf on 2017-03-12.
 */
import React, {Component} from 'react'
const Button= (props) =>(
    <input
        className={props.bsClass}
        size={props.bsSize}
        style={props.bsStyle}
        disabled={props.disabled}
        onClick={props.controlFunc}
        type={props.type}

    />

);


Button.propTypes={
    bsClass:React.PropTypes.string.isRequired,
    bsSize:React.PropTypes.oneOf(['lg','large','sm','small','xs']),
    bsStyle:React.PropTypes.oneOf(['success', 'warning', 'danger', 'info', 'default', 'primary', 'link']),
    disable:React.PropTypes.bool.isRequired,
    controlFunc:React.PropTypes.func.isRequired,
    type:React.PropTypes.oneOf(['button','reset','submit'])
}

export default Button;