var gulp = require("gulp");

var minifyCSS = require("gulp-clean-css");
var uglify = require("gulp-uglify");

gulp.task('minify-css', function () {
  return gulp.src('www/css/*.css')
    .pipe(minifyCSS())
    .pipe(gulp.dest('www-dist/css')); 
});

gulp.task('minify-js', function () {
  return gulp.src('www/js/*.js')
    .pipe(uglify())
    .pipe(gulp.dest('www-dist/js'));
});

gulp.task('copy-resources', function () {
  return gulp.src([
    'www/**',
    '!**/*.js',
    '!**/*.css'])
    .pipe(gulp.dest('www-dist'));
});

gulp.task('default', ['minify-js', 'minify-css', 'copy-resources'])