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
          <RaisedButton label="Get Request" primary={true} onClick={() => this.fetchBasicEndpoint("")} />
          <br />
          <TextField value={this.state.getRequestResponse} />
          <br />
          <br />
          <RaisedButton label="Get Position Endpoint" primary={true} onClick={() => this.fetchBasicEndpoint("/position")} />
          <br />
          <TextField value={this.state.getRequestPositionResponse} />
          <br />
          <RaisedButton label="Search database" primary={true}/>
          <br />
          <br />
          <RaisedButton label="Get data from Id" primary={true}/>
          <br />
          <Table>
            <TableHeader>
              <TableRow>
                <TableHeaderColumn>ID</TableHeaderColumn>
                <TableHeaderColumn>Latitude</TableHeaderColumn>
                <TableHeaderColumn>Longitude</TableHeaderColumn>
              </TableRow>
              <TableBody>
                {this.state.tableData.map((row, Id) => (
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

  // Function for fetching all of the info in the database
  fetchDatabaseInfo() {
    var url = 'http://159.203.178.86:8000/getPositions';
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        this.state.setState(this.tableData = []);
        for(var i = 0; i < res.length; i++) {
          var tempObject;
          tempObject.Id = res[i].Id;
          tempObject.Latitude = res[i].Latitude;
          tempObject.Longitude = res[i].Longitude;
          this.state.tableData.push(tempObject)
        }
      })
  }

  // Function for fetching info from the database give ID
  fetchDatabaseId(id) {
    var url = 'http://159.203.178.86:8000/getPosition?=ID';
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        this.state.setState(this.tableData = []);
        for(var i = 0; i < res.length; i++) {
          var tempObject;
          tempObject.Id = res[i].Id;
          tempObject.Latitude = res[i].Latitude;
          tempObject.Longitude = res[i].Longitude;
          this.state.tableData.push(tempObject);
        }
      })
  }

  // Function for testing communication with the rudimentary API
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
