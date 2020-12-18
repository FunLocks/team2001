import React from 'react';

class Main extends React.Component {

  componentWillMount() {
    return fetch('http://192.168.20.155:8080/get')
      .then((response) => response.json())
      .then((responseJson) => (console.log("get"))
      )
      .catch((error) => {
        console.error(error);
      });
  }

  render() {
    return (
      <div className='main'>
        <h1>this is main.</h1>
      </div>
    );
  }
}

export default Main;