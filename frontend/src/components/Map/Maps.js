/*global google*/
import React, { Component } from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import axios from 'axios';
import Polygon from 'react-google-maps';

class Maps extends Component {
    constructor(props) {
        super(props);

        this.state = {
            polygonArray: []
        }
    }

    render() {
        const { polygonArray } = this.state;
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
                defaultCenter={new google.maps.LatLng(-34.397, 150.644)}>
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
                    onPolygonComplete={(polygon) => this.savePolygonPoints(polygon)}/>
            </GoogleMap>
        );
        return (
            <div>
                <MapWithADrawingManager/>
                <RaisedButton id="polygon_button" label="Get Polygons" primary={true} onClick={() => this.getPolygonPoints()} />
                <RaisedButton id="draw_polygon_button" label="Draw Polygons" primary={true} onClick={() => this.drawPolygons()} />
            </div>
        );
    }

    // Save the coordinates of the polygon drawn on the map
    savePolygonPoints(polygon) {
        var locations = (polygon.getPath().getArray());
        var polygonPointArray = [];
        for (var i = 0; i < locations.length; i++) {
            var polygonObjectTemp = {};
            polygonObjectTemp.Latitude = Number.parseFloat(locations[i].lat());
            polygonObjectTemp.Longitude = Number.parseFloat(locations[i].lng());
            polygonPointArray.push(polygonObjectTemp);
        }

        // Final point is not correct (need same beginning and end point)
        var polygonObject = {};
        polygonObject.Latitude = Number.parseFloat(locations[0].lat());
        polygonObject.Longitude = Number.parseFloat(locations[0].lng());
        polygonPointArray.push(polygonObject);

        console.log(polygonPointArray);
        var data = JSON.stringify({
            id: Number.parseInt(6, 10),
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
                // Split response by delimeter for parsing
                var polyString = response.data["PolygonSummary"];
                var polyStringArray = polyString.split(",");
                var tempPolygonArray = [];
                var tempCoordinateObject = {};

                // Loop through each coordinate to pull out polygons
                polyStringArray.forEach((coordinate) => {
                    var tempCoordinate = coordinate.split(" ");
                    var tempPolygonCoordinateArray = [];
                    for (var i = 0; i < tempCoordinate.length; i++) {
                        tempPolygonCoordinateArray.push(new google.maps.LatLng(
                            parseFloat(tempCoordinate[0]),
                            parseFloat(tempCoordinate[1])
                        ));
                    }

                    // If polygon is complete reset the temp array
                    if (true) {
                        tempPolygonArray.push(tempCoordinateObject);
                        this.state.polygonArray.push(tempPolygonCoordinateArray);
                        tempPolygonArray = [];
                    } else {
                        tempPolygonArray.push(tempPolygonCoordinateArray);
                    }
                    console.log(this.state.polygonArray);
                });
            });
    }

    // Draw the polygons from the database on the map
    drawPolygons() {
        // this.state.polygonArray.forEach( (polygon) => {
        //     var polygonToDraw = new google.maps.Polygon({
        //         paths: polygon,
        //         strokeColor: '#FF0000',
        //         strokeOpacity: 0.8,
        //         strokeWeight: 2,
        //         fillColor: '#FF0000',
        //         fillOpacity: 0.35
        //     });
        //     polygonToDraw.setMap(this.MapWithADrawingManager);
        // });

        var triangleCoords = [
          {lat: -34.16097083888272, lng: 149.91890234375},
          {lat: -34.28814582106447, lng: 149.0454892578125},
          {lat: -34.650442819653406, lng: 149.0674619140625},
          {lat: -34.16097083888272, lng: 149.91890234375}
        ];

        this.state.polygonArray.push(triangleCoords);
    }

    // Prevents the componenet from reloading/updating on every event
    shouldComponentUpdate(nextProps, nextState) {
        return false;
    }
}

export default Maps;