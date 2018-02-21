import React, { Component } from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import {
  Table,
  TableBody,
  TableHeader,
  TableHeaderColumn,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';
import axios from 'axios';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      getRequestResponse: '',
      getRequestPositionResponse: '',
      tableData: []
    }
  }

  render() {
    return (
      <MuiThemeProvider>
        <div>
          <RaisedButton label="Get Request" onClick={() => this.fetchBasicEndpoint("")} />
          <br />
          <TextField value={this.state.getRequestResponse} />
          <br />
          <RaisedButton label="Get Position Endpoint" onClick={() => this.fetchBasicEndpoint("/position")} />
          <br />
          <TextField value={this.state.getRequestPositionResponse} />
          <br />
          <RaisedButton label="Search database" onClick={() => this.fetchDatabaseInfo()} />
          <br />
          <RaisedButton label="Get data from Id" onClick={() => this.fetchDatabaseId()} />
          <br />
          <Table>
            <TableHeader>
              <TableRow>
                <TableHeaderColumn>ID</TableHeaderColumn>
                <TableHeaderColumn>Latitude</TableHeaderColumn>
                <TableHeaderColumn>Longitude</TableHeaderColumn>
              </TableRow>
              <TableBody>
                {tableData.map((row, Id) => (
                  <TableRow key={Id}>
                    <TableRowColumn>{Id}</TableRowColumn>
                    <TableRowColumn>{row.Latitude}</TableRowColumn>
                    <TableRowColumn>{row.Longitude}</TableRowColumn>
                  </TableRow>
                ))}
              </TableBody>
            </TableHeader>
          </Table>
        </div>
      </MuiThemeProvider>
    );
  }

  //Method for testing communication with the rudimentary API
  fetchBasicEndpoint(endpoint) {
    var url = 'http://159.203.178.86:8000' + endpoint;
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        if (endpoint == "/position") {
          var responseString = res['Text'];
          this.setState({ getRequestPositionResponse: responseString })
        } else {
          this.setState({ getRequestResponse: res });
        }
      })
      .catch(error => {
        console.log(error);
      })
  }
}

export default App;
