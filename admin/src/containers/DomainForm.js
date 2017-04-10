/**
 * Created by pdiouf on 2017-04-08.
 */
import React, {Component} from 'react'
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'


var id = 0;
class DomainForm extends Component {

    constructor(props) {
        super(props);
        this.state ={
            id:0,
            name:'',
            description:'',
            tuition:{link:'', amount:''},
            languages: ['en','fr','de','es' ],
            selectedLanguages: ['en','fr'],
            TuitionTmp : {"link":'', "amount":''}

        };
        this.handleClearForm = this.handleClearForm.bind(this);
        this.handleFormSubmit = this.handleFormSubmit.bind(this);
        this.handleDescriptionChange = this.handleDescriptionChange.bind(this);
        this.handleTuitionChange = this.handleTuitionChange.bind(this);
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleLangSelection = this.handleLangSelection.bind(this);
    }
    handleClearForm(e){
        e.preventDefault();
        this.setState({
            name:'',
            description:'',
            selectedLanguages:[],
            tuition: {link:'', amount:''},
        });
    }
    handleFormSubmit(e){
        e.preventDefault();
        const formPayload= {
            id:id++,
            name:this.state.name,
            description: this.state.description,
            selectedLanguages: this.state.selectedLanguages,
            tuition: this.state.tuition,


        };

        this.props.domains.push(formPayload);
        console.log('TODO==> Post Request:', this.props.domains)
        this.handleClearForm(e);

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
    handleNameChange(e){
        this.setState({name:e.target.value}, () => console.log('name :', this.state.name));
    }
    handleDescriptionChange(e){
        this.setState({description:e.target.value}, () => console.log('description:', this.state.description));
    }
    handleTuitionChange(e) {
        const name = e.target.name;
        const value = e.target.value;
        this.state.TuitionTmp[name] = value;
        this.setState({tuition: this.state.TuitionTmp}, () => console.log('tuition:', this.state.tuition));
    }
    render(){
        return (
            <form className="container" onSubmit={this.handleFormSubmit}>
                <h5> All inputs are required</h5>
                <SingleInput
                    inputType={'text'}
                    title={'name'}
                    name={'name'}
                    controlFunc={this.handleNameChange}
                    content={this.state.name}
                    placeholder={'Type the name of the domain'}/>
                <SingleInput
                    inputType={'text'}
                    title={'description'}
                    name={'description'}
                    controlFunc={this.handleDescriptionChange}
                    content={this.state.description}
                    placeholder={'Type the link of the description'}/>

                <CheckboxOrRadioGroup
                    title={"Les langues d'enseignements"}
                    type={'checkbox'}
                    setName={'languages'}
                    options={this.state.languages}
                    controlFunc={this.handleLangSelection}
                    selectedOptions={this.state.selectedLanguages}/>

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
                            inputType={'text'}
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
export default DomainForm;