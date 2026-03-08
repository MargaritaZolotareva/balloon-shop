const {defineConfig} = require('@vue/cli-service')
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
        },
        configureWebpack: {
            optimization: {
                splitChunks: {
                    chunks: 'all',
                    cacheGroups: {
                        vue: {
                            test: /[\\/]node_modules[\\/](vue|vue-router)[\\/]/,
                            name: 'vue-core',
                            chunks: 'all',
                            priority: 20
                        },
                        vendors: {
                            test: /[\\/]node_modules[\\/]/,
                            name: 'vendors',
                            chunks: 'all',
                            priority: 10
                        }
                    }
                }
            }
        },
        pages: {
            index: {
                entry: 'src/main.js',
                title: 'ПОДАРОКК | Воздушные и гелиевые шары'
            }
        }
    });
