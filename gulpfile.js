'use strict';

var gulp = require('gulp');
var uglify = require('gulp-uglify');
// var rename = require('gulp-rename');
// var less = require('gulp-less');
// var cleanCSS = require('gulp-clean-css');
// var webpack = require('webpack-stream');
// var merge = require('merge-stream');
// var concat = require('gulp-concat');
var react = require('gulp-react');
// var reactify = require('gulp-reactify');

gulp.task('jsx', function() {
    return gulp.src('static/src/*.jsx')
        .pipe(react())
        .pipe(uglify())
        .pipe(gulp.dest('static/dist'));
});

// gulp.task('scripts', function () {
//     return gulp.src("src/*.js")
//         .pipe(webpack({
//             externals: {
//                 "vkapi": "VK",
//                 "jquery": "jQuery"
//             }
//         }))
//         .pipe(uglify())
//         .pipe(rename('all.min.js'))
//         .pipe(gulp.dest('dist'));
// });
//
// gulp.task('styles', function () {
//     var lessStream = gulp.src('style/*.less')
//         .pipe(less());
//
//     var cssStream = gulp.src('style/*.css');
//
//     return merge(lessStream, cssStream)
//         .pipe(concat('all.css'))
//         .pipe(cleanCSS({compatibility: 'ie8'}))
//         .pipe(rename('all.min.css'))
//         .pipe(gulp.dest('dist'));
// });
//
gulp.task('watch', function () {
    gulp.watch(["static/src/*.jsx"], ['jsx']);
    // gulp.watch(["style/*.less", "style/*.css"], ['styles']);
});
//
gulp.task('build', ['jsx']);

gulp.task('default', ['build'])
