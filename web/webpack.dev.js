const path = require("path");
const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.config");

module.exports = merge(common, {
	mode: "development",
	optimization: {
		usedExports: true
	},
	devtool: "inline-source-map",
	devServer: {
		hot: true,
		compress: false
	},
	plugins: [new webpack.HotModuleReplacementPlugin()]
})