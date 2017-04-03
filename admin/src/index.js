import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import {HashRouter as Router, Route, browserHistory } from 'react-router-dom'

ReactDOM.render((
    <Router  history={browserHistory}>
        <Route    path="/" component={App}>
        </Route>
    </Router>

),document.getElementById('root'));

