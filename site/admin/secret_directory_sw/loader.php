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
<html>
    <head>
        <meta charset="utf-8" />
        <title>Desafio - Red Team SW</title>
        <link rel="icon" type="image/x-icon" href="../../assets/images/red_team_sw.jpg">
    </head>
</html>

<?php

if(isset($_GET["pagina"])){


    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "admin"){
            include $_GET["pagina"] . ".php";
        }
    } else{
        echo "Usuário não autenticado";
    }

}else{
    echo "Desculpe, não tem nada aqui...";
}

?>