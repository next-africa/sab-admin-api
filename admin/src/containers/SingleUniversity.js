/**
 * Created by pdiouf on 2017-04-03.
 */
import React, {Component} from 'react' ;
import NotFoundPage from '../components/NotFoundPage'
class SingleUniversity extends Component{

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
                    <div className="U-address">
                        <div>icone</div>
                        <div>
                            <p> {university.address.line}</p>
                            <p> {university.address.city} ({university.address.state}),{university.address.code}</p>
                        </div>
                    </div>
                    <div className="U-links">
                        <p> Languages : {university.selectedLanguages}</p>
                        <p> Website : {university.website}</p>
                        <p> ProgramListLink : {university.programListLink}</p>
                        <p> Languages : {university.selectedLanguages}</p>
                    </div>
                    <div className="U-tuition">
                        <p> Tuition link : {university.tuition.link}</p>
                        <p> Tuition amount : {university.tuition.amount}</p>
                    </div>
                    <div className="U-domains">

                    </div>
                </div>
                <button
                    className="btn btn-link float-left"
                    onClick={(event) => {this.props.setCurrentPage(event, {page:'universities'});}}>back to list</button>
            </div>
        )
    }
}

export default SingleUniversity;
