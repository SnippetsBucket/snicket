import 'babel-polyfill';
import path from 'path';
import webpack from 'webpack';
import AssetsPlugin from 'assets-webpack-plugin';

const DEBUG = !process.argv.includes('--release');
const VERBOSE = process.argv.includes('--verbose');
const assetsPluginInstance = new AssetsPlugin();

export default {
  cache: DEBUG,

  debug: DEBUG,

  stats: {
    colors: true,
    reasons: DEBUG,
    hash: VERBOSE,
    version: VERBOSE,
    timings: true,
    chunks: VERBOSE,
    chunkModules: VERBOSE,
    cached: VERBOSE,
    cachedAssets: VERBOSE,
  },

  entry: {
    create: ['./assets/js/create.entry.js'],
    // login: ['./src/js/login.entry.js']
  },

  output: {
    publicPath: 'public/dist/',
    sourcePrefix: '  ',
    path: path.join(__dirname, 'public/dist'),
    filename: '[name].bundle.js',
  },

  target: 'web',

  devtool: DEBUG ? 'cheap-module-eval-source-map' : false,

  plugins: [
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.DefinePlugin({ 'process.env.NODE_ENV': `"${process.env.NODE_ENV || (DEBUG ? 'development' : 'production')}"` }),
    ...(DEBUG ? [] : [
      new webpack.optimize.DedupePlugin(),
      new webpack.optimize.UglifyJsPlugin({ compress: { screw_ie8: true, warnings: VERBOSE, drop_console: true } }),
      new webpack.optimize.AggressiveMergingPlugin(),
    ]),
    assetsPluginInstance
  ],

  resolve: {
    extensions: ['', '.js'],
  },

  module: {
    loaders: [
      { test: /\.js?$/, include: [path.resolve(__dirname, 'app/assets/js')], loader: 'babel' },
    ],
  },

};