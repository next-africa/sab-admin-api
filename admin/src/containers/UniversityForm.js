/**
 * Created by pdiouf on 2017-03-11.
 */
import React, {Component} from 'react';
import SingleInput from '../components/SingleInput'


class UniversityForm extends  Component{
    constructor(props) {
        super(props);
        this.state={
            universityName:  '',
            studyPrograms: [],
            languages: [],
            website: '',
            tuitionLink: '',
            programLink: '',
            address: '',
            city: '',
            state: '',
            minTuition: '',
            maxTuition: ''
        };

        this.handleFormSubmit= this.handleFormSubmit.bind(this);
        this.handleClearForm= this.handleClearForm.bind(this);
        this.handleUniversityNameChange= this.handleUniversityNameChange.bind(this);
    }

    componentDidMount(){
        fetch('./db.json')
            .then(res=> res.json())
            .then(data=>{
                this.setState({
                    universityName: data.name,
                    studyPrograms: data.studyPrograms,
                    languages: data.languages,
                    website: data.website,
                    tuitionLink: data.tuitionLink,
                    programLink: data.programLink,
                    address: data.address,
                    city: data.city,
                    state: data.state,
                    minTuition: data.min,
                    maxTuition: data.maxTuition
                });
            });
    }

    handleClearForm(e){
        e.preventDefault();
        this.setState({
            universityName:'',
            languages: '',
            website: '',
            tuitionLink: '',
            programLink: '',
            address: '',
            city: '',
            state: '',
            minTuition: '',
            maxTuition: ''
        });
    }

    handleFormSubmit(e){
        e.preventDefault();
        const formPayload= {
            universityName: this.state.universityName,
            studyPrograms: this.state.studyPrograms,
            languages: this.state.languages,
            website: this.state.website,
            tuitionLink: this.state.tuitionLink,
            programLink: this.state.programLink,
            address: this.state.address,
            city: this.state.city,
            state: this.state.state,
            minTuition: this.state.min,
            maxTuition: this.state.maxTuition

        };
        console.log('TODO==> Post Request:', formPayload)
        this.handleClearForm(e);
    }

    handleUniversityNameChange(e){
        this.setState({universityName: e.target.value}, () => console.log('name:', this.state.universityName));
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
                    content={this.state.UniversityName}
                    placeholder={'Type the name of the university'}/>

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