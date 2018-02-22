import React, { Component } from 'react';
import { Map, InfoWindow, Marker, GoogleApiWrapper } from 'google-maps-react';

class Maps extends Component {
    render() {
        return (
            <Map style={{ height:300 }}google={this.props.google} zoom={14}>
                <Marker onClick={this.onMarkerClick}
                    name={'Current location'} />
                <InfoWindow onClose={this.onInfoWindowClose}>
                    <div>
                        <h1>Map</h1>
                    </div>
                </InfoWindow>
            </Map>
        );
    }
}

export default GoogleApiWrapper({
    apiKey: ("AIzaSyDqJIIYPWM5hlaGmbENa9SaSSKWPtMOAKs")
})(Maps)