/**
 * Created by pdiouf on 2017-03-11.
 */
import React, {Component} from 'react';
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'
import Select from '../components/Select'
var AddressTmp= {"line" :'', "city":'',"state":'', "code":''};
var TuitionTmp = {"link":'', "amount":''} ;
class UniversityForm extends Component {
    constructor(props) {
        super(props);
        this.state ={
            name: '',
            domains: '',
            languages: ['en','fr','de','es' ],
            selectedLanguages: ['fr'],
            website: '',
            programListLink: '',
            address: {line :'', city:'',state:'', code:''},
            tuition: {"link":'', "amount":0},
            showModal: false,
            countryCode:'',
            setCurrentPage :null
        };

        this.handleProgramListLink = this.handleProgramListLink.bind(this);
        this.handleCountrySelection = this.handleCountrySelection.bind(this);
        this.handleUniversityLinkChange = this.handleUniversityLinkChange.bind(this);
        this.handleLangSelection = this.handleLangSelection.bind(this);
        this.handleClearForm = this.handleClearForm.bind(this);
        this.handleFormSubmit = this.handleFormSubmit.bind(this);
        this.handleProgramListLink = this.handleProgramListLink.bind(this);
        this.handleUniversityNameChange = this.handleUniversityNameChange.bind(this);
        this.handleAddressChange = this.handleAddressChange.bind(this);
        this.handleTuitionChange = this.handleTuitionChange.bind(this);
    }

    setCurrentPage(event, { page, props }) {
        if (event) event.preventDefault();
        this.setState({ currentPage: page, currentPageProps: props });
    }
    handleCountrySelection(e){
        this.setState({countryCode:e.target.value}, () => console.log('Country  :',this.state.countryCode));
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
            name:'',
            selectedLanguages: [],
            website: '',
            programListLink: '',
            address:{},
            tuition: {}
        });
    }

    handleFormSubmit(e){
        e.preventDefault();
        //Create University
        var colors = ['blue', 'purple', 'red', 'yellow'];

        const formPayload = {
            name: this.state.name,
            languages: this.state.selectedLanguages,
            website: this.state.website,
            programListLink: this.state.programListLink,
            address: this.state.address,
            tuition: this.state.tuition
            //color:("card-color "+colors[Math.floor((Math.random()*3)+1)]),
            //countryCode:this.state.countryCode

        };

        fetch("/api/countries/ca/universities", {
            headers:{
                'content-type': 'application/json'
            },
            method: "POST",
            mode:"cors",
            credentials: "same-origin",
            body:JSON.stringify({
                name:this.state.name,
                languages:this.state.languages,
                website: this.state.website,
                programListLink:this.state.programListLink,
                address:this.state.address,
                tuition:this.state.tuition
            })

        });
        console.log("formJson:", formPayload);
        console.log('TODO==> Post Request:', this.props.universitiesList)
        this.handleClearForm(e);
        this.props.setCurrentPage(event, {page:'universities'});

    }
    handleProgramListLink(e){
        this.setState({programListLink:e.target.value}, () => console.log('Program list link :', this.state.programListLink));
    }
    handleUniversityNameChange(e){
        this.setState({name: e.target.value}, () => console.log('University name:', this.state.name));
    }
    handleAddressChange(e) {
        const name = e.target.name;
        const value = e.target.value;
        AddressTmp[name] = value;
        this.setState({address: AddressTmp}, () => console.log('address:', this.state.address));
    }

    handleTuitionChange(e) {
        const name = e.target.name;
        const value = e.target.value;
        TuitionTmp[name] = value;
        this.setState({tuition: TuitionTmp}, () => console.log('tuition:', this.state.tuition));
    }


    render(){
        return (
                    <form className="container" onSubmit={this.handleFormSubmit}>
                        <h5> All inputs are required</h5>
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
                        <div>
                            <label className="university-form">University's adresse</label>
                        <div className="address_items">

                            <SingleInput
                                inputType={'text'}
                                title={'line'}
                                name={'line'}
                                controlFunc={this.handleAddressChange}
                                content={this.state.address['line']}
                                placeholder={'Type the line'}/>
                            <SingleInput
                                inputType={'text'}
                                title={'city'}
                                name={'city'}
                                controlFunc={this.handleAddressChange}
                                content={this.state.address['city']}
                                placeholder={'Type the city'}/>

                            <SingleInput
                                inputType={'text'}
                                title={'state'}
                                name={'state'}
                                controlFunc={this.handleAddressChange}
                                content={this.state.address['state']}
                                placeholder={'Type the state'}/>

                            <SingleInput
                                inputType={'text'}
                                title={'code'}
                                name={'code'}
                                controlFunc={this.handleAddressChange}
                                content={this.state.address['code']}
                                placeholder={'Type the Postal code'}/>
                        </div>
                        </div>
                        <div>
                            <label className="university-form">Tuition of university</label>
                            <div className="tuitions_items">

                                <SingleInput
                                    inputType={'text'}
                                    title={'link'}
                                    name={'link'}
                                    controlFunc={this.handleTuitionChange}
                                    content={this.state.tuition['link']}
                                    placeholder={'Type the link'}/>
                                <SingleInput
                                    inputType={'number'}
                                    title={'amount'}
                                    name={'amount'}
                                    controlFunc={this.handleTuitionChange}
                                    content={this.state.tuition['amount']}
                                    placeholder={'Type the amount'}/>
                            </div>
                        </div>
                        <input
                            type="submit"
                            className="btn btn-primary float-right"
                            value="Add University"
                        />

                        <button
                            className="btn btn-link float-left"
                            onClick={this.handleClearForm}>Cancel</button>

                    </form>

        );
    }
}
export default UniversityForm;