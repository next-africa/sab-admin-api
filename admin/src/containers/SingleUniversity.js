/**
 * Created by pdiouf on 2017-04-03.
 */
import React from 'react' ;
import NotFoundPage from '../components/NotFoundPage'
import {Grid, Col, Row} from 'react-bootstrap'
import DomainForm from '../containers/DomainForm'
import {Button} from 'react-bootstrap'
import {Modal} from 'react-bootstrap'

const styles = {
    divider: {
        margin: '8px 0',
        height: 1,
        backgroundColor: '#757575',
    },
}
const SingleUniversity = React.createClass({
    getInitialState() {
        return {
            showModal: false
        }

    },

    close(){
                this.setState({showModal:false});
    },
    open(){
           this.setState({showModal:true});
    },

    render(){
        const id = this.props.currentPageId;
        console.log("id",id);
        const universities = this.props.universitiesList;
        const university = universities.filter((university) =>university.id === id)[0];

        if(!university){
            return <NotFoundPage/>
        }
        return (
            <div>
                <h4 className="page-header">{university.name}</h4>
                <div className="U-info">
                    <Grid>
                        <Row className="show-grid">
                            <Col xs={6} md={4}>
                                <span>ADRESSE</span>
                                <h5 style={styles.divider}></h5>
                                <div className="U-address infoLeft">
                                    <div className="glyphicon glyphicon-map-marker"></div>
                                    <div className="AddrInfo ">
                                        <p> {university.address.line}</p>
                                        <p> {university.address.city} ({university.address.state}),{university.address.code}</p>
                                    </div>
                                </div>
                            </Col>
                            <Col xs={6} md={4}>
                                <span>INFOS</span>
                                <h5 style={styles.divider}></h5>
                                <div className="U-links infoLeft">
                                    <p><strong> Languages </strong>: {university.selectedLanguages}</p>
                                    <p> Website : {university.website}</p>
                                    <p> ProgramListLink : {university.programListLink}</p>
                                    <p> Languages : {university.selectedLanguages}</p>

                                    <p> Tuition link : {university.tuition.link}</p>
                                    <p> Tuition amount : {university.tuition.amount}</p>
                                </div>
                            </Col>
                            <Col xs={6} md={4}>
                                <span>DOMAINS</span>
                                <h5 style={styles.divider}></h5>
                                <div className="U-domains">
                                    <div className="btnNewUniversity">
                                        <Button bsStyle="primary"
                                                bsSize="large"
                                                className="pull-right"
                                                onClick={this.open}
                                                >
                                            <span className="glyphicon glyphicon-plus"></span>

                                        </Button>
                                    </div>
                                </div>
                                <Modal show={this.state.showModal} onHide={this.close}>
                                    <DomainForm/>
                                </Modal>
                            </Col>
                        </Row>
                    </Grid>


                </div>
                <button
                    className="btn btn-link float-right"
                    onClick={(event) => {this.props.setCurrentPage(event, {page:'universities'});}}>back to the list
                </button>


            </div>
        )
    }
});

export default SingleUniversity;
