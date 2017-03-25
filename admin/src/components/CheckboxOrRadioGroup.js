/**
 * Created by pdiouf on 2017-03-12.
 */
import React from 'react'

const CheckboxOrRadioGroup=(props) => (
    <div>
        <label className="university-form">{props.title}</label>
        <div className="checkbox-group">
            {props.options.map(option => {
                return(
                    <label key={option} className="university-form capitalize">
                        <input
                            className="form-checkbox"
                            name={props.setName}
                            onChange={props.controlFunc}
                            value={option}
                            checked={props.selectedOptions.indexOf(option) > -1}
                            type={props.type}/> {option}
                    </label>
                );

            })}
        </div>
    </div>
);

CheckboxOrRadioGroup.propTypes = {
    title: React.PropTypes.string.isRequired,
    type: React.PropTypes.oneOf(['checkbox', 'radio']).isRequired,
    setName: React.PropTypes.string.isRequired,
    options: React.PropTypes.array.isRequired,
    selectedOptions: React.PropTypes.array,
    controlFunc: React.PropTypes.func.isRequired
};

export default  CheckboxOrRadioGroup;

