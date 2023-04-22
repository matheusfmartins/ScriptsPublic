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

    if ($sessionControl) header('Location: secret_directory_sw/home.php');
?>

<!DOCTYPE html>
<html width="100%" height="100%">
    <head>
        <meta charset="utf-8" />
        <title>Desafio - Red Team SW</title>
        <link rel="icon" type="image/x-icon" href="assets/images/red_team_sw.jpg">
    </head>
    <body width="100%" height="100%" style="display:flex; align-items: center; justify-content: center; flex-direction: column; margin-top: 5%;">
        <h1>Sistema de controle</h1>
        <img src="../assets/images/red_team_sw.jpg">
        <br>
        <form action="index.php" method="POST" style="display:flex; align-items: center; justify-content: center; flex-direction: column;">

            <input type="text" name="username" placeholder="Username">
            <input type="password" name="password" placeholder="Password">
            <br>
            <input type="submit" value="Entrar">
        </form>
        
    </body>
</html>

<style>
    input{
        padding: 5px;
        margin-top: 3px;
    }
</style>


<?php

if(isset($_POST['username'])&& isset($_POST['password'])){
    $username = $_POST['username'];
    $password = $_POST['password'];
    
    if ($username == "admin" && $password == "q1w2e3r4t5"){
        $_SESSION['user'] = "admin";
        header('Location: secret_directory_sw/home.php');
    } else{
        echo "Usuário ou senha incorretos!";
    }
    
}

?>