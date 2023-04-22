<?php 
    if (session_status() !== PHP_SESSION_ACTIVE) {
        session_start();
    }

    $sessionControl = false;

    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "admin"){
            $sessionControl = true;
        }
    }

    if (!$sessionControl) header('Location: ../index.php');
?>

<!DOCTYPE html>
<html lang="en">
    <html>
        <head>
            <meta charset="utf-8" />
            <title>Desafio - Red Team SW</title>
            <link rel="icon" type="image/x-icon" href="../../assets/images/red_team_sw.jpg">
        </head>
    </html>
    <body width="100%" height="100%" style="display:flex; align-items: center; justify-content: center; flex-direction: column; margin-top: 5%;">
        <h1>Bem vindo ao sistema, <b>admin</b>!</h1>
        <img src="../../assets/images/red_team_sw.jpg">
        <br>
        <div style="display:flex; align-items: center; justify-content: center; flex-direction: row; margin-top: 5%;">
            <a href="../../index.html">About</a>
            <a href="loader.php?pagina=projetos">Projetos</a>
            <a href="loader.php?pagina=busque">Buscar</a>
            <a href="loader.php?pagina=inserir">Inserir</a>
            <a href="loader.php?pagina=logout" id="bLogout">Logout</a>
        </div>
    </body>
</html>

<style>

    a{
        text-decoration: none;
        padding: 10px;
        margin: 5px;
        color: black;
        background-color: lightgrey;
        border: 1px solid;
        border-radius: 10px;
    }

    a:hover{
        background-color: grey;
    }

    #bLogout{
        background-color: red;
    }

    #bLogout:hover{
        background-color: black;
        color:white;
    }

</style>
