var gulp = require('gulp');
var sass = require('gulp-sass');

gulp.task('default', ['styles', 'watch']);

gulp.task('styles', function(){
  return gulp.src('./app/assets/css/scss/site.scss')
    .pipe(sass())
    .pipe(gulp.dest('./app/assets/css/'));
});

gulp.task('watch', function () {
  gulp.watch('./app/assets/css/**/*.scss', ['styles']);
});
