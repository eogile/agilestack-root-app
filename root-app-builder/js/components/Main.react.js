//noinspection JSUnresolvedVariable
import React, {Component} from 'react';

var Home = require('./Home.react');

class Main extends Component {
  render() {
    return (
      <div className="main-container">
        <h1>Default main container tavu</h1>
        <div>
          <Home/>
        </div>
        <div>
          {this.props.children}
        </div>
      </div>
    );
  }
}

export default Main;
