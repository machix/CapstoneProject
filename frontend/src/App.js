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
      deleteId: '',
      deleteLatitude: '',
      deleteLongitude: '',
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

  // Handles data change on the delete id text field
  handleDeleteId = (e) => {
    this.setState({ deleteId: e.target.value});
  }

  // Handles data change on the delete latitude field
  handleDeleteLatitude = (e) => {
    this.setState({ deleteLatitude: e.target.value});
  }

  // Handles data change on the delete longitude field
  handleDeleteLongitude = (e) => {
    this.setState({ deleteLongitude: e.target.value});
  }

  render() {
    return (
      <MuiThemeProvider>
        <Grid fluid>
          <Row>
            <Col xs={6} md={3}>
              <RaisedButton id="load_button" label="Load database" primary={true} onClick={() => this.fetchDatabaseInfo()} />
            </Col>
            <Col xs={6} md={3}>
              <Map />
            </Col>
          </Row>
          <br /><br /><br />
          <Row>
            <Col>
              <div id="data_table" className="Table">
                <ReactTable
                  data={this.state.data}
                  columns={this.state.columns}/>
              </div>
            </Col>
            <Col xs={6} md={3}>
              <div>
                <RaisedButton id="insert_button" label="Insert New Point" primary={true} onClick={() => this.insertNewPosition()} />
                <br />
                <TextField id="id_field" hintText="Id" onChange={this.handlePostId} />
                <TextField id="latitude_field" hintText="Latitude" onChange={this.handlePostLatitude} />
                <TextField id="longitude_field" hintText="Longitude" onChange={this.handlePostLongitude} />
              </div>
            </Col>
            <Col xs={6} md={3}>
              <div>
                <RaisedButton id="delete_button" label="Delete Point" primary={true} onClick={() => this.deletePosition()} />
                <br />
                <TextField id="id_delete_field" hintText="Id" onChange={this.handleDeleteId} />
                <TextField id="latitude_delete_field" hintText="Latitude" onChange={this.handleDeleteLatitude} />
                <TextField id="longitude_delete_field" hintText="Longitude" onChange={this.handleDeleteLongitude} />
              </div>
            </Col>
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
  insertNewPosition() {
    var url = 'http://159.203.178.86:8000/postPosition&Id=' + this.state.insertId +
      '&Latitude=' + this.state.insertLatitude + '&Longitude=' + this.state.insertLongitude;
    axios.put(url)
      .then(response => {
        console.log(response);
      });
  }

  // Delete position in the user database
  deletePosition() {
    var url = 'http://159.203.178.86:8000/deletePosition&Id=' + this.state.deleteId +
      '&Latitude=' + this.state.deleteLatitude + '&Longitude=' + this.state.deleteLongitude;
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
