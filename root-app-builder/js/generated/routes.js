Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = [
  {isIndex: true, type: 'content-route', component: require('../components/Home.react').default, routes: []},
  {isIndex: false, path: '*', type: 'content-route', component: require('../components/NotFound.react').default, routes: []}
];