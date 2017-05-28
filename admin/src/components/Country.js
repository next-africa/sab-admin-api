/**
 * Created by pdiouf on 2017-05-13.
 */
import Select from './Select'
import React from 'react'
import {
    createFragmentContainer,
    graphql,
} from 'react-relay';


class Country extends React.Component{


    handleCountrySelection(e){
        this.setState({countryCode:e.target.value}, () => console.log('Country  :',this.state.countryCode));
    }


    render(){
        return(
            <Select
                name={'countryCode'}
                selectedOption={this.state.countryCode}
                controlFunc={this.handleCountrySelection}
                options={["CA","US","SN","CD"]}
                placeholder={'Choose a country code'}/>
        )
    }
}
export default createFragmentContainer(Country,{
    Country: graphql`
        fragment Country on Country(code:"ca"){
            id
            properties{
                code
                name
            }
        }
    `
})