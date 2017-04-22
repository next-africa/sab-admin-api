/**
 * Created by pdiouf on 2017-04-09.
 */
import React from 'react'
var If = React.createClass({
    render:function(){
        if(this.props.items.length){
            return this.props.children;
        }
        else{
            return false;
        }
    }
});
If.propTypes={
    items:React.PropTypes.array.isRequired
}
export default If;