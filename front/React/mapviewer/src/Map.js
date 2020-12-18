import React from 'react';
import GoogleMap from 'google-map-react';


const hakodate = {
    lat: 41.76205157771236,
    lng: 140.70437368224884 
};

const hakodate2 = {
  lat: 41.76205157771236,
  lng: 136.70437368224884 
};
 
var myCircles = [];
var myRadiuss = [];
var myDeleteFrags = [];
class Map extends React.Component {
    static defaultProps = {
      center: {
        lat: 38.6777701,
        lng: 136.9065806
      },
    };
  
    
    componentDidMount = () => {
        this.timerID = setInterval(this.Update, 100);
        console.log("Mount.");
    };

    Update = () => {
        for(let i = 0; i < myRadiuss.length;i++){
          var r = myRadiuss[i] * 1.5;
          if(r > 500000){
            myDeleteFrags[i] = true;
          }
          if(!myDeleteFrags[i]) myRadiuss[i] = r;
        }
        console.log(myCircles[0]);
        console.log(myRadiuss[0]);
        console.log(myDeleteFrags[0]);

        for(let i = 0; i < myRadiuss.length;i++){
          myCircles[i].setRadius(myRadiuss[i]);
          console.log("set");
          if(myDeleteFrags[i]){
            myCircles[i].setVisible(true);
          }
        }
    }

    apiLoaded = (map,maps,lat,lng) => {
        this.setState({mapRef:map})
        this.setState({mapsRef:maps})
        const lotate = {
          lat : lat,
          lmg : lng,
        }
        myDeleteFrags.push(false);
        myRadiuss.push(1000);
        myCircles.push(
          new maps.Circle({
            strokeColor:'red',
            strokeOpacity:0.8,
            strokeWeight:7,
            fillColor:'#FF0000',
            fillOpacity:0,
            map,
            center: lotate,
            radius: myRadiuss[myRadiuss.length-1],
          })
        );
    }

    render() {
      return (
        <div style={{position: 'relative',height: '100%', width: '100%' }}>
          <GoogleMap
            bootstrapURLKeys={{ key: process.env.REACT_APP_GOOGLE_API_KEY}}
            defaultCenter={this.props.center}
            defaultZoom={6}
            onGoogleApiLoaded={({map, maps}) =>this.apiLoaded(map,maps,hakodate.lat,hakodate.lng)}
            />
        </div>
      );
    }
  }
  
  export default Map;