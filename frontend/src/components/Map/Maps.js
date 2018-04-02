import React, { Component } from 'react';
import axios from 'axios';

class Maps extends Component {
    render() {
        const google = window.google;
        const { compose, withProps } = require("recompose");
        const {
            withScriptjs,
            withGoogleMap,
            GoogleMap,
        } = require("react-google-maps");
        const { DrawingManager } = require("react-google-maps/lib/components/drawing/DrawingManager");

        const MapWithADrawingManager = compose(
            withProps({
                googleMapURL: "https://maps.googleapis.com/maps/api/js?key=AIzaSyCru0FXiEWV7vQnvQvTNHyQisAi96J2rlI&v=3.exp&libraries=geometry,drawing,places",
                loadingElement: <div style={{ height: `100%` }} />,
                containerElement: <div style={{ height: `500px`, width: '1400px' }} />,
                mapElement: <div style={{ height: `100%` }} />,
            }),
            withScriptjs,
            withGoogleMap
        )(props =>
            <GoogleMap
                defaultZoom={8}
                defaultCenter={new google.maps.LatLng(-34.397, 150.644)}
            >
                <DrawingManager
                    defaultDrawingMode={google.maps.drawing.OverlayType.POLYGON}
                    defaultOptions={{
                        drawingControl: true,
                        drawingControlOptions: {
                            position: google.maps.ControlPosition.TOP_CENTER,
                            drawingModes: [
                                google.maps.drawing.OverlayType.CIRCLE,
                                google.maps.drawing.OverlayType.POLYGON,
                                google.maps.drawing.OverlayType.POLYLINE,
                                google.maps.drawing.OverlayType.RECTANGLE,
                            ],
                        },
                        circleOptions: {
                            fillColor: `#ffff00`,
                            fillOpacity: 1,
                            strokeWeight: 5,
                            clickable: false,
                            editable: true,
                            zIndex: 1,
                        },
                    }}
                    onPolygonComplete={(polygon) => this.savePolygonPoints(polygon)}
                />
            </GoogleMap>
        );
        return (
            <div>
                <MapWithADrawingManager />
            </div>
        );
    }

    // Save the coordinates of the polygon drawn on the map
    savePolygonPoints(polygon) {
        var locations = (polygon.getPath().getArray());
        var polygonPointArray = [];
        for (var i = 0; i < locations.length; i++) {
            var polygonObject = {};
            polygonObject.Latitude = Number.parseFloat(locations[i].lat());
            polygonObject.Longitude = Number.parseFloat(locations[i].lng());
            polygonPointArray.push(polygonObject);
        }

        // Final point is not correct (need same beginning and end point)
        var polygonObject = {};
        polygonObject.Latitude = Number.parseFloat(locations[0].lat());
        polygonObject.Longitude = Number.parseFloat(locations[0].lng());
        polygonPointArray.push(polygonObject);

        console.log(polygonPointArray);
        var data = JSON.stringify({
            id: Number.parseInt(6),
            name: 'polygon2',
            points: polygonPointArray
        });

        console.log(data);

        var url = 'http://159.203.178.86:8000/savePolygon';
        axios.post(url, data, {
            headers: { 'Content-Type': 'application/json', }
        }).then(response => {
            console.log(response);
        });
    }

    // Download polygons contained within certain area on the map
    getPolygonPoints() {
        var url = "http://159.203.178.86:8000/getPolygons";
        axios.get(url)
            .then(response => {
                console.log(response);
            });
    }

    // Prevents the componenet from reloading/updating on every event
    shouldComponentUpdate(nextProps, nextState) {
        return false;
    }
}

export default Maps;