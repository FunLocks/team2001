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
 
class Map extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        myCircles : [],
        myRadiuss: [],
        myDeleteFrags : []
      };
    }
    static defaultProps = {
      center: {
        lat: 38.6777701,
        lng: 136.9065806
      },
    };
  
    componentDidUpdate(prevProps, prevState, snapshot){
      console.log("change")
      if(prevState.myRadiuss != this.state.myRadiuss){
        for(let i = 0; i < this.state.myRadiuss.length;i++){
          let newState = this.state.myCircles;
          newState[i].setRadius(this.state.myRadiuss[i]);
          this.setState({
            myRadiuss : newState,
          })
          
          console.log("set");
          if(this.state.myDeleteFrags[i]){
            this.state.myCircles[i].setVisible(true);
          }
        }
      }
    }
    componentDidMount = () => {
        this.timerID = setInterval(this.Update, 100);
        console.log("Mount.");
    };

    Update = () => {
        for(let i = 0; i < this.state.myRadiuss.length;i++){
          var r = this.state.myRadiuss[i] * 1.5;
          if(r > 500000){
            let newState = this.state.myDeleteFrags;
            newState[i] = true;
            this.setState({
              myDeleteFrags : newState,
            })
          }
          if(!this.state.myDeleteFrags[i]){
            let newState = this.state.myRadiuss;
            newState[i] = r;
            this.setState({
              myRadiuss : newState,
            })
          }
        }
        console.log(this.state.myCircles[0]);
        console.log(this.state.myRadiuss[0]);
        console.log(this.state.myDeleteFrags[0]);
    }

    apiLoaded = (map,maps,lat,lng) => {
        this.setState({mapRef:map})
        this.setState({mapsRef:maps})
        const lotate = {
          lat : lat,
          lmg : lng,
        }
        this.state.myDeleteFrags.push(false);
        this.state.myRadiuss.push(1000);
        this.state.myCircles.push(
          new maps.Circle({
            strokeColor:'red',
            strokeOpacity:0.8,
            strokeWeight:7,
            fillColor:'#FF0000',
            fillOpacity:0,
            map,
            center: lotate,
            radius: this.state.myRadiuss[this.state.myRadiuss.length-1],
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