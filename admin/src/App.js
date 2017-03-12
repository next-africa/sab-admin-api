import React, { Component } from 'react';
import logo from './logo.svg';
import '../node_modules/spectre.css/dist/spectre.min.css';
import './App.css';
import UniversityForm from './containers/UniversityForm';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Study abroad </h2>
        </div>
          <div className="container">
              <div className="columns">
                  <div className="col-md-9 centered">
                      <h3>Admin view to add a new university to the system</h3>
                        <UniversityForm/>
                  </div>
              </div>
          </div>
      </div>
    );
  }
}

export default App;
