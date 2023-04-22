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
        <h1>Inserir</h1>
        <img src="../../assets/images/construcao.jpg" width="500px">
    </body>
</html>