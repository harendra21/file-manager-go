app.controller("homeCtrl", ['$rootScope','$http','SweetAlert2','$location','fileService', function ($rootScope,$http,SweetAlert2,$location,fileService){
  var scope = $rootScope
  scope.parent = $location.search().p;

	scope.getFiles = () => {
    var parent = ''
    scope.parent == undefined ? (parent = '/') : (parent = scope.parent+"/")
    scope.parent = scope.cleanPath(parent)
    fileService.fetchDir(scope.parent, scope, 'files')
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
        fileService.delete(path, scope)
      }
    })
  }

  scope.movecopy = (parent, folder, action) => {
    scope.moveFolderName = folder
    scope.currentDir = scope.cleanPath(parent+'/'+folder);
    scope.action = action
    var dlgElem = angular.element("#movecopylDlg");
    if (dlgElem) {
      dlgElem.modal("show");
    }
    fileService.fetchDir(parent, scope, 'dirs')
  }

  scope.selectFolder = (folder) => {
    scope.selectedFolder = folder
  }

  scope.openDir = (folder) => {
    scope.parent = scope.cleanPath(folder)
    fileService.fetchDir(scope.parent, scope, 'dirs')
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
    var source = scope.cleanPath(scope.currentDir)
    var destination = scope.cleanPath(parentDir+'/'+scope.moveFolderName)
    fileService.movecopy(scope.action, source, destination, scope)
    var dlgElem = angular.element("#movecopylDlg");
    if (dlgElem) { dlgElem.modal("hide"); }
  }

  scope.create = (type) => {
    scope.msg = "Creating"
    scope.type = type
    var dlgElem = angular.element("#createDlg");
    if (dlgElem) { dlgElem.modal("show"); }
  }

  scope.doCreate = (name, type) => {
    var parent = ''
    scope.parent == undefined ? (parent = '/') : (parent = scope.parent+"/")
    fileService.create(parent, name, type, scope)
    var dlgElem = angular.element("#movecopylDlg");
    if (dlgElem) { dlgElem.modal("hide"); }
  }

  scope.cleanPath = (path) => {
    return path.replace('//','/')
  }

  scope.rename = (file) => {
    scope.msg = "Renaming"
    scope.name = file
    var dlgElem = angular.element("#renameDlg");
    if (dlgElem) { dlgElem.modal("show"); }
  }

  scope.doRename = (newname) => {
    var source = scope.cleanPath(scope.parent+'/'+scope.name)
    var destination = scope.cleanPath(scope.parent+'/'+newname)
    fileService.movecopy('move', source, destination, scope)
    var dlgElem = angular.element("#renameDlg");
    if (dlgElem) { dlgElem.modal("hide"); }
  }

  scope.zip = (filename) => {
    var source = scope.cleanPath(scope.parent+'/'+filename)
    var destination = scope.cleanPath(scope.parent+'/'+filename+'.zip')
    fileService.zipUnzip('zip', source, destination, scope)
  }

  scope.unzip = (filename) => {
    var newfilename = filename.replace('.zip','')
    var source = scope.cleanPath(scope.parent+'/'+filename)
    var destination = scope.cleanPath(scope.parent+'/'+newfilename)
    fileService.zipUnzip('unzip', source, destination, scope)
  }
}]);