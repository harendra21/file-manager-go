app.factory('AuthService', function($q, $location){
    return {
      authenticate : function () {
        let isAuthenticated = localStorage.getItem("token");
        if(isAuthenticated){
          return true;
        } else {
          return $location.path( "/login" );
  //          return $q.reject('Not Authenticated');
        }
      },
      setToken : function(token) {
        localStorage.setItem("token", token);
      },
      login : function(username, password) {
        console.log(username, password)
        this.setToken(username)
      }
    }
  });