var gulp = require("gulp");

var minifyCSS = require("gulp-clean-css");
var uglify = require("gulp-uglify");
var autoprefixer = require('gulp-autoprefixer');
var concat = require("gulp-concat");
var livereload = require("gulp-livereload");

gulp.task('minify-css', function () {
  return gulp.src('www/css/*.css')
    .pipe(autoprefixer({
      browsers: ['last 4 versions'],
      cascade: false,
      flexbox: true
    }))
    .pipe(minifyCSS())
    .pipe(gulp.dest('www-dist/css')); 
});

gulp.task('minify-js', function () {
  return gulp.src(['www/js/*.js', '!www/js/all.js'])
    .pipe(concat('all.js'))
    .pipe(uglify())
    .pipe(gulp.dest('www-dist/js'));
});

gulp.task('js-dev', function () {
  return gulp.src(['www/js/*.js', '!www/js/all.js'])
    .pipe(concat('all.js'))
    .pipe(gulp.dest('www/js'))
    .pipe(livereload());
});

gulp.task('copy-resources', function () {
  return gulp.src([
    'www/**',
    '!**/*.js',
    '!**/*.css'])
    .pipe(gulp.dest('www-dist'));
});

gulp.task('watch', function () {
  livereload.listen();
  gulp.watch('www/js/*.js', ['js-dev']);
});

gulp.task('default', ['minify-js', 'minify-css', 'copy-resources'])