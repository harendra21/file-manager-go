app.controller("homeCtrl", ['$scope','$rootScope','$http','SweetAlert2','$location', function ($scope,$rootScope,$http,SweetAlert2,$location){
  $scope.parent = $location.search().p;

	$scope.getFiles = () => {
    if ($scope.parent == undefined) {
      var parent = '/'
    }else{
      var parent = $scope.parent+"/"
    }
    $scope.parent = parent
    $scope.msg = "Fetching files"
    $http({
      method: 'GET',
      url: base_url+'api/v1/download/files?p='+parent
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            $scope.msg = "File Fetched"
            $scope.files = response.data.data
          }else{
             $scope.msg = response.data.error_msg
          }
        }else{
          $scope.msg = response.statusText
        }
    },function errorCallback(response) {
      $scope.msg = response
    });
  }
}]);