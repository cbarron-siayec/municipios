<!DOCTYPE html>
<html lang="en">
<head>
 <meta charset="utf-8">
 <meta http-equiv="X-UA-Compatible" content="IE=edge">
 <meta name="viewport" content="width=device-width, initial-scale=1">
 <title>Usuarios</title>
 <meta charset = "utf-8">
</head>
<body>
<nav class="navbar navbar-default">
  <div class="container-fluid">
    <div class="navbar-header">
        
        <ul class="nav nav-tabs">
            <li role="presentation" class="active">
                <a class="navbar-brand" target="_blank" href="http://grupo-siayec.com.mx/">    
                <img src="../municipios/img/cropped-iconSiayec-180x180.png" alt="Logo SIAYEC" height="30" width="30">
                </a>
            </li>
            <li>
             <ul class="nav nav-pills" role="tablist">
                 <li role="presentation"> <a><h4>Sistema de Gestión Indicadores Municipales Benito Juárez</h4></a></li>
            </ul>
            </li>
        </ul>
    </div>
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav navbar-right">
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Saul Zamora <span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Indicadores</a></li>
            <li><a href="#">Usuarios</a></li>
            <li role="separator" class="divider"></li>
            <li><a href="#">Cerrar Sesión</a></li>
          </ul>
        </li>
      </ul>
    </div><!-- /.navbar-collapse -->
  </div>
</nav>
<div class="container-fluid">
    <div class="panel panel-default">
  <div class="panel-heading">
    <h3 class="panel-title">Administración de Usuarios</h3>
  </div>
  <div class="panel-body">
         <table class="table table-hover">
            <thead>
              <tr>
                <th><a><span class="glyphicon glyphicon-user" aria-hidden="true"></span></a></th>
                <th><a>Tipo de Usuario &nbsp;<span class="glyphicon glyphicon-leaf" aria-hidden="true"></span></a></th>
                <th><a><span class="glyphicon glyphicon-calendar" aria-hidden="true"></span></a></th>
                <th><a><span class="glyphicon glyphicon-pencil" aria-hidden="true"></span></a></th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Saúl Zamora</td>
                <td>Admin</td>
                <td>11/01/2018</td>
                <td><a href="admin_panel.html">Editar</a></td>
              </tr>
                <?php
                     $result = mysql_query("SELECT count(id) FROM usuarios;")or die(mysql_error());
                     $array = mysql_query("SELECT * FROM usuarios;")or die(mysql_error());
                    for($i = 1;$i <= $result;$i++){
                        <tr>
                            <td>$array["usuario"]</td>
                            <td>$array["tipo"]</td>
                            <td>11/04/2018</td>
                            <td><a href="admin_panel.html">Editar</a></td>
                        <tr>
                    }
                ?>
            </tbody>
          </table>
  </div>
</div>
    <div class="row">
        <div class="col-md-4">
        </div>
        <div class="col-md-4">
        </div>
        <div class="col-md-4">
            <div class="embed-responsive embed-responsive-4by3">
                <iframe class="embed-responsive-item" src="../municipios/img/delegacion_benito_juarez_logo.png"></iframe>
            </div>
        </div>
    </div>
</div>
<hr />
 </body>
    <div>
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <!-- Optional theme -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
    <!-- Latest compiled and minified JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    </div>
    <script>
        
        function verificarPasswords (){
            password = document.getElementById("password").value;
            confirmPassword = document.getElementById("confirmPassword").value;
            goodPassword = document.getElementById("goodPassword");
            badPassword = document.getElementById("badPassword");
            if(password==confirmPassword && password != ""){
               goodPassword.style.display = "block";
               badPassword.style.display = "none";
               }
            else{
                goodPassword.style.display = "none";
               badPassword.style.display = "block";
            }
        }
        
    </script>
</html>