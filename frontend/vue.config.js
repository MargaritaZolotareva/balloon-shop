const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig(
    {
      transpileDependencies: true,
      css: {
        loaderOptions: {
          sass: {
            additionalData: `@import "@/assets/styles/variables.scss";`
          }
        }
      },
      publicPath: '/',
      chainWebpack: config => {
        config.resolve.alias.set('vue$', 'vue/dist/vue.esm-bundler.js');
      }
    });
