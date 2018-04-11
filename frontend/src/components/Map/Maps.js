/*global google*/
import React, { Component } from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import axios from 'axios';
import Notifications, {notify} from 'react-notify-toast';

class Maps extends Component {
    constructor(props) {
        super(props);

        this.state = {
            polygonArray: []
        }
    }

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
                onClick={(point) => this.checkPointInPolygon(point)}
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
                            clickable: true,
                            editable: true,
                            zIndex: 1,
                        },
                    }}
                    onPolygonComplete={(polygon) => this.createGeofence(polygon)}/>
            </GoogleMap>
        );
        return (
            <div>
                <Notifications />
                <MapWithADrawingManager />
            </div>
        );
    }

    // Creates a geofence from the drawn polygon
    createGeofence(polygon) {
        var polygonPointArray = this.createPolygonObject(polygon);
        polygon.addListener('click', (point) => {
            var data = JSON.stringify({
                latitude: Number.parseFloat(point.latLng.lat()),
                longitude: Number.parseFloat(point.latLng.lng())
            });

            // Get request to API to check point in polygon
            var url = 'http://159.203.178.86:8000/checkGeofence';
            axios.post(url, data, {
                headers: { 'Content-Type': 'application/json', }
            }).then(response => {
                if(response.data === 'false') {
                    console.log(response);
                    notify.show('Point not in Polygon!', "error", 3000);
                } else {
                    console.log(response);
                    var polygonSplit = response.data.split(",");
                    var stringNotify = "Point in Polygons (";
                    for(var i = 0; i < polygonSplit.length - 1; i++) {
                        if(i == 0)
                            stringNotify += (Number.parseInt(polygonSplit[i]) + 1);
                        else 
                            stringNotify += " & " + (Number.parseInt(polygonSplit[i]) + 1);
                    }
                    stringNotify += ")";
                    notify.show(stringNotify, "success", 3000);
                }
            })
        });

        var data = JSON.stringify({
            id: Number.parseInt(6, 10),
            name: 'polygon',
            points: polygonPointArray
        });

        // Post request to API to create geofence
        var url = 'http://159.203.178.86:8000/createGeofence';
        axios.post(url, data, {
            headers: { 'Content-Type': 'application/json', }
        }).then(response => {
            if(response.data === 'false') {
                console.log(response);
                notify.show('Point not in Polygon!', "error", 3000);
            }
        });
    }

    // Check point in polygon
    checkPointInPolygon(point) {
        var data = JSON.stringify({
            latitude: Number.parseFloat(point.latLng.lat()),
            longitude: Number.parseFloat(point.latLng.lng())
        });

        // Get request to API to check point in polygon
        var url = 'http://159.203.178.86:8000/checkGeofence';
        axios.post(url, data, {
            headers: { 'Content-Type': 'application/json', }
        }).then(response => {
            console.log(response.data);
            if(response.data === 'false') {
                notify.show('Point not in Polygon!', "error", 3000);
            }
        })
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
                console.log(polyString);
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
        this.state.polygonArray.forEach((polygon) => {
            var polygonToDraw = new google.maps.Polygon({
                paths: polygon,
                strokeColor: '#FF0000',
                strokeOpacity: 0.8,
                strokeWeight: 2,
                fillColor: '#FF0000',
                fillOpacity: 0.35
            });
            polygonToDraw.setMap(this.MapWithADrawingManager);
        });
    }

    // Create polygon object
    createPolygonObject(polygon) {
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
        return polygonPointArray;
    }
}

export default Maps;