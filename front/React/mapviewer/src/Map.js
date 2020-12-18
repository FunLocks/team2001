import React from 'react';
import {Map, InfoWindow, Marker, GoogleApiWrapper} from 'google-maps-react';

const containerStyle = {
  position: 'relative',  
  width: '100%',
  height: '100%'
}
export class MapContainer extends React.Component {
  render() {
    return (
      <div className="container">
        <Map google={this.props.google} zoom={14} containerStyle={containerStyle}>
          <Marker onClick={this.onMarkerClick}
           name={'Current location'} />
        </Map>
      </div>
    );
  }
}
 
export default GoogleApiWrapper({
  apiKey: (process.env.REACT_APP_GOOGLE_API_KEY)
})(MapContainer)