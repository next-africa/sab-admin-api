import React, {Component} from 'react';
import '../node_modules/spectre.css/dist/spectre.min.css';
import './App.css';
import { Grid } from 'react-bootstrap';
import UniversityForm from './containers/UniversityForm';
import NavBarItem from './components/NavBarItem'
import ContentHeader from './containers/ContentHeader';
import SidebarContent from './containers/SidebarContent';
import NewUniversity from './Pages/NewUniversity';
import ViewUniversity from './Pages/ViewUniversity';
import Universities from './containers/Universities'
import  Sidebar  from 'react-sidebar';
import Relay from 'react-relay'
class App extends Component{
    constructor(props){
        super(props);
        this.state= {
            docked: true,
            open: false,
            transitions: true,
            touch: true,
            shadow: true,
            pullRight: false,
            touchHandleWidth: 20,
            dragToggleDistance: 30,
            currentPage:'universities',
            currentPageProps:null,
            currentPageId : 0,
            universitiesList: []
        }
    this.currentPage = this.currentPage.bind(this);
    this.setCurrentPage = this.setCurrentPage.bind(this);
    this.onSetDocked = this.onSetDocked.bind(this);
    this.toggleOpen = this.toggleOpen.bind(this);
    this.toggleDocked = this.toggleDocked.bind(this);
    this.generateItem = this.generateItem.bind(this);
    };
    componentDidMount(){
        let that = this
        fetch("/api/countries/ca/universities", {
            headers:{
                'content-type': 'application/json'
            },
            method: "GET",
            mode:"cors",
            credentials: "same-origin"

        })
            .then((res) => res.json())
            .then(function(data){
                that.setState({universitiesList:data.data});
                console.log("data",that.state.universitiesList)

            })
    }
    setCurrentPage(event, { page, props , id}) {
        if (event) event.preventDefault();
        this.setState({ currentPage: page, currentPageProps: props , currentPageId: id});
    }
    currentPage() {
        return {
            universities:<Universities universitiesList={this.state.universitiesList}/>,
            universityForm: <UniversityForm universitiesList={this.state.universitiesList}/>,
            newUniversity: <NewUniversity universitiesList={this.state.universitiesList}/>,
            viewUniversity: <ViewUniversity universitiesList={this.state.universitiesList} currentPageId={this.state.currentPageId}/>,
        }[this.state.currentPage];
    }
    componentWillMount() {
        const mql = window.matchMedia(`(min-width: 800px)`);
        mql.addListener(this.mediaQueryChanged);
        this.setState({mql: mql, docked: mql.matches});
    }
    componentWillUnmount() {
        this.state.mql.removeListener(this.mediaQueryChanged);
    }
    onSetOpen(open) {
        this.setState({open: open});
    }
    onSetDocked(docked) {
        this.setState({docked: !docked});
    }
    mediaQueryChanged() {
        this.setState({docked: this.state.mql.matches});
    }
    toggleOpen(e) {
        e.preventDefault();
        this.setState({open: !this.state.open});
    }
    toggleDocked(e) {
        e.preventDefault();
        this.setState({docked: !this.state.docked});
    }
    generateItem(item){
        return <NavBarItem controlFunc={this.toggleDocked}style={item.style} key={item.key} text={item.text} url={item.url}/>
    }
    generateUserInfos(infos){
        return <NavBarItem style={infos.style} key={infos.key} text={infos.text} url={infos.url}/>

    }

    render() {
        const data =  [

            {
                "key":"3",
                "text":"",
                "url":"#",
                "style":"glyphicon glyphicon-align-justify"
            },

        ]
       const userInfos = [
            {
                "key": "Admin",
                "text": "Admin",
                "url": "#",
                "style": "glyphicon glyphicon-user"
            },
            {
                "key": "login",
                "text": "Login",
                "url": "#",
                "style": "glyphicon glyphicon-log-in"

            }
        ]
        var items = data.map(this.generateItem);
        var user = userInfos.map(this.generateUserInfos);
        const sidebar = <SidebarContent setCurrentPage={this.setCurrentPage}/>;

        const contentHeader = (

            <div className="sidebar-toggle">
                <ul className="nav navbar-nav navbar-left">
                    {items}
                </ul>
                <ul className="nav navbar-nav navbar-right">

                    {user}
                </ul>

            </div>
        );

        const sidebarProps = {
            sidebar: sidebar,
            docked: this.state.docked,
            open: this.state.open,
            onSetOpen: this.onSetOpen,
        };
        return (
            <Sidebar {...sidebarProps}>
                <ContentHeader title={contentHeader}>
                    <div className="App">
                        <div className="App-header"></div>
                        <div className="container">
                            <div className="columns">
                                <div className="col-md-9 centered">
                                        <Grid>
                                            {
                                                React.cloneElement(this.currentPage(), {
                                                    setCurrentPage: this.setCurrentPage,
                                                    currentPage: this.state.currentPage,
                                                    ...this.state.currentPageProps,
                                                })
                                            }
                                        </Grid>

                                </div>
                            </div>
                            <footer>
                                <div className="copyRight">
                                    <span className="glyphicon glyphicon-copyright-mark"></span>
                                    <span> {new Date().getFullYear()}</span>
                                    -Next Africa Inc.  all rights reserved
                                </div>
                            </footer>
                        </div>
                    </div>
                </ContentHeader>
            </Sidebar>
        );
    }
 };
//App.propTypes = {
    //children: React.propTypes.node,
//}

export default Relay.createContainer(App,{
    initialVariables: {
        code: "ca"
    },
    fragments: {
        country : () => Relay.QL`
            fragment on Country{
                id
                properties{
                    code
                    name
                }
                
            }
        `
    }
});
