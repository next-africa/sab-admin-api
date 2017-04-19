/**
 * Created by pdiouf on 2017-04-08.
 */

import SingleInput from '../components/SingleInput'
import TextArea from '../components/TextArea'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'

var TuitionTmp = {"link":'', "amount":''} ;
domains = [];
class DomainForm extends Component {
    constructor(props) {
        super(props);
        this.state ={

            description:'',
            tuition:'',
            languages:[],
            selectedLanguages:[]
        };
;
        this.handleLangSelection = this.handleLangSelection.bind(this);
        this.handleClearForm = this.handleClearForm.bind(this);
        this.handleFormSubmit = this.handleFormSubmit.bind(this);
        this.handleDescriptionChange = this.handleDescriptionChange.bind(this);
        this.handleTuitionChange = this.handleTuitionChange.bind(this);
    }

    setCurrentPage(event, { page, props }) {
        if (event) event.preventDefault();
        this.setState({ currentPage: page, currentPageProps: props });
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
            description:'',
            languages:[],
            selectedLanguages:[],
            tuition: {link:'', amount:''},
        });
    }

    handleFormSubmit(e){
        e.preventDefault();
        //Create University
        var colors = ['blue', 'purple', 'red', 'yellow'];

        const formPayload= {
            id:id++,
            description: this.state.description,
            selectedLanguages: this.state.selectedLanguages,
            tuition: this.state.tuition,


        };

        this.props.domains.push(formPayload);
        console.log('TODO==> Post Request:', this.props.domains)
        this.handleClearForm(e);
        this.props.setCurrentPage(event, {page:'universities'});

    }
    handleDescriptionChange(e){
        this.setState({description:e.target.value}, () => console.log('description:', this.state.description));
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
                <TextArea
                    title={'Add a description for this domain'}
                    rows={5}
                    resize={false}
                    content={this.state.description}
                    name={'description'}
                    controlFunc={this.handleDescriptionChange}
                    placeholder={'Please be thorough in your descriptions'} />

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