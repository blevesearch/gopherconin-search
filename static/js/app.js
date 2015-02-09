'use strict';


// Declare app level module which depends on filters, and services
angular.module('myApp', [
  'ngRoute',
]).
filter('relDate', function() {
        return function(dstr) {
            return moment(dstr).fromNow();
        };
}).
filter('fosdemDay', function() {
        return function(dstr) {
            return moment(dstr).format('dddd');
        };
}).
filter('fosdemTime', function() {
        return function(dstr) {
            return moment(dstr).format('H:mm');
        };
}).
config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
  $routeProvider.when('/search/', {templateUrl: '/static/partials/search/syntax.html', controller: 'SearchCtrl'});
  $routeProvider.otherwise({redirectTo: '/search/'});
  $locationProvider.html5Mode(true);
}]);
