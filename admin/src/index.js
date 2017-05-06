

import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import Relay from 'react-relay'
 /*ReactDOM.render(
     <App/>,
     document.getElementById('root')
 );*/



class AppRoute extends Relay.Route {
 static queries = {
 country: () => Relay.QL`
 query { 
    country(code: $code)
 }
 `,
 };
 static paramDefinitions = {
 // By setting `required` to true, `ProfileRoute` will throw if a `userID`
 // is not supplied when instantiated.
 code: {required: true},
 };
 static routeName = 'AppRoute';
 }

 let appRoute = new AppRoute({code:"ca"});
 ReactDOM.render(
 <Relay.RootContainer
 Component={App}
 route={appRoute}
 />,
 document.getElementById('root')
 );





