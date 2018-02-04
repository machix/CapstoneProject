import React, { Component } from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import axios from 'axios';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      getRequestResponse: '',
      getRequestPositionResponse: ''
    }
  }

  render() {
    return (
      <MuiThemeProvider>
        <div>
          <RaisedButton label="Get Request" onClick={() => this.fetchBasicEndpoint()} />
          <br />
          <TextField value={this.state.getRequestResponse} />
        </div>
        </MuiThemeProvider>
        );
  }

  //Method for testing communication with the rudimentary API
  fetchBasicEndpoint() {
      var url = 'http://159.203.178.86:8000';
      axios.get(url)
        .then(response => {
          console.log(response.data);
        var res = response.data;
          this.setState({getRequestResponse: res});
        })
        .catch(error => {
          console.log(error);
        })
  }
}

export default App;
