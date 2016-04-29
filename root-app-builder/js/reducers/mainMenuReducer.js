import {LOAD_MAIN_MENU} from '../constants/AppConstants';

/**
 * Reducer in charge of storing the entries of the main menu
 * into the application state.
 *
 * @param menu {Object} The previous state.
 * @param action The action to perform.
 * @returns {Object} The new state.
 */
export default function(menu = {entries: []}, action) {
  /*
   * Ensuring that the state is not modified.
   */
  Object.freeze(menu);

  switch (action.type) {
    /*
     * The menu entries were loaded from the server.
     */
    case LOAD_MAIN_MENU:
      return Object.assign({}, menu, {entries: action.entries, error: action.error});
    default:
      return menu;
  }
}
