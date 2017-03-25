/**
 * Created by pdiouf on 2017-03-18.
 */
/**
 * Created by pdiouf on 2017-03-12.
 */
import React from 'react'
import SingleInput from './SingleInput'

const Address=(props)=> (
    <div>
        <label className="university-form">{props.title}</label>
        <div className="address_items">
                    <SingleInput
                        inputType={'text'}
                        title={props.items[0]}
                        name={props.items[0]}
                        controlFunc={props.controlLineFunc}
                        content={props.contentLine}
                        placeholder={'Type the line'}/>

                    <SingleInput
                        inputType={'text'}
                        title={props.items[1]}
                        name={props.items[1]}
                        controlFunc={props.controlCityFunc}
                        content={props.contentCity}
                        placeholder={'Type the city'}/>

                    <SingleInput
                        inputType={'text'}
                        title={props.items[2]}
                        name={props.items[2]}
                        controlFunc={props.controlSateFunc}
                        content={props.contentState}
                        placeholder={'Type the state'}/>

                    <SingleInput
                        inputType={'text'}
                        title={props.items[3]}
                        name={props.items[3]}
                        controlFunc={props.controlCodeFunc}
                        content={props.contentCode}
                        placeholder={'Type the Postal code'}/>


        </div>
    </div>
);

Address.propTypes= {
    title: React.PropTypes.string.isRequired,
    controlLineFunc: React.PropTypes.func.isRequired,
    controlCityFunc: React.PropTypes.func.isRequired,
    controlSateFunc: React.PropTypes.func.isRequired,
    controlCodeFunc: React.PropTypes.func.isRequired,
    items : React.PropTypes.array.isRequired,
    contentLine : React.PropTypes.string,
    contentCity : React.PropTypes.string,
    contentState: React.PropTypes.string,
    contentCode:  React.PropTypes.string,


}


export default  Address;
