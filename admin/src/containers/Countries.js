/**
 * Created by pdiouf on 2017-05-13.
 */
/**
 * Created by pdiouf on 2017-05-13.
 */
import Select from '../components/Select'
import React from 'react'
import Relay from 'react-relay'
// import {
//     createFragmentContainer,
//     graphql,
//     Relay,
// } from 'react-relay';


class Countries extends React.Component{


    handleCountrySelection(e){
        this.setState({countryCode:e.target.value}, () => console.log('Country  :',this.state.countryCode));
    }


    render(){
        var codes= [];
        console.log(this.props.countries);
        // this.props.countries.reduce(function(a,b,i,tab){
        //     codes[i] = b.properties.code
        // },0)
        return(
            <Select
                name={'countryCode'}
                selectedOption={this.state.countryCode}
                controlFunc={this.handleCountrySelection}
                options={codes}
                placeholder={'Choose a country code'}/>
        )
    }
}
export default Relay.createContainer(Countries,{
    // Countries: graphql`
    //     fragment  on countries{
    //         properties{
    //             code
    //         }
    //     }
    // `
    initialVariables: {
        code: "ca"
    },
    fragments: {
        countries : () => Relay.QL`
            fragment on Countries{
                id
                properties{
                    code
                    name
                }
                
            }
        `
    }
})