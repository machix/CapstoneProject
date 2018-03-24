import React, { Component } from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import ReactTable from 'react-table';
import { Grid, Row, Col } from 'react-flexbox-grid';
import axios from 'axios';

// Client component for demonstrating working service
class Client extends Component {
  constructor(props) {
    super(props);

    this.state = {
      clientId: '',
      clientFirstName: '',
      clientLastName: '',
      clientDeleteId: '',
      data: [],
      columns: [
        {
          Header: "Id",
          accessor: "id"
        },
        {
          Header: "First Name",
          accessor: "first_name"
        },
        {
          Header: "Last Name",
          accessor: "last_name"
        }
      ]
    }
  }

  // Handle change in the client id text field
  handlePostClientId = (e) => {
    this.setState({ clientId: e.target.value });
  }

  // Handle change in the client first name text field
  handlePostFirstName = (e) => {
    this.setState({ clientFirstName: e.target.value });
  }

  // Handle change in the client last name text field
  handlePostLastName = (e) => {
    this.setState({ clientLastName: e.target.value });
  }

  // Handle change in the client delete id text field
  handleDeleteId = (e) => {
    this.setState({ clientDeleteId: e.target.value });
  }

  render() {
    return (
      <Row>
        <Col xs={6} md={3}>
          <div id="data_table" className="Table">
            <ReactTable
              data={this.state.data}
              columns={this.state.columns} />
          </div>
        </Col>
        <Col xs={6} md={3}>
          <RaisedButton id="load_button" label="Load Client Database" primary={true} onClick={() => this.fetchClientDatabaseInfo()} />
        </Col>
        <Col xs={6} md={3}>
          <div>
            <RaisedButton id="insert_button" label="Insert New Client" primary={true} onClick={() => this.insertNewClient()} />
            <br />
            <TextField id="id_field" hintText="Id" onChange={this.handlePostClientId} />
            <TextField id="first_name_field" hintText="Latitude" onChange={this.handlePostFirstName} />
            <TextField id="last_name_field" hintText="Longitude" onChange={this.handlePostLastName} />
          </div>
        </Col>
        <Col xs={6} md={3}>
          <div>
            <RaisedButton id="delete_button" label="Delete Client" primary={true} onClick={() => this.deleteClient()} />
            <br />
            <TextField id="id_delete_field" hintText="Id" onChange={this.handleDeleteId} />
          </div>
        </Col>
      </Row>
    );
  }

  // Inserts a New Client into the database
  insertNewClient() {
    var url = 'http://159.203.178.86:8000/postClient';
    let data = JSON.stringify({
      Id: Number.parseInt(this.state.insertId),
      FirstName: this.state.clientFirstName,
      LastName: this.state.clientLastName
    });
    axios.post(url, data, {
      headers: { 'Content-Type': 'application/json', }
    }).then(response => {
      console.log(response);
    })
  }

  // Fetches the info for the client database
  fetchClientDatabaseInfo() {
    var url = 'http://159.203.178.86:8000/getClients';
    axios.get(url)
      .then(response => {
        console.log(response.data);
        var res = response.data;
        var tempTableArray = [];
        for (var i = 0; i < res.length; i++) {
          var tempObject;
          tempObject.Id = res[i].Id;
          tempObject.FirstName = res[i].FirstName;
          tempObject.LastName = res[i].LastName;
          this.state.tempTableArray.push(tempObject);
        }
        this.setState({ data: tempTableArray });
        console.log(this.state.data);
      });
  }

  // Deletes a client from the database
  deleteClient() {
    var url = 'http://159.203.178.86:8000/deleteClients';
    let data = JSON.stringify({
      Id: Number.parseInt(this.state.clientDeleteId),
    });
    var url = 'http://159.203.178.86:8000/deletePosition';
    axios.delete(url, data, {
      headers: { 'Content-Type': 'application/json', }
    }).then(response => {
      console.log(response);
    });
  }
}

export default Client;