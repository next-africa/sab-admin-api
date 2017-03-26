/**
 * Created by pdiouf on 2017-03-11.
 */
import React, {Component} from 'react';
import SingleInput from '../components/SingleInput'
import CheckboxOrRadioGroup from '../components/CheckboxOrRadioGroup'
import Address from '../components/Address'
import Tuition from '../components/Tuition'
import Modal from '../../node_modules/react-bootstrap/lib/Modal'
import Button from '../../node_modules/react-bootstrap/lib/Button'
var app = {};
app.cards = [{
    "key":0,
    "universityName":  "Ulaval",
    "domains": [
        {
            "key":  " La clef du category/domaine'",
            "tuition": {
                "link": "Le lien vers la page sur les frais du programme",
                "amount": 5000
            },
            "descriptionLink": "Le lien vers la page du domaine",
            "languages": ["language 1", "language 2"]
        },
        {
            "key":  " La clef du category/domaine'",
            "tuition": {
                "link": "Le lien vers la page sur les frais du programme",
                "amount": 17000
            },
            "descriptionLink": "Le lien vers la page du domaine",
            "languages": [ "language 2"]
        }
    ],
    "languages": ["fr", "en"],
    "selectedLanguages":["en"],
    "website": "www.ulaval.ca",
    "programListLink": "",
    "address":  {
        "line": "2325 ",
        "city": "Rue de la vie etudiante ",
        "state": "Quebec ",
        "Postal code": "g1v 3x8"
    },
    "tuition": {
        "link": "Lien vers la description des frais",
        "amount": 55000
    }
}
]
app.dirty = false;
app.nextKey = 3;

var AddressTmp= {"line" :'', "city":'',"state":'', "code":''};
var TuitionTmp = {"link":'', "amount":''} ;
// A card item we want to display.
var CardItem = React.createClass({
    createCard: function(item) {
        var keyId = 'cardkey-' + item.key;
        return (
            <div id={keyId} className="card-item">
                <div className={item.color}>
                    <div className="card-info">
                        <div className="card-name">{item.universityName}</div>
                        <div className="card-desc">{item.website}</div>
                        <div className="card-desc">{item.address.state}</div>
                    </div>
                </div>
                <div className="clear"></div>
                <div className="card-delete" onClick={this.deleteCard}>delete</div>
            </div>
        );
    },
    // Delete the card.
     deleteCard: function(e) {
     e.preventDefault();
     console.log(e.target.parentElement.attributes.id.value.split('-')[1]);
     var keyId = e.target.parentElement.attributes.id.value.split('-')[1];
     var newArr = app.cards.filter(app.cards, function(card) { return (card.key != keyId); });
     document.getElementById('#cardkey-' + keyId).animate({
     opacity: 0,
     left: "+=1000",
     }, 1000, function() {
     (this).hide(); // Ok.. so I don't really delete the card...
     });
     },

     // Render the card.
    render: function() {
        return (
            <div className="cards">
                {this.props.cards.map(this.createCard)}
            </div>
        );
    }

});
// The card manager/holder.
var CardManager = React.createClass({
    getInitialState: function() {
        return {cards: []};
    },

    // When our component mounts we should update the cards with the initial ones.
    componentDidMount: function() {
        var _self = this;
        this.setState({cards: app.cards});

        // We'll just cheat a bit and set an interval to watch changes.
        setInterval(function() {
            if (app.dirty) {
                app.dirty = false;
                _self.setState({cards: app.cards});
            }
        }, 50);
    },

    // Render our cycle of cards.
    render: function() {
        return (
            <div className="card-cycle">
                <CardItem cards={this.state.cards} />
            </div>
        );
    }
});
const UniversityForm = React.createClass({
    getInitialState(){

        return {
            key: '',
            universityName: '',
            domains: [],
            languages: [],
            selectedLanguages: [],
            website: '',
            programListLink: '',
            address: {},
            tuition: {},
            showModal: false
        };


        this.handleFormSubmit= this.handleFormSubmit.bind(this);
        this.handleClearForm= this.handleClearForm.bind(this);
        this.handleUniversityNameChange= this.handleUniversityNameChange.bind(this);
        this.handleLangSelection= this.handleLangSelection.bind(this);
        this.handleUniversityLinkChange= this.handleUniversityLinkChange.bind(this);
        this.handleProgramListLink= this.handleProgramListLink.bind(this);
        this.handleAddressChange= this.handleAddressChange.bind(this);
        this.handleTuitionChange= this.handleTuitionChange.bind(this);
    },

    componentDidMount(){
        fetch('./db.json')
            .then(res=> res.json())
            .then(data=>{
                this.setState({
                    key:'',
                    universityName: data.name,
                    domains: data.domains,
                    languages: data.languages,
                    selectedLanguages:data.selectedLanguages,
                    website: data.website,
                    programListLink: data.programListLink,
                    address: data.address,
                    tuition:data.tuition,
                    showModal: ''
                });
            });

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
    },

    handleFormSubmit(e){
        e.preventDefault();
        //Create University
        var colors = ['blue', 'purple', 'red', 'yellow'];

        const formPayload= {
            key:(app.nextKey++),
            universityName: this.state.universityName,
            domains: this.state.domains,
            selectedLanguages: this.state.selectedLanguages,
            website: this.state.website,
            programListLink: this.state.programListLink,
            address: this.state.address,
            tuition: this.state.tuition,
            color:('card-color'+' '+ colors[Math.floor((Math.random()*3)+1)])

        };
        app.cards.push(formPayload);
        console.log('TODO==> Post Request:', app.cards)
        this.handleClearForm(e);
        app.dirty = true;
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
        AddressTmp[name] = value
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
            <CardManager />
            <Button bsStyle="primary"
                    bsSize="large"
                    onClick={this.open}>Add university</Button>
                <Modal show={this.state.showModal} onHide={this.close}>
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
                </Modal>
            </div>

        );
    },
})
export default UniversityForm;