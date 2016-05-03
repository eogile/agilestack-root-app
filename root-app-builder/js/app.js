/**
 *
 * app.js
 *
 * This is the entry file for the application, mostly just setup and boilerplate
 * code. Routes are configured at the end of this file!
 *
 */

// Load the Cache polyfill, the manifest.json file and the .htaccess file
import 'file?name=[name].[ext]!../manifest.json';
import 'file?name=[name].[ext]!../.htaccess';


// Import all the third party stuff
import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Router, Route, IndexRoute } from 'react-router';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import FontFaceObserver from 'fontfaceobserver';
import { createHistory, useBasename } from 'history';

/**
 * FontFaceObserver is a @font-face loader that can notify when
 * a web-front is loaded.
 *
 * In this case, the expected web front is "Open Sans" that is
 * loaded in index.html with the "link" tag inside the "body" tag.
 */
// Observer loading of Open Sans (to remove open sans, remove the <link> tag in the index.html file and this observer)
const openSansObserver = new FontFaceObserver('Open Sans', {});

// When Open Sans is loaded, add the js-open-sans-loaded class to the body
openSansObserver.check().then(() => {
  document.body.classList.add('js-open-sans-loaded');
}, () => {
  document.body.classList.remove('js-open-sans-loaded');
});

// Import the pages
import HomePage from './components/pages/HomePage.react';
import ReadmePage from './components/pages/ReadmePage.react';
import NotFoundPage from './components/pages/NotFound.react';
import App from './components/App.react';

// Import the CSS file, which HtmlWebpackPlugin transfers to the build folder
import '../css/main.css';

// Create the store with the redux-thunk middleware, which allows us
// to do asynchronous things in the actions
import rootReducer from './reducers/rootReducer';
const createStoreWithMiddleware = applyMiddleware(thunk)(createStore);
const store = createStoreWithMiddleware(rootReducer);

// Make reducers hot reloadable, see http://stackoverflow.com/questions/34243684/make-redux-reducers-and-other-non-components-hot-loadable
if (module.hot) {
  module.hot.accept('./reducers/rootReducer', () => {
    const nextRootReducer = require('./reducers/rootReducer').default;
    store.replaceReducer(nextRootReducer);
  });
}

const history = useBasename(createHistory)({
  basename: window.baseUrl
});

/**
 * FIXME Removes these imports when components will be provided by other plugins.
 */
import {Plugin1Feature1, Plugin2Feature1, Plugin2Feature2} from './components/pages/TemporaryMenuPages.react';

import pluginsRoutes from './generated/routes';

// Mostly boilerplate, except for the Routes. These are the pages you can go to,
// which are all wrapped in the App component, which contains the navigation etc
ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRoute component={HomePage}/>
        <Route key="readme" path="/readme" component={ReadmePage}/>

        <Route key="plugin1/feature1" path="/plugin1/feature1" component={Plugin1Feature1}/>
        <Route key="plugin2/feature1" path="/plugin2/feature1" component={Plugin2Feature1}/>
        <Route key="plugin2/feature2" path="/plugin2/feature2" component={Plugin2Feature2}/>

        {
          pluginsRoutes.map((route) =>
            <Route key={route.href} path={route.href} component={route.component} />
          )
        }

        <Route path="*" component={NotFoundPage}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('app')
);

/*
 * Loading the main menu
 */
import {asyncLoadMainMenu} from './actions/AppActions';
store.dispatch(asyncLoadMainMenu());