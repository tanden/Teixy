# support

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn run serve
```

### Compiles and minifies for production
```
yarn run build
```

### Run your tests
```
yarn run test
```

### Lints and fixes files
```
yarn run lint
```

### Run your end-to-end tests
```
yarn run test:e2e
```

### Run your unit tests
```
yarn run test:unit
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


## dependencies
### vue
言わずと知れたVue.jsの本体
2019/01の最新は2.5.21

### vue-router
Vue.js公式のルータ
2019/01の最新は3.0.1

### vuex
Vue.js公式の「状態管理パターン + ライブラリ」で、アプリケーションが複雑になってくると活躍しだす
2019/01の最新は3.0.1

## devDependencies
### babel
ES6 -> ES5に変換するトランスパイラ

### cypress
chrome限定のe2eテストツール<br>
seleniumを使うより幸せになれるらしい

### eslint
lintツール<br>
こまめにアップデートしないと、その都度コードのフォーマットを大きく直さないといけなくなるので注意

### jest
Facebook謹製のunit testツール<br>
jsのunitテストといえばmocha + chaiだが、今後はjestが主流になる予感

### node-sass
Node.jsからLibSass（C++で書かれたsassのコンパイラ）にバインドしてくれるので、コンパイルが早くなる、らしい...<br>
webpackでsassをコンパイルしてバンドルするときに、sass-loaderともに必要なライブラリ

### sass-loader
sassをcssにコンパイルするためのライブラリ