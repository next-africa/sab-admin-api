/**
 * Created by pdiouf on 2017-03-11.
 */
import React, {Component} from 'react';
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'
import Select from '../components/Select'
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
            addressLine :'',
            addressCity:'',
            addressState:'',
            addressCode:'',
            tuitionLink: '',
            tuitionAmount: 0,
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
        this.handleAddressLineChange = this.handleAddressLineChange.bind(this);
        this.handleAddressCityChange = this.handleAddressCityChange.bind(this);
        this.handleAddressStateChange = this.handleAddressStateChange.bind(this);
        this.handleAddressCodeChange = this.handleAddressCodeChange.bind(this);
        this.handleTuitionAmountChange = this.handleTuitionAmountChange.bind(this);
        this.handleTuitionLinkChange = this.handleTuitionLinkChange.bind(this);
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
        const formPayload = {
            name: this.state.name,
            languages: this.state.selectedLanguages,
            website: this.state.website,
            programListLink: this.state.programListLink,
            address: this.state.address,
            tuition: this.state.tuition

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
    handleAddressCityChange(e) {
        this.setState({addressCity: e.target.value}, () => console.log('City:', this.state.addressCity));
    }
    handleAddressLineChange(e) {
        this.setState({addressLine: e.target.value}, () => console.log('Line:', this.state.addressLine));
    }
    handleAddressStateChange(e) {
        this.setState({addressstate: e.target.value}, () => console.log('state:', this.state.addressState));
    }
    handleAddressCodeChange(e) {
        this.setState({addressCode: e.target.value}, () => console.log('Code:', this.state.addressCode));
    }
    handleTuitionAmountChange(e){
        this.setState({tuitionAmount:e.target.value}, () => console.log('tuitionAmount:', this.state.tuitionAmount));
    }

    handleTuitionLinkChange(e) {
        this.setState({tuitionLink:e.target.value}, () => console.log('tuitionAmount:', this.state.tuitionLink));
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
                                controlFunc={this.handleAddressLineChange}
                                content={this.state.addressLine}
                                placeholder={'Type the line'}/>
                            <SingleInput
                                inputType={'text'}
                                title={'city'}
                                name={'city'}
                                controlFunc={this.handleAddressCityChange}
                                content={this.state.addressCity}
                                placeholder={'Type the city'}/>

                            <SingleInput
                                inputType={'text'}
                                title={'state'}
                                name={'state'}
                                controlFunc={this.handleAddressStateChange}
                                content={this.state.addressState}
                                placeholder={'Type the state'}/>

                            <SingleInput
                                inputType={'text'}
                                title={'code'}
                                name={'code'}
                                controlFunc={this.handleAddressCodeChange}
                                content={this.state.addressCode}
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
                                    controlFunc={this.handleTuitionLinkChange}
                                    content={this.state.tuitionLink}
                                    placeholder={'Type the link'}/>
                                <SingleInput
                                    inputType={'number'}
                                    title={'amount'}
                                    name={'amount'}
                                    controlFunc={this.handleTuitionAmountChange}
                                    content={this.state.tuitionAmount}
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