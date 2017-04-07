/**
 * Created by pdiouf on 2017-04-03.
 */
import React from 'react';
import { HashRouter as Router, Route,  browserHistory } from 'react-router-dom';
import App from '../App';
import IndexPage from '../containers/UniversityForm';
import UniversityPage from '../containers/SingleUniversity';
import NotFoundPage from '../components/NotFoundPage';

class AppRoutes extends React.Component {
    render() {
        return (
            <Router  history={browserHistory} onUpdate={() => window.scrollTo(0, 0)}>
                <Route path="/" component={App}>
                    <indexRoute component={IndexPage}/>
                    <Route path="university/id" component={UniversityPage}/>
                    <Route path="*" component={NotFoundPage}/>
                </Route>
            </Router>
        );
    }
}
export default AppRoutes;
