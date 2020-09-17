let mix = require('laravel-mix');

mix.setPublicPath('public/')
    .react('resources/js/app.js', 'js')
    .webpackConfig({
        output: { chunkFilename: 'js/[name].js?id=[chunkhash]' },
        resolve: {
            alias: {
                '@': path.resolve('resources/js')
            }
        }
    })
    .version()
    .sourceMaps();