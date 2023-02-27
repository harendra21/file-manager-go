app.controller("downloadCtrl", ['$scope','$http','SweetAlert2', function ($scope,$http,SweetAlert2){
    $scope.msg = ""
    $scope.files = []
    $scope.download = (file) => {
        $scope.msg = "Downloading ..."
        var url = encodeURI(file.url.toString())
        var name = file.name
        console.log(url)
        $http({
            method: 'POST',
            url: base_url+'api/v1/download',
            data: {url, name}
        }).then(function successCallback(response) {
            if(response.statusText == "OK"){
                if (response.data.status == 1) {
                    $scope.msg = "Download Complete"
                }else{
                    $scope.msg = response.data.error_msg
                }
            }else{
                $scope.msg = response.statusText
            }
        }, function errorCallback(response) {
            $scope.msg = response
        });
    }
}])