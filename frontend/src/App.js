import React, { Component } from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import ReactTable from 'react-table';
import { Grid, Row, Col } from 'react-flexbox-grid';
import axios from 'axios';
import Map from "./components/Map";

class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      getRequestResponse: '',
      getRequestPositionResponse: '',
      data: [],
      columns: [
        {
          Header: "Id",
          accessor: "id"
        },
        {
          Header: "Latitude",
          accessor: "latitude"
        },
        {
          Header: "Longitude",
          accessor: "longitude"
        }
      ]
    }
  }



  render() {
    return (
      <MuiThemeProvider>
        <Grid fluid>
          <Row>
            <Col xs={6} md={3}>
              <RaisedButton label="Load database" primary={true} onClick={() => this.fetchDatabaseInfo()} />
              <br />
              <br />
              <div>
                <RaisedButton label="Insert New Point" primary={true} />
                <br />
                <TextField placeholder="Id" />
                <TextField placeholder="Latitude" />
                <TextField placeholder="Longitude" />
              </div>
              <br />
            </Col>
            <Col xs={6} md={3}>
              <Map />
            </Col>
          </Row>
          <br /><br /><br /><br /><br />
          <Row>
              <div className="Table">
                <ReactTable
                  data={this.state.data}
                  columns={this.state.columns}
                  />
              </div>
          </Row>
        </Grid>
      </MuiThemeProvider>
    );
  }

  // Function for fetching all of the info in the database
  fetchDatabaseInfo() {
    var url = 'http://159.203.178.86:8000/position';
    axios.get(url)
      .then(response => {
        var res = response.data;
        this.setState({ data: [] });
        var tempTableArray = [];
        for (var key in res) {
          if (res.hasOwnProperty(key)) {
            var value = res[key];
            var tempObject = {};
            tempObject.id = value[0].Id;
            tempObject.latitude = value[0].Latitude;
            tempObject.longitude = value[0].Longitude;
            tempTableArray.push(tempObject);
          }
        }
        this.setState({data: tempTableArray});
        console.log(this.state.data);
      })
    this.showTable();
  }

  showTable() {
    //Show the table once the data is loaded
  }

  // Function for fetching info from the database give ID
  fetchDatabaseId(id) {
    var url = 'http://159.203.178.86:8000/getPosition';
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        var tempTableArray = [];
        for (var i = 0; i < res.length; i++) {
          var tempObject;
          tempObject.Id = res[i].Id;
          tempObject.Latitude = res[i].Latitude;
          tempObject.Longitude = res[i].Longitude;
          this.state.tempTableArray.push(tempObject);
        }
        this.setState({data: tempTableArray});
        console.log(this.state.data);
      })
  }

  // Function for testing communication with the rudimentary API
  fetchBasicEndpoint(endpoint) {
    var url = 'http://159.203.178.86:8000' + endpoint;
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        if (endpoint === "/position") {
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
