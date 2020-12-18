import React from 'react';

class AhchooCounter extends React.Component{
  constructor(props){
    super(props);
    this.state = {
      dayOf:0,
      weekOf:0,
      monthOf:0
    }
  }
  
  componentWillMount() {
    //return fetch('http://192.168.20.155:8080/get')
    return fetch('http://153.120.166.49:8080/get')
      .then((response) => response.json())
      .then((responseJson) => (
          this.setState({
            dayOf:responseJson.latitude,
            weekOf:responseJson.longitude,
          })
        )
      )
      .catch((error) => {
        console.error(error);
      });
  }

  render(){
    return(
       <div className="menu">
         <h1 className="menuTitle">今日のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.dayOf} 回</h2>
         <h1 className="menuTitle">今週のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.weekOf} 回</h2>
         <h1 className="menuTitle">今月のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.monthOf} 回</h2>
      </div>
    );
  }
}


class Header extends React.Component {
  render() {
    return (
      <div className='sidebar'>
          <p className="title">AHCHOO!</p>
          <p className="sub">くしゃみ/咳のマップビューワー</p>
          <AhchooCounter></AhchooCounter>
      </div>
    );
  }
}

export default Header;