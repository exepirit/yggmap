const path = require("path");
const { merge } = require("webpack-merge");
const common = require("./webpack.config");

module.exports = merge(common, {
	mode: "production",
	optimization: {
		usedExports: true,
		minimize: true
	}
});