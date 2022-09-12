const path = require('path');
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require('mini-css-extract-plugin')

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
		}),
		new MiniCssExtractPlugin()
  ],
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
			{
        test: /\.css$/,
        use: [
					{
						loader: MiniCssExtractPlugin.loader
					},
					'css-loader',
				]
      }
		]
	}
}