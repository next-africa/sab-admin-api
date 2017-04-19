/**
 * Created by pdiouf on 2017-04-15.
 */

var gulp  = require ('gulp');
var watch = require('gulp-watch');
var shell = require('gulp-shell');
var util = require('util');
var del   = require('del')

var config = {
    assetsDir : 'build',
    appAdminDir : '../appengine/admin-module'
}
gulp.task('watch', function (){
    gulp.watch('src/**/*', ['default'], {ignoreInitial: false})


})
gulp.task('clean', function(callback) {
    del([ config.assetsDir + '/**/*' ], callback);
    setTimeout(function() {
        util.log('clean task executed.');
        cb(null);
    }, 3000);
});


gulp.task('buildSrc', function (){
    shell('npm run build')
    setTimeout(function() {
        util.log('build task executed.');
    }, 3000);

});
gulp.task('copyBuildFolder',['buildSrc'], function(){
    return gulp.src('build/**/*')
        .pipe(gulp.dest(config.appAdminDir))

});
gulp.task('default', ['copyBuildFolder']);
