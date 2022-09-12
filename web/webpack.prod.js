const path = require("path");
const { merge } = require("webpack-merge");
const common = require("./webpack.config");

module.exports = merge(common, {
	mode: "production",
	optimization: {
		usedExports: true,
		minimize: true
	},
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
});