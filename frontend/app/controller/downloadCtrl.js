app.controller("downloadCtrl", ['$rootScope','SweetAlert2','fileService', function ($rootScope,SweetAlert2,fileService){
    var scope = $rootScope
    scope.msg = ""
    scope.download = (file) => {
        fileService.download(file, scope)
    }
}])