/**
 * Created by pdiouf on 2017-03-25.
 */
import React from 'react'
import UniversityPreview from '../components/UniversityPreview'
import {Button} from 'react-bootstrap'
import If from '../components/If'
var app =  {};

// Some initial universities to start with.
const styles = {
    divider: {
        margin: '8px 0',
        height: 1,
        backgroundColor: '#757575',
    },
}
app.dirty = false;

// The card manager/holder.
class CardManager  extends React.Component{
    constructor(props)
    {
        super(props)
        this.state = {
            universities: []
        }
    }
    // When our component mounts we should update the universities with the initial ones.
    componentDidMount() {
        var _self = this;
        this.setState({universities: this.props.universitiesList});

        // We'll just cheat a bit and set an interval to watch changes.
        setInterval(function() {
            if (app.dirty) {
                app.dirty = false;
                _self.setState({universities: this.props.universitiesList});
            }
        }, 500);
    }


    // Render our cycle of universities.
    render() {
        return (

            <div className="card-cycle">
                <UniversityPreview universitiesList={this.props.universitiesList} setCurrentPage={this.props.setCurrentPage}/>
            </div>
        );
    }
};


// The card application.
const Universities =(props) => (


                <div>
                    <div className="btnNewUniversity">
                        <Button bsStyle="primary"
                                bsSize="large"
                                className="pull-right"
                                onClick={(event) => {props.setCurrentPage(event, {page:'newUniversity'});}}>
                            New University

                        </Button>
                    </div>
                    <h3>Universities</h3>
                    <div style={styles.divider}></div>

                    <If items={props.universitiesList}>

                        <CardManager  universitiesList={props.universitiesList} setCurrentPage={props.setCurrentPage}/>
                    </If>


                </div>


);

Universities.propTypes= {
    setCurrentPage : React.PropTypes.func,
}
export default Universities;
