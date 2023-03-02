var app = angular.module("mainApp",["ngRoute"]);
var host = window.location.host; 
host = host.split(":")[0]
var base_url =  window.location.protocol + "//" +host+":8083/"

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
app.service('fileService', ['$http',function($http) {
  this.fetchDir = (dir, scope, type) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file?p='+dir
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File Fetched"
            if (type == "files"){
              scope.files = response.data.data
            }else{
              scope.dirs = response.data.data
            }
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  },
  this.delete = (path, scope) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/delete?source='+path
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "Deleted"
            scope.getFiles()
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  },
  this.movecopy = (action, source, destination, scope) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/'+action+'?source='+source+'&destination='+destination
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File "+scope.action
            scope.getFiles()
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  },
  this.create = (parent, name, type, scope) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/create?parent='+parent+'&name='+name+'&type='+type
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "Created successfully"
            scope.getFiles()
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }, this.zipUnzip = (action, source, destination, scope) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/'+action+'?source='+source+'&destination='+destination
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File zipped successfully"
            scope.getFiles()
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  },
  this.download = (file, scope) => {
    var url = encodeURI(file.url.toString())
    var name = file.name
      $http({
          method: 'POST',
          url: base_url+'api/v1/download?p='+scope.parent,
          data: {url, name}
      }).then(function successCallback(response) {
          if(response.statusText == "OK"){
              if (response.data.status == 1) {
                  scope.msg = "Download Complete"
                  scope.getFiles()
              }else{
                  scope.msg = response.data.error_msg
              }
          }else{
              scope.msg = response.statusText
          }
      }, function errorCallback(response) {
          scope.msg = response
      });
    }
}]);