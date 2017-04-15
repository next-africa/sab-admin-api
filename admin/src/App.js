import React  from 'react';
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
var universitiesList = [];

const App = React.createClass({
    getInitialState() {
        return {
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
            currentPageId : 0
        };

    },
    componentDidMount(){
        fetch("/api/countries/ca/universities", {
            headers:{
                'content-type': 'application/json'
            },
            method: "GET",
            mode:"same-origin",
            credentials: "same-origin"
        })
            .then((res) => res.json())
            .then(function(data){
                console.log(data.data);
                universitiesList : data.data
            })


    },
    setCurrentPage(event, { page, props , id}) {
        if (event) event.preventDefault();
        this.setState({ currentPage: page, currentPageProps: props , currentPageId: id});
    },
    currentPage() {
        return {

            universities:<Universities universitiesList={universitiesList}/>,
            universityForm: <UniversityForm universitiesList={universitiesList}/>,
            newUniversity: <NewUniversity universitiesList={universitiesList}/>,
            viewUniversity: <ViewUniversity universitiesList={universitiesList} currentPageId={this.state.currentPageId}/>,
        }[this.state.currentPage];
    },

    componentWillMount() {
        const mql = window.matchMedia(`(min-width: 800px)`);
        mql.addListener(this.mediaQueryChanged);
        this.setState({mql: mql, docked: mql.matches});
    },

    componentWillUnmount() {
        this.state.mql.removeListener(this.mediaQueryChanged);
    },

    onSetOpen(open) {
        this.setState({open: open});
    },
    onSetDocket(docked) {
        this.setState({docked: !docked});
    },

    mediaQueryChanged() {
        this.setState({docked: this.state.mql.matches});
    },

    toggleOpen(ev) {
        this.setState({open: !this.state.open});

        if (ev) {
            ev.preventDefault();
        }
    },
    toggleDocked(ev) {
        this.setState({docked: !this.state.docked});

        if (ev) {
            ev.preventDefault();
        }
    },
    generateItem : function(item){
        return <NavBarItem controlFunc={this.toggleDocked}style={item.style} key={item.key} text={item.text} url={item.url}/>
    },
    generateUserInfos: function(infos){
        return <NavBarItem style={infos.style} key={infos.key} text={infos.text} url={infos.url}/>

    },

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
        const sidebar = <SidebarContent />;

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
                                <p>
                                    <strong>S</strong>tudy <strong>A</strong>broad <strong>B</strong>ot.
                                </p>
                            </footer>
                        </div>
                    </div>
                </ContentHeader>
            </Sidebar>
        );
    }
 });
//App.propTypes = {
    //children: React.propTypes.node,
//}

export default App;
