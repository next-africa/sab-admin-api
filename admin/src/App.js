import React  from 'react';
import '../node_modules/spectre.css/dist/spectre.min.css';
import './App.css';
import UniversityForm from './containers/UniversityForm';
import NavBarItem from './components/NavBarItem'
import ContentHeader from './containers/ContentHeader';
import SidebarContent from './containers/SidebarContent';
import  Sidebar  from 'react-sidebar';
const styles = {
    contentHeaderMenuLink: {
        textDecoration: 'none',
        color: 'white',
        padding: 8,
    },
    content: {
        padding: '16px',
    },
};
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
            dragToggleDistance: 30,};
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

            <span>
                <ul className="nav navbar-nav navbar-left">
                    {items}
                </ul>
                <ul className="nav navbar-nav navbar-right">

                    {user}
                </ul>
                <span>
                    {this.state.docked &&
                    <a onClick={this.toggleDocked} href="#" style={styles.contentHeaderMenuLink}></a>}
                    <span> S-A-B</span>
                </span>
            </span>
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
                        <div className="App-header">
                        </div>

                        <div className="container">

                            <div className="columns">
                                <div className="col-md-9 centered">
                                    <div style={styles.content}>
                                        <h3>Admin view to add a new university to the system</h3>
                                        <UniversityForm/>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </ContentHeader>
            </Sidebar>
        );
    }
 });

export default App;
