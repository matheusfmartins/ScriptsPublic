<?php 
    if (session_status() !== PHP_SESSION_ACTIVE) {
        session_start();
    }

    $sessionControl = false;

    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "keeggo"){
            $sessionControl = true;
        }
    }

    if ($sessionControl) header('Location: secret_directory_kg/home.php');

    if(isset($_POST['username'])&& isset($_POST['password'])){
        $username = $_POST['username'];
        $password = $_POST['password'];
        
        if ($username == "keeggo" && $password == "1q2w3e"){
            $_SESSION['user'] = "keeggo";
            header('Location: secret_directory_kg/home.php');
        } else{
            if ($username == "keeggo"){
                echo "</br>Senha incorreta!";
            }else{
                echo "</br>Usuário incorreto!";
            }
        }
        
    }
?>

<!DOCTYPE html>
<html width="100%" height="100%">
    <head>
        <meta charset="utf-8" />
        <title>Desafio Keeggo</title>
    </head>
    <body width="100%" height="100%" style="display:flex; align-items: center; justify-content: center; flex-direction: column; margin-top: 5%;">
        <h1>Sistema de controle da Keeggo</h1>
        <br>
        <img src="../assets/images/keeggo.png" height="150px">
        <form action="index.php" method="POST" style="display:flex; align-items: center; justify-content: center; flex-direction: column;">
            <input type="text" name="username" placeholder="Username">
            <input type="password" name="password" placeholder="Password">
            </br>
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