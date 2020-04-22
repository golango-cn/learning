const conf = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  plugins: ['@babel/plugin-syntax-dynamic-import']
}
// 如果是发布模式则启用的插件
if (process.env.NODE_ENV === 'production') {
  conf.plugins.push('transform-remove-console')
}
module.exports = conf
