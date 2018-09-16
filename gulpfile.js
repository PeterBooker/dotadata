var gulp = require('gulp');
var csso = require('gulp-csso');
var rename = require('gulp-rename')

var gulpCopy = require('gulp-copy');

var sass = require('gulp-sass');
var scss = {
    in: 'assets/scss/main.scss',
    out: 'assets/css/',
    watch: 'assets/scss/**/*',
    sassOpts: {
        outputStyle: 'nested',
        precison: 3,
        errLogToConsole: true,
        includePaths: ['./node_modules/bootstrap/scss']
    }
};

// copy deps to build dir
gulp.task('copydeps', function () {
    return gulp.src([
        'node_modules/jquery/dist/jquery.slim.min.js',
        'node_modules/bootstrap/dist/js/bootstrap.bundle.min.js',
        'web/assets/css/main.min.css'
        ])
        .pipe(gulp.dest('build/assets/'));
});

// copy files to build dir
gulp.task('copy', function () {
    return gulp.src([
        'public/'
        ])
        .pipe(gulp.dest('build/'));
});

// compile scss
gulp.task('sass', function () {
    return gulp.src(scss.in)
        .pipe(sass(scss.sassOpts))
        .pipe(csso())
        .pipe(rename('main.min.css'))
        .pipe(gulp.dest(scss.out));
});

// default task
gulp.task('default', ['sass'], function () {
     gulp.watch(scss.watch, ['sass']);
});

// build task
gulp.task('build', ['copy', 'copydeps'], function () {
    // nothing else
});