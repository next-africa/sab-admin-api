/**
 * Created by pdiouf on 2017-03-12.
 */
import React from 'react';
const Select= (props) => (
    <div className="form-group">
        <label className="university-form"></label>
        <textarea
            className="form-input"
            style={props.resize ? null : {resize: 'none'}}
            name={props.name}
            rows={props.rows}
            value={props.content}
            onChange={props.controlFunc}
            placeholder={props.placeholder}/>
    </div>
);


Select.propTypes={
    title: React.PropTypes.string.isRequired,
    rows: React.PropTypes.number.isRequired,
    name: React.PropTypes.string.isRequired,
    content: React.PropTypes.string.isRequired,
    resize: React.PropTypes.bool,
    placeholder: React.PropTypes.string,
    controlFunc: React.PropTypes.func.isRequired
};

export default TextArea;