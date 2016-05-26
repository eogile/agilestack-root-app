/**
 * Presentational component displaying a menu.
 */
import React from 'react';
import {Link} from 'react-router';
import {connect} from 'react-redux';

/**
 * Defines the Menu component as a stateless functional component.
 *
 * @param entries {Array} The menu entries.
 * @param error {Boolean} A flag indicating the menu load failed.
 */
const Menu = ({entries, error}) => {
  console.log('Menu', entries, error);
  return (
    <div className="menu-container">
      <h2>Default menu</h2>
      {error ? (
        <p style={{color: 'red'}}>There was an error while loading the menu</p>
      ) : (
        <MenuEntries entries={entries}/>
      )}
    </div>
  );
};

const MenuEntries = ({entries}) => {
  console.log('MenuEntries', entries);
  return(
  <ul>
    {
      entries && entries.map((entry) =>
        <li key={entry.route}>
          <Link to={entry.route}>{entry.name}</Link>
          {entry.entries && <MenuEntries entries={entry.entries} />}
        </li>
      )
    }
  </ul>
)
};

export default connect(({MainMenu}) => ({...MainMenu}))(Menu);
