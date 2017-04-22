import React from 'react'
import {Accordion, Panel} from 'react-bootstrap'

var DomainPreview = React.createClass({

    createItem : function(item){
        console.log("item:", item);
        var keyId = 'domain-' + item.id;
        return(
            <Panel header={item.name} key={keyId} id={keyId} eventKey={keyId}>
                <p> Languages : {item.selectedLanguages}</p>
                <p> Description  : {item.description}</p>
                <p> Tuition link : {item.tuition.link}</p>
                <p> Tuition amount : {item.tuition.amount}</p>
            </Panel>
        );

    },

    render: function(){
        return(
            <div>
                <Accordion>
                     {this.props.domains.map(this.createItem)}
                </Accordion>
            </div>

        )

    }
})


export  default  DomainPreview;