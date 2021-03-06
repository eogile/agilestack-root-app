var path = require('path');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var AppCachePlugin = require('appcache-webpack-plugin');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var failPlugin = require('webpack-fail-plugin');

module.exports = function(options) {
  var entry, jsLoaders, plugins, cssLoaders, resolversExtensions;

  // If production is true
  if (options.prod) {
    // Entry
    entry = [
      path.resolve(__dirname, 'js/app.js') // Start with js/app.js...
    ];
    cssLoaders = ExtractTextPlugin.extract('style-loader', 'css-loader!postcss-loader');
    // Plugins
    plugins = [// Plugins for Webpack
      new webpack.optimize.UglifyJsPlugin({ // Optimize the JavaScript...
        compress: {
          warnings: false // ...but do not show warnings in the console (there is a lot of them)
        },
        sourceMap:false
      }),

      /*
       * This plugin is used to add DOM elements in the "index.html" file to insert artifacts
       * generated by Webpack.
       */
      new HtmlWebpackPlugin({
        /*
         * No file should be generated. Instead, the "index.html" file must be modified.
         */
        template: 'index.html', // Move the index.html file...

        /*
         * The artifacts must be minified. The given options are for "html-minifier".
         *
         * The options details are available here :
         * https://github.com/kangax/html-minifier#options-quick-reference
         */
        minify: {
          removeComments: true,
          collapseWhitespace: true,
          /*
           * Some HTML attributes have default values. For example, the default value
           * of the "type" attribute of the "input" element is "text".
           *
           * This option removes an attribute when its value is the default one.
           */
          removeRedundantAttributes: true,
          useShortDoctype: true,
          /*
           * If an attribute value is blank, then the attribute is removed.
           */
          removeEmptyAttributes: true,
          /*
           * If there is an "type" attribute whose value is "text/css" in a "link" tag
           * or in a "style" tag, then the attribute is removed.
           */
          removeStyleLinkTypeAttributes: true,
          /*
           * For HTML tags that are not required to have a closing tag (such as "br"),
           * adds a closing slash ("<br>" => "<br/>")
           */
          keepClosingSlash: true,
          /*
           * Minifying JavaScript sources using UglifyJS. It's possible to pass
           * options.
           *
           * TODO Check why the UglifyJS plugin is configured outside
           */
          minifyJS: true,
          /*
           * Minifying CSS with clean-css
           */
          minifyCSS: true,
          /*
           * Minifying URLs when possible by replacing absolute URLs by relative ones.
           */
          minifyURLs: true
        },
        inject: true // inject all files that are generated by webpack, e.g. bundle.js, main.css with the correct HTML tags
      }),
      new ExtractTextPlugin("css/main.css"),
      new webpack.DefinePlugin({
        "process.env": {
          NODE_ENV: JSON.stringify("production")
        }
      })
    ];

    /*
     * Adding a new extension to load the production configuration.
     * It's crucial that ".agilestack.prod.js" is before ".js".
     */
    resolversExtensions = ["", ".agilestack.prod.js", ".webpack.js", ".web.js", ".js"];

  /*
   * Configuration for development environment.
   */
  } else {
    /*
     * The Webpack entry points.
     */
    entry = [
      /*
       * The webpack-server-server supports two ways to refresh the page : the iframe
       * mode and the inline mode.
       *
       * For inline mode, the "webpack-dev-server" client entry point must be manually
       * added (webpack-dev-server/client?http://localhost:3000). For explanation, see
       * http://webpack.github.io/docs/webpack-dev-server.html#inline-mode-with-node-js-api
       */
      "webpack-dev-server/client?http://localhost:3000",
      /*
       * The "webpack/hot/dev-server" entry point is added for hot module replacement.
       * See http://webpack.github.io/docs/webpack-dev-server.html#hot-module-replacement-with-node-js-api
       */
      "webpack/hot/only-dev-server",
      path.resolve(__dirname, 'js/app.js')
    ];

    /*
     * The CSS loader.
     *
     * The "style-loader" loads the CSS files that are referenced by some entry points.
     *
     * With the "css-loader", the CSS (and images ?) referenced inside a CSS (with @import for
     * example) will be resolved as CSS and will be treated as such by the next loader.
     *
     * The "postcss-loader" executes a treatment on loaded CSS. This treatment is specified
     * in the "postcss" function assigned to the returned object.
     */
    cssLoaders = 'style-loader!css-loader!postcss-loader';

    //noinspection JSUnresolvedFunction
    plugins = [
      /*
       * This plugin must be added to enable hot module replacement.
       * See http://webpack.github.io/docs/webpack-dev-server.html#hot-module-replacement-with-node-js-api
       */
      new webpack.HotModuleReplacementPlugin(),

      /*
       * This plugin is used to add DOM elements in the "index.html" file to insert artifacts
       * generated by Webpack.
       */
      new HtmlWebpackPlugin({
        /*
         * No file should be generated. Instead, the "index.html" file must be modified.
         */
        template: 'index.html',
        /*
         * The artifacts will be inserted in the HTML body, not the head.
         */
        inject: true
      })
    ];

    /*
     * Adding a new extension to load the development configuration.
     * It's crucial that ".agilestack.dev.js" is before ".js".
     */
    resolversExtensions = ["", ".agilestack.dev.js", ".webpack.js", ".web.js", ".js"];
  }

  plugins.push(new AppCachePlugin({ // AppCache should be in both prod and dev env
    exclude: ['.htaccess'] // No need to cache that. See https://support.hostgator.com/articles/403-forbidden-or-no-permission-to-access
  }));

  /*
   * This plugin makes Webpack returns an exit code equal to 1
   * when there is a compilation error.
   */
  plugins.push(failPlugin);

  return {
    entry: entry,
    output: { // Compile into build/js/bundle.js
      path: path.resolve(__dirname, 'build'),
      filename: "js/bundle.js"
    },

    resolve: {
      /*
       * The order of extensions matters.
       */
      extensions: resolversExtensions

    },
    module: {
      loaders: [{
          test: /\.js$/, // Transform all .js files required somewhere within an entry point...
          loader: 'babel', // ...with the specified loaders...
          exclude: path.join(__dirname, '/node_modules/') // ...except for the node_modules folder.
        }, {
          test:   /\.css$/, // Transform all .css files required somewhere within an entry point...
          loader: cssLoaders // ...with PostCSS
        }, {
          test: /\.jpe?g$|\.gif$|\.png$/i,
          loader: "url-loader?limit=10000"
        }
      ]
    },
    plugins: plugins,
    postcss: function() {
      return [
        require('postcss-import')({ // Import all the css files...
          /*
           * TODO ????
           */
          glob: true,
          /*
           * Function called after the "import" process.
           *
           * We add each imported file as a dependency to the loader result.
           *
           * So, when one these files change, Webpack (dev server) knows that the result
           * should change. Is this why hot reloading work ?
           */
          onImport: function (files) {
              files.forEach(this.addDependency); // ...and add dependecies from the main.css files to the other css files...
          }.bind(this) // ...so they get hot–reloaded when something changes...
        }),
        /*
         * Replaces the places where CSS variables are used by their values.
         */
        require('postcss-simple-vars')(),
        /*
         * For all the CSS selectors with a ":hover" pseudo-class, this plugin adds
         * the same CSS selector but the a ":focus" pseudo-class to the same CSS rule.
         */
        require('postcss-focus')(),
        /*
         * Parses the CSS files and adds vendors prefixes for CSS rules using
         * values from CanIUse.
         */
        require('autoprefixer')({
          browsers: ['last 2 versions', 'IE > 8'] // ...supporting the last 2 major browser versions and IE 8 and up...
        }),
        /*
         * Displays in the console warnings emitted by other plugins.
         */
        require('postcss-reporter')({
          /*
           * Once a message is logged by this plugin, it's removed from the result's messages.
           * This ensures that the message will be logged only once.
           */
          clearMessages: true
        })
      ];
    },
    target: "web", // Make web variables accessible to webpack, e.g. window
    stats: false, // Don't show stats in the console
    progress: true
  }
};
