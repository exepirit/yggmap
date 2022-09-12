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
	plugins: [new webpack.HotModuleReplacementPlugin()],
	module: {
		rules: [
			{
				test: /\.jsx?$/,
				loader: "babel-loader",
				exclude: /node_modules/,

				options: {
					presets: [
            "@babel/preset-env",
					],
					plugins: [
						["@babel/plugin-transform-react-jsx", {
              "pragma": "h",
              "pragmaFrag": "Fragment",
            }]
					]
				}
			},
		]
	}
})