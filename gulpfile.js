'use strict';

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
    return gulp.src('static/src/main.jsx')
        .pipe(browserify({
            debug: false,
            extensions: ['.jsx', '.js', '.json'],
            transform: [reactify, literalify.configure({
                react: 'window.React',
                reactdom: 'window.ReactDOM'
            })]
        }))
        .on('error', function(err) {
            gutil.log(err.message)
        })
        .pipe(uglify())
        .pipe(rename('client.min.js'))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('vendor', function() {
    gulp.src([
        'bower_components/react/react.min.js',
        'bower_components/react/react-dom.min.js',
        'bower_components/jquery/dist/jquery.min.js'
    ])
        .pipe(concat('vendor.js'))
        .pipe(uglify())
        .pipe(rename('vendor.min.js'))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('styles', function () {
    var lessStream = gulp.src('static/styles/*.less')
        .pipe(less());

    var cssStream = gulp.src('static/styles/*.css');

    return merge(lessStream, cssStream)
        .pipe(concat('all.css'))
        .pipe(cleanCSS({compatibility: 'ie8'}))
        .pipe(rename('styles.min.css'))
        .pipe(gulp.dest('static/dist'));
});

gulp.task('watch', function () {
    gulp.watch(["static/src/*.jsx"], ['scripts']);
    gulp.watch(["static/style/*.less", "static/style/*.css"], ['styles']);
});

gulp.task('build', ['scripts', 'vendor', 'styles']);

gulp.task('default', ['build'])
