app.controller("authCtrl", ['$scope','AuthService','SweetAlert2', function ($scope,AuthService,SweetAlert2){
    
//  SweetAlert2.fire({
//    title: 'Are you sure?',
//    text: "You won't be able to revert this!",
//    icon: 'warning',
//    showCancelButton: true,
//    confirmButtonColor: '#3085d6',
//    cancelButtonColor: '#d33',
//    confirmButtonText: 'Yes, delete it!'
//  }).then((result) => {
//    if (result.value) {
//      Swal.fire(
//        'Deleted!',
//        'Your file has been deleted.',
//        'success'
//      )
//    }
//  })
  
  $scope.login = (data) => {
    AuthService.login(data.username, data.password)
  }
}])

//http://recepuncu.github.io/ngSweetAlert2/#!/home