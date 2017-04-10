/**
 * Created by pdiouf on 2017-04-09.
 */
import React from 'react'
import DomainPreview from '../components/DomainPreview'

var Domains = React.createClass({
  render : function(){
        return(
            <DomainPreview domains={this.props.domains}/>
        );
    }
});

export default Domains;