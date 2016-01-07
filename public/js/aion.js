var BASE_API = 'http://localhost:9898/api/v1/'

var app = angular.module('aionApp', []);


app.controller('mainController', function($scope, $http) {
	$http.get(BASE_API + 'job').then(function(response) {$scope.jobs = response.data.jobs});
	$http.get(BASE_API + 'task').then(function(response) {$scope.tasks = response.data.tasks});
});