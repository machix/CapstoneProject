import React, { Component } from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import ReactTable from 'react-table';
import { Row, Col } from 'react-flexbox-grid';
import axios from 'axios';

// User component to demonstrate working service
class User extends Component {
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
        this.handlePostId = this.handlePostId.bind(this);
        this.handlePostLatitude = this.handlePostLatitude.bind(this);
        this.handlePostLongitude = this.handlePostLongitude.bind(this);
        this.handleDeleteId = this.handleDeleteId.bind(this);
        this.handleDeleteLatitude = this.handleDeleteLatitude.bind(this);
        this.handleDeleteLongitude = this.handleDeleteLongitude.bind(this);
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
        this.setState({ deleteId: e.target.value });
    }

    // Handles data change on the delete latitude field
    handleDeleteLatitude = (e) => {
        this.setState({ deleteLatitude: e.target.value });
    }

    // Handles data change on the delete longitude field
    handleDeleteLongitude = (e) => {
        this.setState({ deleteLongitude: e.target.value });
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
                    <RaisedButton id="load_button" label="Load database" primary={true} onClick={() => this.fetchDatabaseInfo()} />
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
            });
    }

    // Insert new position into the user database
    insertNewPosition() {
        let data = JSON.stringify({
            Id: Number.parseInt(this.state.insertId, 10),
            Latitude: Number.parseFloat(this.state.insertLatitude),
            Longitude: Number.parseFloat(this.state.insertLongitude)
        });
        var url = 'http://159.203.178.86:8000/postPosition';
        axios.post(url, data, {
            headers: { 'Content-Type': 'application/json', }
        }).then(response => {
            console.log(response);
        });
    }

    // Delete position in the user database
    deletePosition() {
        let data = JSON.stringify({
            Id: Number.parseInt(this.state.deleteId, 10),
            Latitude: Number.parseFloat(this.state.deleteLatitude),
            Longitude: Number.parseFloat(this.state.deleteLongitude)
        });
        var url = 'http://159.203.178.86:8000/deletePosition';
        axios.delete(url, data, {
            headers: { 'Content-Type': 'application/json', }
        }).then(response => {
            console.log(response);
        });
    }
}

export default User;