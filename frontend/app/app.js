var app = angular.module("mainApp",["ngRoute"]);
var base_url = "http://localhost:8083/"

app.run(function($rootScope, $location){
  //If the route change failed due to authentication error, redirect them out
  $rootScope.$on('$routeChangeError', function(event, current, previous, rejection){
    if(rejection === 'Not Authenticated'){
      $location.path('/');
    }
  })
});
app.factory('SweetAlert2', ['$rootScope', '$q',function ($rootScope, $q) {
  return {
    fire: function (args1, args2, args3) {
        var deferred = $q.defer();
        $rootScope.$evalAsync(function () {
            if (args1 != null && args2 == null && args3 == null) {
                let opened = Swal.fire(args1);
                deferred.resolve(opened);
            }
            else if (typeof args1 === 'string' && typeof args2 === 'string' && typeof args3 === 'string') {
                let opened = Swal.fire(args1, args1, args3);
                deferred.resolve(opened);
            }
        });
        return deferred.promise;
    }
  };
}]);