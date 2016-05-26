//noinspection JSUnresolvedVariable
import React, {Component} from 'react';

class Home extends Component {

  constructor(props) {
    super(props);
    this.state = {home: true}
  }

  render() {
    return (
      <div className="home-container">
        <h2>Default home page</h2>
      </div>
    );
  }
}

export default Home;
