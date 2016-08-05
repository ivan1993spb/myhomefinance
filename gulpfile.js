'use strict';

var entryPoint = 'static/src/main.jsx';

var browserify = require('gulp-browserify'),
    cleanCSS = require('gulp-clean-css'),
    concat = require('gulp-concat'),
    gulp = require('gulp'),
    gutil = require('gulp-util'),
    less = require('gulp-less'),
    merge = require('merge-stream'),
    reactify = require('reactify'),
    rename = require('gulp-rename'),
    uglify = require('gulp-uglify');

gulp.task('scripts', function() {
    return gulp.src(entryPoint)
        .pipe(browserify({
            debug: false,
            extensions: ['.jsx', '.js', '.json'],
            transform: [reactify]
        }))
        .on('error', function(err) {
            gutil.log(err.message)
        })
        //.pipe(uglify())
        .pipe(rename('client.js'))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('styles', function() {
    return merge(
        gulp.src('static/src/styles/*.less').pipe(less()),
        gulp.src('static/src/styles/*.css')
    )
        .pipe(concat('styles.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('watch', function() {
    gulp.watch(["static/src/*.jsx", "static/src/*.js"], ['scripts']);
    gulp.watch(["static/src/styles/*.less", "static/src/style/*.css"], ['styles']);
});

gulp.task('build', ['scripts', 'styles']);

gulp.task('default', ['build']);
