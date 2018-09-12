const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const extractSass = new ExtractTextPlugin({
    filename: 'ddtip.css',
});

module.exports = {
  mode: 'production',
  entry: [
      './src/js/index.ts',
      './src/scss/main.scss'
  ],
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/
      },
      {
        test: /\.scss$/,
        include: [
            path.resolve(__dirname, 'src/scss')
        ],
        use: extractSass.extract({
            use: [{
                loader: 'css-loader',
                options: { minimize: true }
            }, {
                loader: 'sass-loader'
            }],
        })
      }
    ]
  },
  plugins: [
    extractSass
  ],
  resolve: {
    extensions: [ '.tsx', '.ts', '.js' ]
  },
  output: {
    filename: 'ddtip.js',
    path: path.resolve(__dirname, 'dist')
  }
};