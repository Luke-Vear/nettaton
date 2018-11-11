// Required for the sagas
require("@babel/polyfill");

const port = process.env.PORT || 3002;
 
// Html plugin puts index.html into the dist dir with a script tag pointing to bundle.js
const HtmlWebPackPlugin = require("html-webpack-plugin");

const htmlPlugin = new HtmlWebPackPlugin({
  template: "./public/index.html",
  filename: "./index.html"
});

// TL;DR - https://github.com/webpack-contrib/extract-text-webpack-plugin
const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
	entry: ["@babel/polyfill", "./src/index.js"],
	module: {
		rules: [
			{
				test: /\.(js|jsx)$/,
				exclude: /node_modules/,
				use: ['babel-loader']
			},
			{
				test: /\.scss$/,
				use: ExtractTextPlugin.extract({
					fallback: "style-loader",
					use: "css-loader!sass-loader",
				})
			}
		]
	},
	resolve: {
		extensions: ['*', '.js', '.jsx']
	},
	output: {
		path: __dirname + '/dist',
		publicPath: '/',
		filename: 'bundle.js'
	},
	devServer: {
		contentBase: './dist',
		port: port
	},
	plugins: [
		htmlPlugin,
		new ExtractTextPlugin('style.css')
	]
};
