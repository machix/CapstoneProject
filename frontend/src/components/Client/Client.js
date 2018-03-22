import React, { Component } from 'react';

class Client extends Component {
    constructor() {
        this.state = {
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
    
    render() {
        return(
            <Row>
            <Col xs={6} md={3}>
              <div id="data_table" className="Table">
                <ReactTable
                  data={this.state.data}
                  columns={this.state.columns}/>
              </div>
            </Col>
            <Col xs={6} md={3}>
              <RaisedButton id="load_button" label="Load Client Database" primary={true} onClick={() => this.fetchClientDatabaseInfo()} />
            </Col>
            <Col xs={6} md={3}>
              <div>
                <RaisedButton id="insert_button" label="Insert New Client" primary={true} onClick={() => this.insertNewClient()} />
                <br />
                <TextField id="id_field" hintText="Id" onChange={this.handlePostId} />
                <TextField id="latitude_field" hintText="Latitude" onChange={this.handlePostLatitude} />
                <TextField id="longitude_field" hintText="Longitude" onChange={this.handlePostLongitude} />
              </div>
            </Col>
            <Col xs={6} md={3}>
              <div>
                <RaisedButton id="delete_button" label="Delete Client" primary={true} onClick={() => this.deleteClient()} />
                <br />
                <TextField id="id_delete_field" hintText="Id" onChange={this.handleDeleteId} />
                <TextField id="first_name_delete_field" hintText="First Name" onChange={this.handleDeleteLatitude} />
                <TextField id="last_name_delete_field" hintText="Last Name" onChange={this.handleDeleteLongitude} />
              </div>
            </Col>
          </Row>
        );
    }

    // Inserts a New Client into the database
    insertNewClient() {

    }

    // Fetches the info for the client database
    fetchClientDatabaseInfo() {

    }

    // Deletes a client from the database
    deleteClient() {

    }
}

export default Client;