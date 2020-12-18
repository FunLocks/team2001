import React from 'react';
import {Map, InfoWindow,Circle, Marker, GoogleApiWrapper} from 'google-maps-react';

const containerStyle = {
  position: 'relative',  
  width: '100%',
  height: '100%'
}
const initialCenter={
  lat: 38.6777701,
  lng: 136.9065806
}
const hakodate = {
  lat: 41.76205157771236,
  lng: 140.70437368224884 
};

export class MapContainer extends React.Component {
  
  constructor(props){
    super(props);
    this.state = {
      radius: 1000,
      hoge: true,
    };
  }
  componentDidMount = () => {
    setInterval(this.addRadius, 100);
    console.log("Mount.");
  };

  addRadius = () => {
    console.log(this.state.radius);
    var r = this.state.radius * 2;
    if(r > 500000){
      this.setState({
        hoge:false,
      });
    }
    this.setState({
      radius:r,
    });
  }
  
  render() {
    return (
      <div className="container">
        <Map google={this.props.google} 
          zoom={6} 
          containerStyle={containerStyle}
          initialCenter={initialCenter}>
            {this.state.hoge &&
              <Circle
                radius={this.state.radius}
                center={hakodate}
                strokeColor='red'
                strokeOpacity={0.8}
                strokeWeight={7}
                fillColor='#FF0000'
                fillOpacity={0}
              />
            }
        </Map>
      </div>
    );
  }
}
 
export default GoogleApiWrapper({
  apiKey: (process.env.REACT_APP_GOOGLE_API_KEY)
})(MapContainer)