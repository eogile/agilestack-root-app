//noinspection JSUnresolvedVariable
import React, {Component} from 'react';
import Menu from './Menu.react';

export default ({children}) => (
  <div className="main-container">
    <h1>Default main container</h1>
    <div className="menu-container">
      <Menu/>
    </div>
    <div className="content-container">
      {children}
    </div>
  </div>
);
