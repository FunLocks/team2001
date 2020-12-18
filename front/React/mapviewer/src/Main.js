import React from 'react';
import Map from './Map';

const containerStyle = {
  position: 'relative',  
  width: '100%',
  height: '100%'
}

class Main extends React.Component {
  render() {
    return (
      <div className="main" id="main">
        <Map/>
      </div>
    );
  }
}

export default Main;