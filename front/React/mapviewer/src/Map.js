import React from 'react';
import GoogleMap from 'google-map-react';


const hakodate = {
    lat: 41.76205157771236,
    lng: 140.70437368224884 
};
const hakodate2 = {
  lat: 42.76205157771236,
  lng: 134.70437368224884 
};
 
var myCircle = [];
class Map extends React.Component {
    constructor(props) {
        super(props)
        this.state ={
            deleteFlag:[],
            radius:[],
            loaded: false,
        }
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
        this.timerID = setInterval(this.addRadius, 100);
        console.log("Mount.");
    };
    componentWillUnmount(){
      clearInterval(this.timerID);
    }

    componentDidUpdate = (prevProps, prevState) => {
            //再描画
            if(myCircle.length > 0 ){
              for(var i = 0; i < myCircle.length;i++){
                myCircle[0].setRadius(this.state.radius[0])
                if(this.state.deleteFlag[0]){
                  console.log("deleted");
                  myCircle[0].setVisible(false)
                  myCircle.shift()
                  //flag
                  var newstate = this.state.deleteFlag;
                  newstate.shift();
                  this.setState({
                    deleteFlag:newstate,
                  });
                  //
                  //radius
                  var newState = this.state.radius;
                  newState.shift();
                  this.setState({
                    radius:newState,
                  });
                  //
                }
              }
            }
    };
    
    addRadius = () => {
      if(this.state.loaded){
        for(var i = 0;i<myCircle.length;i++){
          var r = this.state.radius[i] * 1.5;
          if(r > 500000){
            var newstate = this.state.deleteFlag;
            newstate[i] = true;
            this.setState({
              deleteFlag:newstate,
            });
          }
          //Add radius
          var newState = this.state.radius;
          this.state.radius[i] = r;
          this.setState({
            radius:newState,
          });
        }
      }
    }

    apiLoaded = (map,maps,object) => {
        this.setState({mapRef:map})
        this.setState({mapsRef:maps})
        var newstate = this.state.deleteFlag;
        newstate.push(false)
        this.setState({
          deleteFlag:newstate,
        })
        var newState = this.state.radius;
        newState.push(1000)
        this.setState({
          radius:newState,
        })
        myCircle.push(
          new maps.Circle({
            strokeColor:'red',
            strokeOpacity:0.8,
            strokeWeight:7,
            fillColor:'#FF0000',
            fillOpacity:0,
            map,
            center: object,
            radius: this.state.radius[this.state.radius.length -1],
          })
        )
        this.setState({
          loaded:true,
        })
    }

    render() {
      return (
        <div style={{position: 'relative',height: '100%', width: '100%' }}>
          <GoogleMap
            bootstrapURLKeys={{ key: process.env.REACT_APP_GOOGLE_API_KEY}}
            defaultCenter={this.props.center}
            defaultZoom={6}
            onGoogleApiLoaded={({map, maps}) =>this.apiLoaded(map,maps,hakodate2)}
            />
        </div>
      );
    }
  }
  
  export default Map;