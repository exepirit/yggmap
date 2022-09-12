const path = require('path');
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  entry: './src/app.js',
  output: {
    filename: "[name].bundle.js",
		path: path.resolve(__dirname, "./dist")
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"]
  },
  plugins: [
		new HtmlWebpackPlugin({
			title: "WebPack - Preact",
			template: "./public/index.html",
			inlineSource: ".(js|css)$",
			minify: {
				collapseWhitespace: true,
				removeComments: true
			}
		})
  ]
}