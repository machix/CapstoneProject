import React, { Component } from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import ReactTable from 'react-table';
import { Grid, Row, Col } from 'react-flexbox-grid';
import axios from 'axios';
import Map from "./components/Map";
import Client from "./components/Client";
import User from "./components/User";

class App extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <MuiThemeProvider>
        <Grid fluid>
          <Row>
            <Col>
              <Map />
            </Col>
          </Row>
          <br /> <br /> <br />
          <Row>
            <User />
          </Row>
          <br /> <br />
          <Row>
            <Client />
          </Row>
        </Grid>
      </MuiThemeProvider>
    );
  }
}

export default App;
