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