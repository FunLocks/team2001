import React from 'react';
import GoogleMap from 'google-map-react';


const hakodate = {
    lat: 41.76205157771236,
    lng: 140.70437368224884 
};

const myCircle = null;
class Map extends React.Component {
    constructor(props) {
        super(props)
        this.state ={
            deleteFlag:false,
            radius:10000,
        }
    }
    static defaultProps = {
      center: {
        lat: 38.6777701,
        lng: 136.9065806
      },
    };
  
    
    componentDidMount = () => {
        this.timerID = setInterval(this.addRadius, 5000);

        console.log("Mount.");
    };

    componentDidUpdate = (prevProps, prevState) => {
        if (prevState.radius !== this.state.radius) {
            //再描画
            console.log("changed.");
            this.apiLoaded(prevState.mapRef,prevState.mapsRef)
        }
    };
    
    addRadius = () => {
        console.log(this.state.radius);
        var r = this.state.radius * 2;
        if(r > 500000){
            this.setState({
            deleteFlag:true,
        });
        }
        this.setState({
            radius:r,
        });
    }

    apiLoaded = (map,maps) => {
        this.setState({mapRef:map})
        this.setState({mapsRef:maps})
            if(this.state.deleteFlag)return;
            new maps.Circle({
                strokeColor:'red',
                strokeOpacity:0.8,
                strokeWeight:7,
                fillColor:'#FF0000',
                fillOpacity:0,
                map,
                center: hakodate,
                radius: this.state.radius,
            })
    }

    render() {
      return (
        <div style={{position: 'relative',height: '100%', width: '100%' }}>
          <GoogleMap
            bootstrapURLKeys={{ key: process.env.REACT_APP_GOOGLE_API_KEY}}
            defaultCenter={this.props.center}
            defaultZoom={6}
            onGoogleApiLoaded={({map, maps}) =>this.apiLoaded(map,maps)}
            />
        </div>
      );
    }
  }
  
  export default Map;