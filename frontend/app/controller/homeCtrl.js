app.controller("homeCtrl", ['$rootScope','$http','SweetAlert2','$location', function ($rootScope,$http,SweetAlert2,$location){
  var scope = $rootScope
  scope.parent = $location.search().p;

	scope.getFiles = () => {
    if (scope.parent == undefined) {
      var parent = '/'
    }else{
      var parent = scope.parent+"/"
    }
    scope.parent = scope.cleanPath(parent)
    scope.msg = "Fetching files"
    $http({
      method: 'GET',
      url: base_url+'api/v1/file?p='+parent
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File Fetched"
            scope.files = response.data.data
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }

  scope.delete = (path) => {
    SweetAlert2.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
      if (result.isConfirmed) {

        scope.msg = "Deleting"
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
                  Swal.fire(
                    'Not Deleted!',
                    scope.msg,
                    'error'
                  )
              }
            }else{
              scope.msg = response.statusText
                Swal.fire(
                  'Not Deleted!',
                  scope.msg,
                  'error'
                )
            }
        },function errorCallback(response) {
          scope.msg = response
          Swal.fire(
            'Not Deleted!',
            scope.msg,
            'error'
          )
        });
      }
    })
  }

  scope.movecopy = (parent, folder, action) => {
    scope.moveFolderName = folder
    scope.currentDir = scope.cleanPath(parent+'/'+folder);
    scope.action = action

    $http({
      method: 'GET',
      url: base_url+'api/v1/file?p='+parent
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File Fetched"
            scope.dirs = response.data.data
           
            var dlgElem = angular.element("#movecopylDlg");
            if (dlgElem) {
              dlgElem.modal("show");
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
  }

  scope.selectFolder = (folder) => {
    scope.selectedFolder = folder
  }

  scope.openDir = (folder) => {
    scope.parent = scope.cleanPath(folder)
    $http({
      method: 'GET',
      url: base_url+'api/v1/file?p='+scope.parent
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File Fetched"
            scope.dirs = response.data.data   
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }

  scope.goUp = (type) => {
    var paths = scope.parent.split("/")
    var last = paths[paths.length - 1]
    if (last == ""){
      var path = paths.slice(0, -2)
    }else{
      var path = paths.slice(0, -1)
    }
    path = path.join("/")
    scope.parent = path
    if (type == "file"){
      $location.search('p', scope.parent);
    }else{
      scope.openDir(scope.parent)
    }
    
  }

  scope.doMove = (parentDir) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/'+scope.action+'?source='+scope.cleanPath(scope.currentDir)+'&destination='+scope.cleanPath(parentDir)+'/'+scope.moveFolderName
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File "+scope.action
            setTimeout(() => {
              var dlgElem = angular.element("#movecopylDlg");
              if (dlgElem) {
                dlgElem.modal("hide");
              }
              scope.getFiles()
            },100)
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }

  scope.create = (type) => {
    scope.msg = "Creating"
    scope.type = type
    var dlgElem = angular.element("#createDlg");
    if (dlgElem) {
      dlgElem.modal("show");
    }
  }

  scope.doCreate = (name, type) => {

    if (scope.parent == undefined) {
      var parent = '/'
    }else{
      var parent = scope.parent+"/"
    }

    $http({
      method: 'GET',
      url: base_url+'api/v1/file/create?parent='+parent+'&name='+name+'&type='+type
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "Created successfully"
            setTimeout(() => {
              var dlgElem = angular.element("#createDlg");
              if (dlgElem) {
                dlgElem.modal("hide");
              }
              scope.getFiles()
            },100)
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }

  scope.cleanPath = (path) => {
    path = path.replace('//','/')
    return path
  }

  scope.rename = (file) => {
    console.log(scope.parent)
    scope.msg = "Renaming"
    scope.name = file
    var dlgElem = angular.element("#renameDlg");
    if (dlgElem) {
      dlgElem.modal("show");
    }
  }

  scope.doRename = (newname) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/move?source='+scope.cleanPath(scope.parent)+'/'+scope.name+'&destination='+scope.cleanPath(scope.parent)+'/'+newname
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File Renamed successfully"
            setTimeout(() => {
              var dlgElem = angular.element("#renameDlg");
              if (dlgElem) {
                dlgElem.modal("hide");
              }
              scope.getFiles()
            },100)
            
          }else{
             scope.msg = response.data.error_msg
          }
        }else{
          scope.msg = response.statusText
        }
    },function errorCallback(response) {
      scope.msg = response
    });
  }

  scope.zip = (filename) => {
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/zip?source='+scope.cleanPath(scope.parent)+'/'+filename+'&destination='+scope.cleanPath(scope.parent)+'/'+filename+'.zip'
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File zipped successfully"
            Swal.fire(
              'Success',
              scope.msg,
              'success'
            )
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
  }
  scope.unzip = (filename) => {
    var newfilename = filename.replace('.zip','')
    $http({
      method: 'GET',
      url: base_url+'api/v1/file/unzip?source='+scope.cleanPath(scope.parent)+'/'+filename+'&destination='+scope.cleanPath(scope.parent)+'/'+newfilename
    }).then(function successCallback(response) {
        if(response.statusText == "OK"){
          if (response.data.status == 1) {
            scope.msg = "File unipped successfully"
            Swal.fire(
              'Success',
              scope.msg,
              'success'
            )
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
  }

}]);