import React, { Component } from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import axios from 'axios';

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
                // Split response by delimeter for parsing
                var polyString = response.data["PolygonSummary"];
                var polyStringArray = polyString.split(",");
                var tempPolygonArray = [];

                // Loop through each coordinate to pull out polygons
                polyStringArray.forEach( (coordinate) => {
                    var tempCoordinate = coordinate.split(" ");
                    var tempCoordinateObject = {};
                    tempCoordinateObject.lat = Number.parseFloat(tempCoordinate[0]);
                    tempCoordinateObject.lng = Number.parseFloat(tempCoordinate[1]);

                    // If polygon is complete reset the temp array
                    if(tempPolygonArray.find(x => x.lat  === tempCoordinateObject.lat)) {
                        tempPolygonArray.push(tempCoordinateObject);
                        this.state.polygonArray.push(tempPolygonArray);
                        tempPolygonArray = [];
                    } else {
                        tempPolygonArray.push(tempCoordinateObject);
                    }
                    console.log(this.state.polygonArray);
                });
            });
    }

    // Draw the polygons from the database in the database
    drawPolygons() {
        this.state.polygonArray.forEach( (polygon) => {
            var polygonToDraw = this.google.maps.Polygon({
                paths: polygon,
                strokeColor: '#FF0000',
                strokeOpacity: 0.8,
                strokeWeight: 2,
                fillColor: '#FF0000',
                fillOpacity: 0.35
            });
            polygonToDraw.setMap(this.props.GoogleMap);
        });
    }

    // Prevents the componenet from reloading/updating on every event
    shouldComponentUpdate(nextProps, nextState) {
        return false;
    }
}

export default Maps;