var gulp = require('gulp')
var browserify = require('browserify')
var source = require('vinyl-source-stream')

gulp.task('browserify', function() {
  return browserify('./client/app.js')
    .bundle()
    .pipe(source('main.js'))
    .pipe(gulp.dest('./build/static/'));
});
