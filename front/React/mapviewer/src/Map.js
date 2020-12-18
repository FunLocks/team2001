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
 
var myCircle = null;
var myCircle2 = null;
class Map extends React.Component {
    constructor(props) {
        super(props)
        this.state ={
            deleteFlag:false,
            radius:10000,
            deleteFlag2:false,
            radius2:10000,
        }
    }
    static defaultProps = {
      center: {
        lat: 38.6777701,
        lng: 136.9065806
      },
    };
  
    
    componentDidMount = () => {
        this.timerID = setInterval(this.addRadius, 100);

        console.log("Mount.");
    };

    componentDidUpdate = (prevProps, prevState) => {
        if (prevState.radius !== this.state.radius) {
            //再描画
            console.log("changed.");
            myCircle.setRadius(this.state.radius)
            myCircle2.setRadius(this.state.radius)
            if(this.state.deleteFlag){
              myCircle.setVisible(false)
            }
            
            if(this.state.deleteFlag2){
              myCircle2.setVisible(false)
            }
        }
    };
    
    addRadius = () => {
        console.log(this.state.radius);
        var r = this.state.radius * 1.5;
        var r2 = this.state.radius2 * 1.5;
        if(r > 500000){
            this.setState({
              deleteFlag:true,
            });
        }
        if(r2 > 500000){
          this.setState({
            deleteFlag2:true,
          });
        }
        
        this.setState({
            radius:r,
        });
        this.setState({
          radius2:r2,
      });
    }

    apiLoaded = (map,maps) => {
        this.setState({mapRef:map})
        this.setState({mapsRef:maps})
        myCircle = new maps.Circle({
          strokeColor:'red',
          strokeOpacity:0.8,
          strokeWeight:7,
          fillColor:'#FF0000',
          fillOpacity:0,
          map,
          center: hakodate,
          radius: this.state.radius,
        })
        myCircle2 = new maps.Circle({
          strokeColor:'red',
          strokeOpacity:0.8,
          strokeWeight:7,
          fillColor:'#FF0000',
          fillOpacity:0,
          map,
          center: hakodate2,
          radius: this.state.radius2,
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