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
      insertId: '',
      insertLongitude: '',
      insertLatitude: '',
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

  // Handles data change on the insert id text field
  handlePostId = (e) => {
    this.setState({ insertId: e.target.value });
  }

  // Handles data change on the insert latitude text field
  handlePostLatitude = (e) => {
    this.setState({ insertLatitude: e.target.value });
  }

  // Handles data change on the insert longitude text field
  handlePostLongitude = (e) => {
    this.setState({ insertLongitude: e.target.value });
  }

  render() {
    return (
      <MuiThemeProvider>
        <Grid fluid>
          <Row>
            <Col xs={6} md={3}>
              <RaisedButton id="load_button" label="Load database" primary={true} onClick={() => this.fetchDatabaseInfo()} />
              <br />
              <br />
              <div>
                <RaisedButton id="insert_button" label="Insert New Point" primary={true} onClick={() => this.insertNewPosition()} />
                <br />
                <TextField id="id_field" hintText="Id" onChange={this.handlePostId} />
                <TextField id="latitude_field" hintText="Latitude" onChange={this.handlePostLatitude} />
                <TextField id="longitude_field" hintText="Longitude" onChange={this.handlePostLongitude} />
              </div>
              <br />
            </Col>
            <Col xs={6} md={3}>
              <Map />
            </Col>
          </Row>
          <br /><br /><br /><br /><br />
          <Row>
            <div id="data_table" className="Table">
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
        console.log(response);
        var res = response.data;
        this.setState({ data: [] });
        var tempTableArray = [];
        for (var key in res) {
          if (res.hasOwnProperty(key)) {
            var value = res[key];
            for (var i = 0; i < value.length; i++) {
              var tempObject = {};
              tempObject.id = value[i].Id;
              tempObject.latitude = value[i].Latitude;
              tempObject.longitude = value[i].Longitude;
              tempTableArray.push(tempObject);
            }
          }
        }
        this.setState({ data: tempTableArray });
        console.log(this.state.data);
      })
  }

  // Insert new position into the user database
  insertNewPosition(position) {
    var url = '159.203.178.86:8000/postPosition&id=' + this.state.insertId +
      '&latitude=' + this.state.insertLatitude + '&longitude=' + this.state.insertLongitude;
    axios.put(url)
      .then(response => {
        console.log(response);
      });
  }

  // Delete position in the user database
  deletePosition(id) {
    var url = '159.203.178.86:8000/deletePosition';
    axios.delete(url)
      .then(response => {
        console.log(response);
      })
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
        this.setState({ data: tempTableArray });
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
