app.controller("siteCtrl", ['$scope','$http','SweetAlert2', function ($scope,$http,SweetAlert2){
    $scope.addSite = (data) => {
        $http({
            method: 'POST',
            url: base_url+'api/v1/site/',
            headers: { 'Content-Type': 'application/json'},
            data: data
        }).then(function successCallback(response) {
            if(response.statusText == "OK"){
                Swal.fire(
                    'Added!',
                    'Your site has been added.',
                    'success'
                )
            }else{
                Swal.fire(
                    'Error!',
                    'Error while adding new site.',
                    'error'
                )
                console.log(response.statusText)
            }
        }, function errorCallback(response) {
            console.log("Here 111",response)
        });
    } 
}])