/**
 * Presentational component displaying a menu.
 */
import React from 'react';
import {Link} from 'react-router';

/**
 * Defines the Menu component as a stateless functional component.
 *
 * @param menuEntries {Array} The menu entries.
 * @param loadFailure {Boolean} A flag indicating the menu load failed.
 */
const Menu = ({menuEntries, loadFailure}) => {
  if (loadFailure) {
    return (<p style={{color: 'red'}}>There was an error while loading the menu</p>);
  }
  return (
    <ul>
      {
        menuEntries.map((menuEntry) =>
          <li key={menuEntry.route}>
            <Link to={menuEntry.route}>{menuEntry.name}</Link>
          </li>
        )
      }
    </ul>
  );
};

export default Menu;

