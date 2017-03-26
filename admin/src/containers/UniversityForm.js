/**
 * Created by pdiouf on 2017-03-11.
 */
import React, {Component} from 'react';
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'
import Address from '../components/Address'
import Tuition from '../components/Tuition'

var AddressTmp= {"line" :'', "city":'',"state":'', "code":''};
var TuitionTmp = {"link":'', "amount":''}
class UniversityForm extends  Component{
    constructor(props) {
        super(props);
          this.state={
            universityName:  '',
            domains: [],
            languages: [],
            selectedLanguages:[],
            website: '',
            programListLink: '',
            address: {},
            tuition: {}
        };

        this.handleFormSubmit= this.handleFormSubmit.bind(this);
        this.handleClearForm= this.handleClearForm.bind(this);
        this.handleUniversityNameChange= this.handleUniversityNameChange.bind(this);
        this.handleLangSelection= this.handleLangSelection.bind(this);
        this.handleUniversityLinkChange= this.handleUniversityLinkChange.bind(this);
        this.handleProgramListLink= this.handleProgramListLink.bind(this);
        this.handleAddressChange= this.handleAddressChange.bind(this);
        this.handleTuitionChange= this.handleTuitionChange.bind(this);
    }

    componentDidMount(){
        fetch('./db.json')
            .then(res=> res.json())
            .then(data=>{
                this.setState({
                    universityName: data.name,
                    domains: data.domains,
                    languages: data.languages,
                    selectedLanguages:data.selectedLanguages,
                    website: data.website,
                    programListLink: data.programListLink,
                    address: data.address,
                    tuition:data.tuition
                });
            });

    }

    handleUniversityLinkChange(e){
        this.setState({website:e.target.value}, () => console.log('website link :',this.state.website));
    }

    handleLangSelection(e){
        const newSelection= e.target.value;
        let newSelectionArray;
        if(this.state.selectedLanguages.indexOf(newSelection)> -1){
            newSelectionArray= this.state.selectedLanguages.filter(s => s !== newSelection)
        }else{
            newSelectionArray= [...this.state.selectedLanguages, newSelection];
        }
        this.setState({selectedLanguages:newSelectionArray}, () => console.log('universite Language:', this.state.selectedLanguages));
    }

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
    }

    handleFormSubmit(e){
        e.preventDefault();
        const formPayload= {
            universityName: this.state.universityName,
            domains: this.state.domains,
            selectedLanguages: this.state.selectedLanguages,
            website: this.state.website,
            programListLink: this.state.programListLink,
            address: this.state.address,
            tuition: this.state.tuition

        };
        console.log('TODO==> Post Request:', formPayload)
        this.handleClearForm(e);
    }
    handleProgramListLink(e){
        this.setState({programListLink:e.target.value}, () => console.log('Program list link :', this.state.programListLink));
    }
    handleUniversityNameChange(e){
        this.setState({universityName: e.target.value}, () => console.log('University name:', this.state.universityName));
    }

    handleAddressChange(e) {
        const name = e.target.name;
        const value = e.target.value;
        AddressTmp[name] = value
        this.setState({address:AddressTmp}, () => console.log('address:', this.state.address));

    }
    handleTuitionChange(e){
        const name = e.target.name;
        const value = e.target.value;
        TuitionTmp[name] = value;
        this.setState({tuition:TuitionTmp}, () => console.log('Tuition ',this.state.tuition));

    }
    render(){
        return (
            <form className="container" onSubmit={this.handleFormSubmit}>
                <h5> Ajouter une université </h5>
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
                    items={["line","city","state", "code"]}
                    controlLineFunc={this.handleAddressChange}
                    controlCityFunc={this.handleAddressChange}
                    controlSateFunc={this.handleAddressChange}
                    controlCodeFunc={this.handleAddressChange}
                    contentLine={this.state.address['line']}
                    contentCity={this.state.address['city']}
                    contentState={this.state.address['state']}
                    contentCode={this.state.address['Postal code']}
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
                    value="Submit"/>

                <button
                    className="btn btn-link float-left"
                    onClick={this.handleClearForm}>clear form</button>

            </form>

        );
    }

}
export default UniversityForm;