module.exports = {
  lintOnSave: true,
  devServer: {
    port: 8083,
    proxy: {
      '/api': {
        ws: true,
        target: 'http://localhost:8092',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
}
