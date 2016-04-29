/**
 * Combine all reducers in this file and export the combined reducers.
 * If we were to do this in store.js, reducers wouldn't be hot reloadable.
 */

import { combineReducers } from 'redux';
import homeReducer from './homeReducer';
import MainMenu from './mainMenuReducer';

import reducers from '../generated/reducers';

// Replace line below once you have several reducers with
// import { combineReducers } from 'redux';
// const rootReducer = combineReducers({ homeReducer, yourReducer })
const rootReducer = combineReducers({
  ...reducers,
  homeReducer,
  MainMenu
});

export default rootReducer;
