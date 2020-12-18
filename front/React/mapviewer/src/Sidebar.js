import React from 'react';

class AhchooCounter extends React.Component{
  constructor(props){
    super(props);
    this.state = {
      hourOf:0,
      dayOf:0,
      weekOf:0,
      monthOf:0
    }
  }
  
  fetchAPI = () =>{
    //return fetch('http://192.168.20.155:8080/get')//test server
    fetch('http://153.120.166.49:8080/ahchoo/one-hour')
      .then((response) => response.json())
      .then((responseJson) => (
          this.setState({
            hourOf:responseJson.length,
          })
        )
      )
      .catch((error) => {
        console.error(error);
      });
    fetch('http://153.120.166.49:8080/ahchoo/one-day')
      .then((response) => response.json())
      .then((responseJson) => (
          this.setState({
            dayOf:responseJson.length,
          })
        )
      )
      .catch((error) => {
        console.error(error);
      });
    fetch('http://153.120.166.49:8080/ahchoo/seven-days')
      .then((response) => response.json())
      .then((responseJson) => (
          this.setState({
            weekOf:responseJson.length,
          })
        )
      )
      .catch((error) => {
        console.error(error);
      });
    fetch('http://153.120.166.49:8080/ahchoo/thiry-days')
      .then((response) => response.json())
      .then((responseJson) => (
          this.setState({
            monthOf:responseJson.length,
          })
        )
      )
      .catch((error) => {
        console.error(error);
      });
  }

  componentWillMount() {
    this.fetchAPI()
    this.timerID = setInterval(this.fetchAPI, 10000);
  }
  
  componentWillUnmount(){
    clearInterval(this.timerID);
  }

  render(){
    return(
       <div className="menu">
         <h1 className="menuTitle">1時間以内のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.hourOf} 回</h2>
         <h1 className="menuTitle">1日以内のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.dayOf} 回</h2>
         <h1 className="menuTitle">1週間以内のAHCHOO!</h1>
         <h2 className="menuValue">{this.state.weekOf} 回</h2>
         <h1 className="menuTitle">1ヶ月以内のAHCHOO!</h1>
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