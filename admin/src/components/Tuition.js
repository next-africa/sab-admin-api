/**
 * Created by pdiouf on 2017-03-25.
 */

import React from 'react'
import SingleInput from './SingleInput'

const Tuition=(props)=> (
    <div>
        <label className="university-form">{props.title}</label>
        <div className="address_items">
            <SingleInput
                inputType={'text'}
                title={props.items[0]}
                name={props.items[0]}
                controlFunc={props.controlLinkFunc}
                content={props.contentLink}
                placeholder={'Type the line'}/>

            <SingleInput
                inputType={'text'}
                title={props.items[1]}
                name={props.items[1]}
                controlFunc={props.controlAmountFunc}
                content={props.contentAmount}
                placeholder={'Type the city'}/>
        </div>
    </div>
);

Tuition.propTypes= {
    title: React.PropTypes.string.isRequired,
    controlLinkFunc: React.PropTypes.func.isRequired,
    controlAmountFunc: React.PropTypes.func.isRequired,
    items : React.PropTypes.array.isRequired,
    contentLink : React.PropTypes.string,
    contentAmount : React.PropTypes.number,


}


export default  Tuition;