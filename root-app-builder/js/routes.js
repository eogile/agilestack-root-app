//noinspection JSUnresolvedVariable
import React from 'react';
import { Route, IndexRoute } from 'react-router';

import { App, Main } from './generated/components';
import pluginsRoutes from './generated/routes';

let counter = 0;

const renderRoute = (route) => {
  if (route.routes && route.routes.length > 0) {
    return renderRouteWithChildren(route)
  }

  //noinspection JSUnresolvedVariable
  if (route.isIndex) {
    return renderIndexRoute(route);
  }

  /*
   * If it has no children it must have a path.
   */
  console.log('renderRouteWithNoChildren');
  return (
    <Route key={route.href} path={route.href} component={route.component} />
  )
};

const renderRouteWithChildren = (route) => {
  console.log('renderRouteWithChildren');
  const props = {
    key: route.href || (counter++),
    component: route.component
  };
  if (route.href) {
    props.path= route.href;
  }
  return (
    <Route {...props}>
      {
        route.routes.map((subRoute) =>
          renderRoute(subRoute)
        )
      }
    </Route>
  )
};

/**
 * Returns the JSX statement for the given index route.
 */
const renderIndexRoute = (route) => {
  console.log('renderIndexRoute');
  return <IndexRoute key={counter++} component={route.component} />
};

export default () => {
  let contentRoutes = pluginsRoutes.filter((route) => "content-route" === route.type);
  let fullScreenRoutes = pluginsRoutes.filter((route) => "full-screen-route" === route.type);

  return (
    <Route path="/" component={App}>
      {
        fullScreenRoutes.map((route, index) =>
          renderRoute(route)
        )
      }
      <Route component={Main}>
        {
          contentRoutes.map((route, index) =>
            renderRoute(route)
          )
        }
      </Route>
    </Route>
  )
}
