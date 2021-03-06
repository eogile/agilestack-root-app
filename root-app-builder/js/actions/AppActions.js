/*
 * Actions change things in your application
 * Since this boilerplate uses a uni-directional data flow, specifically redux,
 * we have these actions which are the only way your application interacts with
 * your application state. This guarantees that your state is up to date and nobody
 * messes it up weirdly somewhere.
 *
 * To add a new Action:
 * 1) Import your constant
 * 2) Add a function like this:
 *    export function yourAction(var) {
 *        return { type: YOUR_ACTION_CONSTANT, var: var }
 *    }
 * 3) (optional) Add an async function like this:
 *    export function asyncYourAction(var) {
 *        return (dispatch) => {
 *             // Do async stuff here
 *             return dispatch(yourAction(var));
 *        };
 *    }
 *
 *    If you add an async function, remove the export from the function
 *    created in the second step
 */

// Disable the no-use-before-define eslint rule for this file
// It makes more sense to have the asnyc actions before the non-async ones
/* eslint-disable no-use-before-define */

import {
  CHANGE_OWNER_NAME, CHANGE_PROJECT_NAME,
  LOAD_MAIN_MENU
} from '../constants/AppConstants';

import {ApiConfig} from '../config/config';

export function asyncChangeProjectName(name) {
  return (dispatch) => {
    // You can do async stuff here!
    // API fetching, Animations,...
    // For more information as to how and why you would do this, check https://github.com/gaearon/redux-thunk
    return dispatch(changeProjectName(name));
  };
}

export function asyncChangeOwnerName(name) {
  return (dispatch) => {
    // You can do async stuff here!
    // API fetching, Animations,...
    // For more information as to how and why you would do this, check https://github.com/gaearon/redux-thunk
    return dispatch(changeOwnerName(name));
  };
}

export function changeProjectName(name) {
  return {type: CHANGE_PROJECT_NAME, name};
}

export function changeOwnerName(name) {
  return {type: CHANGE_OWNER_NAME, name};
}

/**
 * Loads the menu entries from the server and then dispatches an action
 * to update the store.
 *
 * @returns {function()} The action to dispatch.
 */
export function asyncLoadMainMenu() {
  return (dispatch) => {
    // noinspection JSValidateTypes
    return fetch(ApiConfig.apiURL + '/menu-entries')
      .then(response => response.json())
      .then(entries => dispatch({
        type: LOAD_MAIN_MENU,
        entries: entries,
        error: false
      }))
      .catch(error => {
        console.error('Error while loading the menu entries', error);
        dispatch({
          type: LOAD_MAIN_MENU,
          entries: [],
          error: true
        });
      });
  };
}
