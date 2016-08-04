'use strict';

var entryPoint = 'static/src/main.jsx';

var browserify = require('gulp-browserify'),
    cleanCSS = require('gulp-clean-css'),
    concat = require('gulp-concat'),
    gulp = require('gulp'),
    gutil = require('gulp-util'),
    less = require('gulp-less'),
    literalify = require('literalify'),
    merge = require('merge-stream'),
    reactify = require('reactify'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify');

gulp.task('scripts', function() {
    return gulp.src(entryPoint)
        .pipe(browserify({
            debug: false,
            extensions: ['.jsx', '.js', '.json'],
            transform: [reactify, literalify.configure({
                'react':          'window.React',
                'react-dom':      'window.ReactDOM',
                'jquery':         'window.jQuery',
                'swagger-client': 'window.SwaggerClient'
            })]
        }))
        .on('error', function(err) {
            gutil.log(err.message)
        })
        //.pipe(uglify())
        .pipe(rename('client.js'))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('vendor', function() {
    gulp.src([
        'bower_components/react/react.min.js',
        'bower_components/react/react-dom.min.js',
        'bower_components/jquery/dist/jquery.min.js',
        'bower_components/swagger-js/browser/swagger-client.min.js',
        'bower_components/jpillora/jquery.rest/dist/1/jquery.rest.min.js'
    ])
        .pipe(concat('vendor.js'))
        .pipe(uglify())
        .pipe(gulp.dest('static/dist'));
});

gulp.task('styles', function () {
    return merge(
        gulp.src('static/src/styles/*.less').pipe(less()),
        gulp.src('static/src/styles/*.css')
    )
        .pipe(concat('styles.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('watch', function () {
    gulp.watch(["static/src/*.jsx", "static/src/*.js"], ['scripts']);
    gulp.watch(["static/src/styles/*.less", "static/src/style/*.css"], ['styles']);
});

gulp.task('build', ['scripts', 'vendor', 'styles']);

gulp.task('default', ['build']);
