import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { geolocated } from 'react-geolocated';

class App extends Component {
  render() {
    return (
      !this.props.isGeolocationAvailable
        ? <div>Your browser does not support Geolocation</div>
        : !this.props.isGeolocationEnabled
          ? <div>Geolocation is not enabled</div>
          : this.props.coords
            ? <table>
              <tbody>
                <tr><td>latitude</td><td>{this.props.coords.latitude}</td></tr>
                <tr><td>longitude</td><td>{this.props.coords.longitude}</td></tr>
                <tr><td>altitude</td><td>{this.props.coords.altitude}</td></tr>
              </tbody>
            </table>
            : <div>Getting the location data&hellip; </div>
    );
  }

  //Method for testing communication with the rudimentary API
  fetchCoordinateData() {
    //TODO: Add URL for digital ocean droplet here
    var url = '';
    axios.get(url)
      .then(response => {
        console.log(response);
      })
      .catch(error => {
        console.log(error);
      })
  }
}

export default geolocated({
  positionOptions: {
    enableHighAccuracy: false,
  },
  userDecisionTimeout: 5000,
})(App);
