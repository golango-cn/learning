module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/essential',
    '@vue/standard'
  ],
  parserOptions: {
    parser: 'babel-eslint'
  },
  rules: {
    indent: ['off', 2],
    'space-before-function-paren': 0,
    'space-before-keywords': 0,
    'no-useless-return': 0,
    'padded-blocks': 0,
    'eol-last': 0,
    'no-trailing-spaces': 0,
    'object-curly-spacing': 0,
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off'
  }
}
