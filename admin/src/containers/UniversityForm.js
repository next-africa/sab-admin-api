/**
 * Created by pdiouf on 2017-03-11.
 */
import React from 'react';
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'
import Address from '../components/Address'
import Tuition from '../components/Tuition'
import Modal from '../../node_modules/react-bootstrap/lib/Modal'
import Button from '../../node_modules/react-bootstrap/lib/Button'
import Select from '../components/Select'
import Universities from './Universities'
var AddressTmp= {"line" :'', "city":'',"state":'', "code":''};
var TuitionTmp = {"link":'', "amount":''} ;
var universitiesList = [];
// A card item we want to display.
var If = React.createClass({
    render:function(){
        if(this.props.numberOfUniversities){
            return this.props.children;
        }
        else{
            return false;
        }
    }
});
const UniversityForm = React.createClass({
    getInitialState(){
        return {
            id: 0,
            universityName: '',
            domains: '',
            languages: ['en','fr','de','es' ],
            selectedLanguages: [],
            website: '',
            programListLink: '',
            address: '',
            tuition: {},
            showModal: false,
            countryCode:''
        };
    },

    handleCountrySelection(e){
        this.setState({countryCode:e.target.value}, () => console.log('Country  :',this.state.countryCode));
    },
    handleUniversityLinkChange(e){
        this.setState({website:e.target.value}, () => console.log('website link :',this.state.website));
    },

    handleLangSelection(e){
        const newSelection= e.target.value;
        let newSelectionArray;
        if(this.state.selectedLanguages.indexOf(newSelection)> -1){
            newSelectionArray= this.state.selectedLanguages.filter(s => s !== newSelection)
        }else{
            newSelectionArray= [...this.state.selectedLanguages, newSelection];
        }
        this.setState({selectedLanguages:newSelectionArray}, () => console.log('universite Language:', this.state.selectedLanguages));
    },

    handleClearForm(e){
        e.preventDefault();
        this.setState({
            universityName:'',
            selectedLanguages: [],
            website: '',
            programListLink: '',
            address:{},
            tuition: {}
        });
        this.close();
    },

    handleFormSubmit(e){
        e.preventDefault();
        //Create University
        var colors = ['blue', 'purple', 'red', 'yellow'];

        const formPayload= {
            id:this.state.id++,
            universityName: this.state.universityName,
            domains: this.state.domains,
            selectedLanguages: this.state.selectedLanguages,
            website: this.state.website,
            programListLink: this.state.programListLink,
            address: this.state.address,
            tuition: this.state.tuition,
            color:("card-color "+colors[Math.floor((Math.random()*3)+1)]),
            countryCode:this.state.countryCode

        };

        universitiesList.push(formPayload);
        console.log('TODO==> Post Request:', universitiesList)
        this.handleClearForm(e);

        this.close();
    },
    handleProgramListLink(e){
        this.setState({programListLink:e.target.value}, () => console.log('Program list link :', this.state.programListLink));
    },
    handleUniversityNameChange(e){
        this.setState({universityName: e.target.value}, () => console.log('University name:', this.state.universityName));
    },

    handleAddressChange(e) {
        const name = e.target.name;
        const value = e.target.value;
        AddressTmp[name] = value;
        this.setState({address:AddressTmp}, () => console.log('address:', this.state.address));

    },
    handleTuitionChange(e){
        const name = e.target.name;
        const value = e.target.value;
        TuitionTmp[name] = value;
        this.setState({tuition:TuitionTmp}, () => console.log('Tuition ',this.state.tuition));

    },
    close(){
        this.setState({showModal:false});
    },
    open(){
        this.setState({showModal:true});
    },

    render(){
        return (
            <div>
                <If numberOfUniversities={universitiesList.length}>
                    <Universities  universitiesList={universitiesList}/>
                </If>

                <Button bsStyle="primary"
                        bsSize="large"
                        onClick={this.open}>Add university
                </Button>
                <Modal show={this.state.showModal} onHide={this.close}>
                    <form className="container" onSubmit={this.handleFormSubmit}>
                        <h5> Ajouter une université </h5>
                        <Select
                            name={'countryCode'}
                            selectedOption={this.state.countryCode}
                            controlFunc={this.handleCountrySelection}
                            options={["CA","US","SN","CD"]}
                            placeholder={'Choose a country code'}/>
                        <SingleInput
                            inputType={'text'}
                            title={'University name'}
                            name={'name'}
                            controlFunc={this.handleUniversityNameChange}
                            content={this.state.universityName}
                            placeholder={'Type the name of the university'}/>
                        <SingleInput
                            inputType={'text'}
                            title={'University website link'}
                            name={'website'}
                            controlFunc={this.handleUniversityLinkChange}
                            content={this.state.website}
                            placeholder={'Type the website link of the university'}/>
                        <SingleInput
                            inputType={'text'}
                            title={'link of the  programs list'}
                            name={'programListLink'}
                            controlFunc={this.handleProgramListLink}
                            content={this.state.programListLink}
                            placeholder={'Program list link'}/>
                        <CheckboxOrRadioGroup
                            title={"Les langues d'enseignements"}
                            type={'checkbox'}
                            setName={'languages'}
                            options={this.state.languages}
                            controlFunc={this.handleLangSelection}
                            selectedOptions={this.state.selectedLanguages}/>
                        <Address
                            title="Address of university"
                            items={["line","city","state", "postalCode"]}
                            controlLineFunc={this.handleAddressChange}
                            controlCityFunc={this.handleAddressChange}
                            controlSateFunc={this.handleAddressChange}
                            controlCodeFunc={this.handleAddressChange}
                            contentLine={this.state.address['line']}
                            contentCity={this.state.address['city']}
                            contentState={this.state.address['state']}
                            contentCode={this.state.address['postalCode']}
                        />
                        <Tuition
                            title="Tuition of university"
                            items={["link","amount"]}
                            controlLinkFunc={this.handleTuitionChange}
                            controlAmountFunc={this.handleTuitionChange}
                            contentLink={this.state.tuition['link']}
                            contentAmount={this.state.tuition['amount']}
                        />

                        <input
                            type="submit"
                            className="btn btn-primary float-right"
                            value="Add"/>

                        <button
                            className="btn btn-link float-left"
                            onClick={this.handleClearForm}>Cancel</button>

                    </form>
                </Modal>
            </div>

        );
    },
})
export default UniversityForm;